package encode

import (
	b64 "encoding/base64"

	"github.com/alfred/cmd/inputData"
)

func getData() string {
	return inputData.InputString()
}

func EncB64() string {
	var line = getData()
	sEnc := b64.StdEncoding.EncodeToString([]byte(line))
	return sEnc
}

func DecB64() string {
	var line = getData()
	sDec, _ := b64.StdEncoding.DecodeString(line)
	return string(sDec)
}
