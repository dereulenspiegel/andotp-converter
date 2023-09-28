package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	andotpconverter "github.com/dereulenspiegel/andotp-converter"
	"github.com/dereulenspiegel/andotp-converter/andotp"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatalf("You need to specify an input and output file")
	}
	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	run(inputFile, outputFile)
}

func run(inputFile, outputFile string) {
	andOtpData, err := andotp.Import(inputFile)
	if err != nil {
		log.Fatalf("failed to load andOTP data: %s", err)
	}
	twofasData, err := andotpconverter.FromAndOtpTo2Fas(andOtpData)
	if err != nil {
		log.Fatalf("failed to convert data to 2FAS format: %s", err)
	}
	twoFasBytes, err := json.Marshal(twofasData)
	if err != nil {
		log.Fatalf("failed to marshal 2FAS data: %s", err)
	}
	if err := os.WriteFile(outputFile, twoFasBytes, 0660); err != nil {
		log.Fatalf("failed to write 2FAS data to file %s: %s", outputFile, err)
	}
}
