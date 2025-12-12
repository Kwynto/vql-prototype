package main

import (
	"fmt"
	"log/slog"
	"time"

	_ "embed"

	"github.com/Kwynto/vql-prototype/pkg/lib/incolor"
	"github.com/Kwynto/vql-prototype/pkg/lib/ordinarylogger"
)

var (
	//go:embed LICENSE
	sLicense string
)

func main() {
	// Greeting
	fmt.Println(incolor.StringYellowH(sLicense))

	ordinarylogger.Init("./logs", "develop")
	slog.Info("A prototype of the language analyzer. Experimental repository.")

	// запускать здесь.

	time.Sleep(1 * time.Second)
}
