package main

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	_ "embed"

	"github.com/Kwynto/vql-prototype/internal/engine/languages/vqlang/vql1step1analyz"
	"github.com/Kwynto/vql-prototype/internal/engine/languages/vqlang/vql1step2sintax"
	"github.com/Kwynto/vql-prototype/internal/engine/languages/vqlang/vql1step3simant"
	"github.com/Kwynto/vql-prototype/pkg/lib/ecowriter"
	"github.com/Kwynto/vql-prototype/pkg/lib/incolor"
	"github.com/Kwynto/vql-prototype/pkg/lib/ordinarylogger"
)

var (
	//go:embed LICENSE
	sLicense string
)

func readCode(sPath, sName string) (string, error) {
	sNameFile := filepath.Join(sPath, sName)
	code, err := ecowriter.FileRead(sNameFile)
	return code, err
}

func main() {
	// Greeting
	fmt.Println(incolor.StringYellowH(sLicense))

	ordinarylogger.Init("./logs", "develop")
	slog.Info("A prototype of the language analyzer. Experimental repository.")

	// запускать здесь.

	// получаем код от пользователя
	sFileName := "code.txt"
	sCode, errRead := readCode("./", sFileName)
	if errRead != nil {
		slog.Error("File not reading.", slog.String("file", sFileName))
		return
	}
	slog.Info("Code was read.", slog.String("file", sFileName))

	// анализ кода
	stLines := vql1step1analyz.Analyz(sCode)
	slog.Info("Произведен анализ кода.")

	// синтаксический разбор кода
	stSintax := vql1step2sintax.Sintax(stLines)
	slog.Info("Произведен синтаксический разбор кода.")

	// симантический разбор кода
	stSimant := vql1step3simant.Simant(stSintax)
	slog.Info("Произведен симантический разбор кода.")
	_ = stSimant

	// запуск кода
	slog.Info("Начало выполнения кода.")
	slog.Info("Код выполнен.")

	time.Sleep(3 * time.Second)
}
