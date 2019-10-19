package main

import "log"

type (
	metadata struct {
		// ResName is the name of the resource to be created.
		Package            string
		ResName            string `yaml:"name"`
		APIVersion         string `yaml:"apiVer"`
		APIVersionUpper    string
		Plural             string `yaml:"plural"`
		SingularLowercase  string
		PluralLowercase    string
		SingularCamelCase  string
		PluralCamelCase    string
		SingularPascalCase string
		PluralPascalCase   string
		SingularSnakeCase  string
		PluralSnakeCase    string
		SingularDashed     string
		PluralDashed       string
		PackageDir         string
		PropDefs           []propDef `yaml:"propertyDefs"`
		NonVirtualPropDefs []propDef
		ModelMatchCond     string
		CreateStatement    string
		InsertStatement    string
		RESTCreateJSON     string
		RESTUpdateJSON     string
	}

	propDef struct {
		Name               string  `yaml:"name"`
		Type               string  `yaml:"type"`
		Length             int     `yaml:"length"`
		IsVirtual          bool    `yaml:"isVirtual"`
		IsKey              bool    `yaml:"isKey"`
		IsUnique           bool    `yaml:"isUnique"`
		AdmitNull          bool    `yaml:"admitNull"`
		Ref                propRef `yaml:"references"`
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

func (g *gen) procMetadata() error {
	log.Println("*metadata.process() not implemented")
	return nil
}
