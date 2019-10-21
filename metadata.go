package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/camelcase"
	"github.com/gedex/inflector"
	"github.com/twinj/uuid"
)

type (
	metadata struct {
		// ResName is the name of the resource to be created.
		ResName             string `yaml:"name"`
		APIVersion          string `yaml:"apiVer"`
		APIVersionUpper     string
		Plural              string `yaml:"plural"`
		SingularLowercase   string
		PluralLowercase     string
		SingularCamelCase   string
		PluralCamelCase     string
		SingularPascalCase  string
		PluralPascalCase    string
		SingularSnakeCase   string
		PluralSnakeCase     string
		SingularDashed      string
		PluralDashed        string
		Package             string
		PackageDir          string
		MigrationNumber     string
		PropDefs            []propDef `yaml:"propDefs"`
		NonVirtualPropDefs  []propDef
		ClientPropDefs      []propDef
		ModelMatchCond      string
		CreateStatement     string
		AlterStatement      []string
		DropStatement       string
		InsertStatement     string
		FixtureData         []string
		TestCreateForm      string
		TestUpdateForm      string
		TestCreateJSON      string
		TestUpdateJSON      string
		RESTCreateJSON      string
		RESTUpdateJSON      string
		TestInsertValueVars string
		TestUpdateValueVars string
		TestJSONVars        string
	}

	propDef struct {
		Name               string  `yaml:"name"`
		Type               string  `yaml:"type"`
		Length             int     `yaml:"length"`
		IsVirtual          bool    `yaml:"isVirtual"`
		IsKey              bool    `yaml:"isKey"`
		IsUnique           bool    `yaml:"isUnique"`
		AdmitNull          bool    `yaml:"admitNull"`
		Ref                propRef `yaml:"ref"`
		IsEmbedded         bool
		IsBackendOnly      bool
		ModelType          string
		SafeType           string
		SafeTypeMaker      string
		SingularCamelCase  string
		SingularPascalCase string
		FormatString       string
		DashedName         string
		SQLColumn          string
		SQLType            string
		SQLModifier        string
		Value              interface{}
	}

	propRef struct {
		Model    string `yaml:"model"`
		Property string `yaml:"property"`
		FKName   string
		TrgTable string
	}
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func (g *gen) procMetadata() error {
	md := g.meta
	md.ResName = upercaseFirst(md.ResName)
	md.genNameForms()
	md.addIdentification()
	md.addAudit()
	props := md.PropDefs
	for i := range props {
		prop := &props[i]
		// fmt.Println("generators#GenerateMetadata: ", prop.ResName, " - ", i)
		prop.setTypes() //SafeType = safeType(prop)
		prop.SingularCamelCase = toCamelCase(prop.Name)
		prop.SingularPascalCase = toPascalCase(prop.Name)
		prop.DashedName = dashedName(prop.Name)
		prop.SQLColumn = sqlColumn(prop)
		prop.SQLType = sqlType(prop)
		sqlColModifiers(prop)
		sqlFKData(prop, md)
		setShowInClient(prop)
	}
	md.selectNonVirtualPropDefs()
	md.selectClientPropDefs()
	return nil
}

// MakeMetadata - md *constructor.
func MakeMetadata(pkg string) metadata {
	md := metadata{}
	md.PackageDir = getPackageRootDir()
	md.Package = decidePkgName(pkg)
	md.EnsureAPIVersion()
	return md
}

func getPackageRootDir() string {
	pkgDir, err := os.Getwd()
	if err != nil {
		log.Printf("%s", err.Error())
	}
	return pkgDir
}

func decidePkgName(usrProvided string) string {
	if strings.Trim(usrProvided, " ") != "" {
		return usrProvided
	}
	return guessPkgName()
}

func guessPkgName() (name string) {
	// Infer current dir
	pkgDir, err := os.Getwd()
	if err != nil {
		log.Printf("%s", err.Error())
	}
	// Get GOPATH (It can consist of several directories (*) separated by colons)
	gpall := os.Getenv("GOPATH")
	gpsplit := strings.Split(gpall, ":")
	for _, gp := range gpsplit {
		gps := filepath.Join(gp, "src")
		sep := string(filepath.Separator)
		// (*) One of these entries plus filepath separator (i.e. "/")
		gopathsrc := fmt.Sprintf("%s%s", gps, sep)
		// TrimPrefix returns original first argument if it does not find the prefix to trim off.
		name = strings.TrimPrefix(pkgDir, gopathsrc)
		if name != pkgDir {
			// If prefix (gopath + separator + src) can be trimed
			// Package name is the trimmed part from current directory.
			return name
		}
	}
	return "github.com/[username]/[project]"
}

// GenerateMetadata - Generate aditional struct properties
func (md *metadata) GenerateMetadata() {
	md.ResName = upercaseFirst(md.ResName)
	md.genNameForms()
	md.addIdentification()
	md.addAudit()
	props := md.PropDefs
	for i := range props {
		prop := &props[i]
		prop.setTypes()
		prop.SingularCamelCase = toCamelCase(prop.Name)
		prop.SingularPascalCase = toPascalCase(prop.Name)
		prop.DashedName = dashedName(prop.Name)
		prop.SQLColumn = sqlColumn(prop)
		prop.SQLType = sqlType(prop)
		sqlColModifiers(prop)
		sqlFKData(prop, md)
		setShowInClient(prop)
	}
	md.selectNonVirtualPropDefs()
	md.selectClientPropDefs()
}

func (md *metadata) addIdentification() {
	pd := makePropDef("ID", "primary_key", 0, false, true, true, false, true, true, "")
	md.PropDefs = append([]propDef{pd}, md.PropDefs...)
	pd = makePropDef("Name", "string", 0, false, false, false, true, true, false, "")
	md.PropDefs = append(md.PropDefs, pd)
}

func (md *metadata) addDetail() {
	pd := makePropDef("Description", "text", 0, false, false, false, true, true, false, "")
	md.PropDefs = append(md.PropDefs, pd)
}

func (md *metadata) addAudit() {
	pd := makePropDef("CreatedBy", "uuid", 36, false, false, false, false, true, true, "")
	pd.SQLColumn = "created_by_id"
	// propDef.Ref = PropertyRef{Property: "id"}
	md.PropDefs = append(md.PropDefs, pd)
	pd = makePropDef("UpdatedBy", "uuid", 36, false, false, false, false, true, true, "")
	pd.SQLColumn = "updated_by_id"
	// propDef.Ref = PropertyRef{Property: "id"}
	md.PropDefs = append(md.PropDefs, pd)
	pd = makePropDef("CreatedAt", "timestamptz", 0, false, false, false, false, true, true, "")
	md.PropDefs = append(md.PropDefs, pd)
	pd = makePropDef("UpdatedAt", "timestamptz", 0, false, false, false, false, true, true, "")
	md.PropDefs = append(md.PropDefs, pd)
}

func (md *metadata) genNameForms() {
	md.genPlural()
	md.genSingularLowercase()
	md.genPluralLowercase()
	md.genSingularCamelCase()
	md.genPluralCamelCase()
	md.genSingularPascalCase()
	md.genPluralPascalCase()
	md.genSingularSnakeCase()
	md.genPluralSnakeCase()
	md.genSingularDashed()
	md.genPluralDashed()
}

func (md *metadata) genPlural() {
	if md.Plural == "" {
		md.Plural = inflector.Pluralize(md.ResName)
	}
}

// Lowercase
func (md *metadata) genSingularLowercase() {
	md.SingularLowercase = strings.ToLower(md.ResName)
	// fmt.Println("SingularLowercase: ", md.SingularSnakeCase)
	return
}

func (md *metadata) genPluralLowercase() {
	md.PluralLowercase = strings.ToLower(md.Plural)
	// fmt.Println("PluralLowercase: ", md.PluralLowercase)
	// md.PluralLowercase = sqlTable(md *)
}

// CamelCase
func (md *metadata) genSingularCamelCase() {
	md.SingularCamelCase = toCamelCase(md.ResName)
	// fmt.Println("SingularCamelCase: ", md.SingularCamelCase)
	return
}

func (md *metadata) genPluralCamelCase() {
	md.PluralCamelCase = toCamelCase(md.Plural)
	// fmt.Println("PluralCamelCase: ", md.PluralCamelCase)
}

// PascalCase
func (md *metadata) genSingularPascalCase() {
	md.SingularPascalCase = toPascalCase(md.ResName)
	// fmt.Println("SingularPascalCase: ", md.SingularPascalCase)
	return
}

func (md *metadata) genPluralPascalCase() {
	md.PluralPascalCase = toPascalCase(md.Plural)
	// fmt.Println("PluralPascalCase: ", md.PluralCamelCase)
	// md.PluralLowercase = sqlTable(md *)
}

// SnakeCase
func (md *metadata) genSingularSnakeCase() {
	md.SingularSnakeCase = toSnakeCase(md.ResName)
	// fmt.Println("SingularSnakeCase: ", md.SingularSnakeCase)
}

func (md *metadata) genPluralSnakeCase() {
	md.PluralSnakeCase = toSnakeCase(md.Plural)
	// fmt.Println("PluralSnakeCase: ", md.PluralSnakeCase)
}

// Dashed
func (md *metadata) genSingularDashed() {
	md.SingularDashed = toDashedCase(md.ResName)
	// fmt.Println("SingularDashed: ", md.SingularDashed)
}

func (md *metadata) genPluralDashed() {
	md.PluralDashed = toDashedCase(md.Plural)
	// fmt.Println("PluralDashed: ", md.PluralDashed)
}

func (md *metadata) selectNonVirtualPropDefs() {
	props := md.PropDefs
	for i := range md.PropDefs {
		prop := props[i]
		if !prop.IsVirtual {
			md.NonVirtualPropDefs = append(md.NonVirtualPropDefs, prop)
		}
	}
}

func (md *metadata) selectClientPropDefs() {
	props := md.PropDefs
	for i := range md.PropDefs {
		prop := props[i]
		if !prop.IsBackendOnly {
			md.ClientPropDefs = append(md.ClientPropDefs, prop)
		}
	}
}

func (property *propDef) setTypes() {
	propType := property.Type
	switch propType {
	case "id":
		property.ModelType = "Int64"
		property.SafeType = "nulls.Int64"
		property.SafeTypeMaker = "NullsZeroInt64()"
	case "uuid":
		property.ModelType = "UUID"
		property.SafeType = "nulls.UUID"
		property.SafeTypeMaker = "NullsZeroUUID()"
	case "binary":
		property.ModelType = "ByteSlice"
		property.SafeType = "nulls.ByteSlice"
		property.SafeTypeMaker = "NullsEmptyByteSlice()"
	case "boolean":
		property.ModelType = "Bool"
		property.SafeType = "nulls.Bool"
		property.SafeTypeMaker = "NullsFalseBool()"
	case "date":
		property.ModelType = "Time"
		property.SafeType = "nulls.Time"
		property.SafeTypeMaker = "NullsZeroTime()"
	case "datetime":
		property.ModelType = "Time"
		property.SafeType = "nulls.Time"
		property.SafeTypeMaker = "NullsZeroTime()"
	case "decimal":
		property.ModelType = "Float"
		property.SafeType = "nulls.Float"
		property.SafeTypeMaker = "NullsEmptyFloat64()"
	case "float":
		property.ModelType = "Float"
		property.SafeType = "nulls.Float"
		property.SafeTypeMaker = "NullsEmptyFloat64()"
	case "geolocation":
		property.ModelType = "Point"
		property.SafeType = "nulls.NullPoint"
		property.SafeTypeMaker = "NullsZeroPoint()"
	case "integer":
		property.ModelType = "Int64"
		property.SafeType = "nulls.Int64"
		property.SafeTypeMaker = "NullsZeroInt64()"
	case "json":
		property.ModelType = "String"
		property.SafeType = "sqlxtypes.JSONText"
		property.SafeTypeMaker = "NullsEmptyByteSlice()"
	case "primary_key":
		property.ModelType = "UUID"
		property.SafeType = "nulls.UUID"
		property.SafeTypeMaker = "NullsZeroUUID()"
	case "string":
		property.ModelType = "String"
		property.SafeType = "nulls.String"
		property.SafeTypeMaker = "NullsEmptyString()"
	case "text":
		property.ModelType = "String"
		property.SafeType = "nulls.String"
		property.SafeTypeMaker = "NullsEmptyString()"
	case "password":
		property.ModelType = "String"
		property.SafeType = "nulls.String"
		property.SafeTypeMaker = "NullsEmptyString()"
	case "password_confirmation":
		property.ModelType = "String"
		property.SafeType = "nulls.String"
		property.SafeTypeMaker = "NullsEmptyString()"
	case "time":
		property.ModelType = "Time"
		property.SafeType = "nulls.Time"
		property.SafeTypeMaker = "NullsZeroTime()"
	case "timestamp":
		property.ModelType = "Time"
		property.SafeType = "nulls.Time"
		property.SafeTypeMaker = "NullsZeroTime()"
	case "timestamptz":
		property.ModelType = "Time"
		property.SafeType = "nulls.Time"
		property.SafeTypeMaker = "NullsZeroTime()"
	default:
		property.ModelType = "String"
		property.SafeType = "nulls.String"
		property.SafeTypeMaker = "NullsEmptyString()"
	}
}

func toCamelCase(str string) string {
	camelCased := toCamelCaseString(str)
	splitted := camelcase.Split(camelCased)
	splitted[0] = strings.ToLower(splitted[0])
	return strings.Join(splitted, "")
}

func toPascalCase(str string) string {
	camelCased := toCamelCaseString(str)
	splitted := camelcase.Split(camelCased)
	splitted[0] = upercaseFirst(splitted[0])
	return strings.Join(splitted, "")
}

func dashedName(str string) string {
	return toDashedCase(str)
}

func sqlTable(md *metadata) string {
	//return strings.ToLower(inflector.Pluralize(md.Plural))
	// return strings.ToLower(md.Plural)
	return md.PluralSnakeCase
}

func sqlColumn(prop *propDef) string {
	if prop.SQLColumn != "" {
		return prop.SQLColumn
	}
	return toSnakeCase(prop.Name)
}

func sqlType(prop *propDef) string {
	propType := prop.Type
	//propSize := prop.Length
	switch propType {
	case "binary":
		return "BYTEA"
	case "boolean":
		return "BOOLEAN"
	case "date":
		return "DATE"
	case "datetime":
		return "TIMESTAMP"
	case "decimal":
		return "FLOAT(24)"
	case "float":
		return "FLOAT(24)"
	case "geolocation":
		return "GEOGRAPHY(Point,4326)"
	case "integer":
		return "BIGINT"
	case "json":
		return "JSONB"
	case "primary_key":
		return "UUID"
	case "string":
		return "VARCHAR(64)"
	case "text":
		return "TEXT"
	case "time":
		return "TIME"
	case "timestamp":
		return "TIMESTAMP"
	case "timestamptz":
		return "TIMESTAMP WITH TIME ZONE"
	case "uuid":
		return "UUID"
	default:
		return "VARCHAR(64)"
	}
}

func sqlColModifiers(property *propDef) {
	mq := 0
	var m bytes.Buffer
	if property.IsKey {
		mq = mq + 1
		m.WriteString("PRIMARY KEY")
	}
	if property.IsUnique {
		if mq > 0 {
			m.WriteString(" ")
		}
		m.WriteString("UNIQUE")
	}
	if property.AdmitNull {
		if mq > 0 {
			m.WriteString(" ")
		}
		m.WriteString("NULL")
	}
	property.SQLModifier = m.String()
}

func sqlFKData(property *propDef, md *metadata) {
	ref := &property.Ref
	emptyRef := &propRef{}
	if *emptyRef != *ref {
		refModel := inflector.Pluralize(ref.Model)
		refProp := ref.Property
		ref.FKName = strings.ToLower(fmt.Sprintf("%s_%s_fkey", refModel, refProp))
		ref.TrgTable = strings.ToLower(md.Plural)
	}
}

func setShowInClient(property *propDef) {
	if !property.IsEmbedded {
		property.IsBackendOnly = false
	}
}

func newMigrationPrefix() string {
	t := time.Now()
	return t.Format("20060102150405")
}

func getFileWriter(outputFile string, force bool) (*os.File, error) {
	return getFileWriterWithPerm(outputFile, force, 0666)
}

func getFileWriterWithPerm(outputFile string, force bool, perm int) (*os.File, error) {
	flag := os.O_CREATE | os.O_EXCL
	if force {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	p := os.FileMode(perm)
	return os.OpenFile(outputFile, flag, p)
}

func toCamelCaseString(str string) string {
	var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")
	byteSrc := []byte(str)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	//fmt.Printf("Original: %s - Camel: %s", str, string(bytes.Join(chunks, nil)))
	return string(bytes.Join(chunks, nil))
}

func upercaseFirst(str string) string {
	temp := []rune(str)
	temp[0] = unicode.ToUpper(temp[0])
	return string(temp)
}

func lowercaseFirst(str string) string {
	temp := []rune(str)
	temp[0] = unicode.ToLower(temp[0])
	return string(temp)
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func toDashedCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	dashed := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	dashed = matchAllCap.ReplaceAllString(dashed, "${1}-${2}")
	return strings.ToLower(dashed)
}

func generateUUID() uuid.UUID {
	return uuid.NewV4()
}

func generateUUIDString() string {
	return fmt.Sprintf("%v", generateUUID())
}

func generateTestUUIDString(index int) string {
	if index < 6 {
		return fmt.Sprintf("%v%d", "00000000-0000-0000-0000-00000000000", index)
	}
	return fmt.Sprintf("%v", generateUUID())
}

// BoolGen - Bool generator
type BoolGen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func newBoolGen() *BoolGen {
	return &BoolGen{src: rand.NewSource(time.Now().UnixNano())}
}

func (b *BoolGen) dampleBool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}
	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--
	return result
}

