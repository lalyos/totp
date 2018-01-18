package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	otp "github.com/hgfischer/go-otp"
	"github.com/tjgq/clipboard"
)

//var version, gitrev string
var (
	version string = "0.1.0"
	gitrev  string
)

const TOKEN_ENV_NAME = "TOTP_TOKEN"
const usage = `
usage: totp <token-secret>
   either use the first argument to set the totp secret,
   or use the TOTP_TOKEN env variable
`

var clpb bool

func main() {
	vers := flag.Bool("version", false, "prints version")
	flag.BoolVar(&clpb, "clipboard", false, "copies token to system clipboard")
	flag.BoolVar(&clpb, "c", false, "(short) copies token to system clipboard")
	flag.Parse()

	if *vers {
		fmt.Printf("version: %v-%v\n", version, gitrev)
		os.Exit(0)
	}

	token := os.Getenv(TOKEN_ENV_NAME)
	if token == "" {
		token = flag.Arg(0)
		if token == "" {
			fmt.Println(usage)
			os.Exit(1)
		}
	}

	totp := &otp.TOTP{Secret: strings.ToUpper(token), IsBase32Secret: true}
	if clpb {
		clipboard.Set(totp.Get())
		fmt.Fprintln(os.Stderr, "totp token writen to CLIPBOARD")
		os.Exit(0)
	}
	fmt.Print(totp.Get())
}
