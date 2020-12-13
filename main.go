package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alfred/cmd/encode"
	"github.com/alfred/cmd/httpCurl"
	"github.com/alfred/cmd/loadYaml"
)

func main() {

	base64Cmd := flag.NewFlagSet("base64", flag.ExitOnError)
	base64Enc := base64Cmd.Bool("encode", false, "Encode a string in Base64 encrypted")
	base64Dec := base64Cmd.Bool("decode", false, "Decode a Base64 encrypted string to raw string")

	urlCmd := flag.NewFlagSet("url", flag.ExitOnError)
	var fileName string
	urlCmd.StringVar(&fileName, "file", "", "YAML file with parameters.")
	urlReq := urlCmd.Bool("request", false, "Make a HTTP request. Allow all methods (GET, POST, PUT, etc). Use with -file <yaml file>.")

	if len(os.Args) < 2 {
		fmt.Println("Expected a subcommands, found none!")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "base64":
		base64Cmd.Parse(os.Args[2:])
		if *base64Enc && *base64Dec {
			fmt.Println("Choose only one of the following options: -encode or -decode")
			os.Exit(1)
		} else if *base64Enc {
			fmt.Printf(" ~> Encoded text: %v\n", encode.EncB64())
		} else if *base64Dec {
			fmt.Printf(" ~> Decoded text: %v\n", encode.DecB64())
		} else {
			fmt.Println("Choose at least one of the follow option: -encode or -decode")
			os.Exit(1)
		}
	case "url":
		urlCmd.Parse(os.Args[2:])
		if *urlReq {
			if fileName == "" {
				fmt.Println("Please provide yaml file by using -file <file> option")
				os.Exit(1)
			}
			fileReturn := loadYaml.ReqYamlFile(fileName)
			httpCurl.ReadStruct(fileReturn.Url, fileReturn.Method, fileReturn.Body, fileReturn.Headers)
		} else {
			fmt.Println("Choose at least one option. To see them use -h or -help")
			os.Exit(1)
		}

	default:
		fmt.Printf("The subcommand \"%v\" is not valid!\n", os.Args[1])
		os.Exit(1)
	}

}
