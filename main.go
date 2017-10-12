package main

import (
	"fmt"
	"os"
	"strings"

	otp "github.com/hgfischer/go-otp"
)

var version string = "0.9"

func main() {
	totp := &otp.TOTP{Secret: strings.ToUpper(os.Getenv("GITHUB_TOTP_SECRET")), IsBase32Secret: true}
	fmt.Println(totp.Get())
	fmt.Println("version", version)
}
