package provider

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

/*
Env Variable Needed ex.
$ export ESHIP_USER=user
$ export ESHIP_PASS=pass
$ export ESHIP_URL=http://www.eshipper.com/rpc2
*/
// eShipping API credentials
var username = os.Getenv("ESHIP_USER")
var password = os.Getenv("ESHIP_PASS")
var apiURL = os.Getenv("ESHIP_URL")

type quoteRequest struct {
	XMLName  xml.Name `xml:"EShipper"`
	XMLNS    string   `xml:"xmlns,attr"`
	Username string   `xml:"username,attr"`
	Password string   `xml:"password,attr"`
	Version  string   `xml:"version,attr"`
	Packages struct {
		PackageType string `xml:"type,attr"`
		Package     []struct {
			Weight      float64 `xml:"weight"`
			Description string  `xml:"description"`
			Type        string  `xml:"type"`
			Length      float64 `xml:"length"`
			Height      float64 `xml:"height"`
			Width       float64 `xml:"width"`
		} `xml:"Package"`
	} `xml:"QuoteRequest>Packages"`
}
type qPackage struct {
	Weight      float64 `xml:"Package>weight"`
	Description string  `xml:"Package>description"`
	Type        string  `xml:"Package>type"`
	Length      float64 `xml:"Package>length"`
	Height      float64 `xml:"Package>height"`
	Width       float64 `xml:"Package>width"`
}

func getShippingQuote(weight float64) error {
	// Check existense of credentials from OS Environmental Variables
	if len(username) == 0 {
		log.Fatal("Retrieved no username from Env. Check OS Env vars.")
	}
	if len(password) == 0 {
		log.Fatal("Retrieved no password from Env. Check OS Env vars.")
	}
	if len(apiURL) == 0 {
		log.Fatal("Retrieved no API URL from Env. Check OS Env vars.")
	}
	//var quote quoteRequest
	q := &quoteRequest{XMLNS: "http://www.eshipper.net/XMLSchema", Username: username, Password: password, Version: "3.0.0"}
	q.Packages.PackageType = "Package"
	//q.Package = append(q.Package, qPackage{Weight: 1, Description: "Testing", Type: "Package", Length: 12, Height: 9, Width: 0.5})
	//q.P.Packages = {Weight: 1, Description: "Testing", Type: "Package", Length: 12, Height: 9, Width: 0.5})
	//q.Packages.Package = append(q.Packages.Package, qPackage{Weight: 1, Description: "Testing", Type: "Package", Length: 12, Height: 9, Width: 0.5})

	output, err := xml.MarshalIndent(q, "", " ")
	if err != nil {
		log.Fatalf("Could not encode/marshall XML. Error: %v", err)
	}

	fmt.Printf("\nXML:\n%v\n", string(output))

	// Create and Send Request
	client := &http.Client{}
	xmlBody, err := xml.Marshal(q)
	if err != nil {
		log.Fatalf("Error marshalling/encoding XML Request. Error: %v", err)
	}
	body := new(bytes.Buffer)
	_, err = body.Write(xmlBody)
	if err != nil {
		log.Fatalf("Could not write XML to buffer. Error: %v", err)
	}
	req, err := http.NewRequest("POST", apiURL, body)
	/*
		resp, err := resty.R().SetBody(q).Post(apiURL)
		if err != nil {
			log.Fatalf("HTTP Request failure. Error: %v", err)
		}
	*/
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatalf("Could not dump HTTP request. Error: %v", err)
	}
	fmt.Printf("\nRequest: %v\n", string(requestDump))
	resp, err := client.Do(req)
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalf("Could not dump HTTP response. Error: %v", err)
	}
	fmt.Printf("\nResponse: %v\n", string(responseDump))

	var qq SnipcartQuote
	qq.Rates = append(qq.Rates, Rate{Cost: 1.50, Description: "test"})
	jsonBody, err := json.Marshal(qq)
	if err != nil {
		log.Fatalf("Could not marshall/encode Quote json. Error: %v", err)
	}
	fmt.Printf("\nOutput: %v", string(jsonBody))

	return err
}
