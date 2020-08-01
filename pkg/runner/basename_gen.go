package runner

import (
	"fmt"
	"math/rand"
	"time"
)

type randomIntegerGenerator interface {
	Intn(int) int
}

type dateTimeProvider func() string

var __rnd randomIntegerGenerator
var __dtp dateTimeProvider

func init() {
	__rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	__dtp = func() string {
		return time.Now().Format("2006-01-02T15:04:05.000")
	}
}

func generateBasename() string {
	pattern := "gonel_%s_%07d.go"
	now := __dtp()
	rnd := __rnd.Intn(9999999)

	return fmt.Sprintf(pattern, now, rnd)
}
