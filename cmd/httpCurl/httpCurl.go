package httpCurl

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urlLib "net/url"
	"strings"
)

func ReadStruct(url, method, body string, headers []string) {

	// Parse URL and Body to usi into http.Request
	reqURL, _ := urlLib.Parse(url)
	reqBody := ioutil.NopCloser(strings.NewReader(body))

	// Contruct Request
	req := &http.Request{
		Method: method,
		URL:    reqURL,
		Body:   reqBody,
	}

	//Inicialize header
	req.Header = make(http.Header)

	// Looping to add the headers
	for i := 0; i < len(headers); i += 1 {
		header := strings.Split(headers[i], ":")
		req.Header.Add(header[0], strings.TrimSpace(header[1]))
	}

	// Make a request
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error:", err)
	}

	// read response body
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	// print response status and body
	fmt.Printf("status: %d\n", res.StatusCode)
	fmt.Printf("body: %s\n", data)
}
