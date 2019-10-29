package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/twinj/uuid"
)

var (
	bg = NewBoolGen()
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

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

func logErr(err error, message string) {
	if err != nil {
		if message != "" {
			log.Printf("%s - %s", message, err.Error())
		}
		log.Printf("%s", err.Error())
	}
}

func checkErr(err error, message string) {
	if err != nil {
		if message != "" {
			log.Fatalf("%s - %s", message, err.Error())
		}
		log.Fatalf("%s", err.Error())
	}
}

// BoolGen - Bool generator
type BoolGen struct {
	src       rand.Source
	cache     int64
	remaining int
}

// NewBoolGen - Build a bool generator
func NewBoolGen() *BoolGen {
	return &BoolGen{src: rand.NewSource(time.Now().UnixNano())}
}

// SampleBool - Get a random bool
func (b *BoolGen) SampleBool() bool {
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