func sampleInt(max int) int {
	return r.Intn(max)
}

func sampleDecimal(max int) float32 {
	if max == 0 {
		max = 2
	} else if max < 0 {
		max = -max - 1
	} else {
		max = max + 1
	}
	return float32(max)*r.Float32() - 1
}

func sampleString(prefix string, prefixLength, maxLength int) string {
	pl := min(len(prefix), prefixLength)
	p := prefix[0:pl]
	rl := maxLength - pl
	s := randString(rl)
	return fmt.Sprintf("%s%s", p, s)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func randString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func ensureDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, 0755)
	}
}

// propDef --------------------------------------------------------------------
// makePropDef - Creates an instence of propDef
func makePropDef(name, propType string, length int, isVirtual, isKey, isUnique, admitNull, isEmbedded, isBackendOnly bool, value interface{}) propDef {
	return propDef{
		Name:          name,
		Type:          propType,
		Length:        length,
		IsVirtual:     isVirtual,
		IsKey:         isKey,
		IsUnique:      isUnique,
		AdmitNull:     admitNull,
		IsEmbedded:    isEmbedded,
		IsBackendOnly: isBackendOnly,
		Value:         value,
	}
}

// EnsureAPIVersion - Todo: complete comment
func (md *metadata) EnsureAPIVersion() {
	//fmt.Printf("API Version is %s\n", md.APIVersion)
	if md.APIVersion == "" {
		md.DefaultAPIVersion()
	}
	md.APIVersionUpper = strings.ToUpper(md.APIVersion)
}

// DefaultAPIVersion - Todo: complete comment
func (md *metadata) DefaultAPIVersion() {
	md.APIVersion = "v1"
}

func (prop propDef) isKeyType() bool {
	return prop.Type == "primary_key"
}

func (prop propDef) isUUIDType() bool {
	return prop.Type == "uuid"
}

func (prop propDef) isPasswordType() bool {
	return prop.Type == "password" || prop.Type == "password_confirmation"
}

func (prop propDef) isPasswordHashType() bool {
	return prop.Type == "password_hash"
}

func (prop propDef) isTextType() bool {
	return prop.Type == "string" || prop.Type == "text"
}

func (prop propDef) isIntType() bool {
	return prop.Type == "bigint" || prop.Type == "integer" || prop.Type == "smallint"
}

func (prop propDef) isDecimalType() bool {
	return prop.Type == "float" || prop.Type == "decimal"
}

func (prop propDef) isBooleanType() bool {
	return prop.Type == "boolean"
}

func (prop propDef) isTimeType() bool {
	return prop.Type == "date" || prop.Type == "datetime" || prop.Type == "time" || prop.Type == "timestamp" || prop.Type == "timestamptz"
}

func (prop propDef) isGeoType() bool {
	return prop.Type == "geolocation"
}
