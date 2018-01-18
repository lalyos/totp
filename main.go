package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	otp "github.com/hgfischer/go-otp"
)

var version string = "0.9"
const TOKEN_ENV_NAME = "TOTP_TOKEN"

func main() {
	fmt.Println(totp.Get())
	fmt.Println("version", version)
	flag.Parse()
	token := os.Getenv(TOKEN_ENV_NAME)
	if token == "" {
		token = flag.Arg(0)
		if token == "" {
			fmt.Println(usage)
			os.Exit(1)
		}
	}

	totp := &otp.TOTP{Secret: strings.ToUpper(token), IsBase32Secret: true}
}
