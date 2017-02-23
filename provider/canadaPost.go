package provider

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type cpRequest struct {
	XMLName        xml.Name `xml:"mailing-scenario"`
	XMLNS          string   `xml:"xmlns,attr"`
	CustomerNumber string   `xml:"customer-number"`
	Weight         float64  `xml:"parcel-characteristics>weight"`
	ServiceCode    string   `xml:"services>service-code"`
	OriginPostcode string   `xml:"origin-postal-code"`
	Destination    struct {
		DestinationPostcode string `xml:"postal-code"`
	} `xml:"destination>domestic"`
}
type cpResponse struct {
	XMLName     xml.Name `xml:"price-quotes"`
	ServiceCode string   `xml:"price-quote>service-code"`
	ServiceName string   `xml:"service-name"`
}

/*
Env Variable Needed ex.
$ export CAPost_USER=user
$ export CAPost_PASS=pass
$ export CAPost_URL=http://www.eshipper.com/rpc2
$ export CAPost_CustNum=000
*/
func GetCanadaPostRate(weight float64, destPostcode string, originPostcode string) (Rate, error) {

	var cpRate Rate

	// Encode API Credentials
	username := os.Getenv("CAPost_USER")
	password := os.Getenv("CAPost_PASS")
	apiURL := os.Getenv("CAPost_URL")
	custnum := os.Getenv("CAPost_CustNum")
	if len(username) == 0 {
		log.Fatalf("Username for CanadaPost API cannot be empty. Check OS Env.")
	}
	if len(password) == 0 {
		log.Fatalf("Password for CanadaPost API cannot be empty. Check OS Env.")
	}
	if len(apiURL) == 0 {
		log.Fatalf("API URL for CanadaPost API cannot be empty. Check OS Env.")
	}
	if len(custnum) == 0 {
		log.Fatalf("Customer Number for CanadaPost API cannot be empty. Check OS Env.")
	}
	// Encode credentials to base64
	credentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	// Create XML for Request
	q := &cpRequest{XMLNS: "http://www.canadapost.ca/ws/ship/rate-v3", CustomerNumber: custnum}
	q.Weight = weight
	if weight < 0.3 {
		q.Weight = 0.3
	}
	q.ServiceCode = "DOM.EP"          // Service Code - Shipping service pre-selected
	q.OriginPostcode = originPostcode // Spaces and other chars should be removed from Postcodes before submission
	q.Destination.DestinationPostcode = destPostcode

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
	req.Header.Add("Authorization", "Basic "+credentials)
	req.Header.Add("Content-Type", "application/vnd.cpc.ship.rate-v3+xml")
	req.Header.Add("Accept", "application/vnd.cpc.ship.rate-v3+xml")
	req.Header.Add("Accept-Language", "en-CA")

	log.Println("Sending Rate Request to Canada Post API")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return cpRate, errors.New("Could not complete request to Canada Post API")
	}

	// Read XML Response from HTTP Body
	xmlResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Could not read HTTP response body. Error: %v", err)
		return cpRate, err
	}
	if resp.StatusCode != 200 {
		log.Printf("Canada Post API Error: HTTP Status %v, Body:\n%v", resp.StatusCode, string(xmlResp))
		return cpRate, errors.New("Canada Post API Response of Unexpected Type")
	}
	var Quote PriceQuotes
	err = xml.Unmarshal(xmlResp, &Quote)
	if err != nil {
		log.Printf("Body: %v", xmlResp)
		log.Printf("Could not unmarshall response body. Error: %v", err)
		return cpRate, err
	}
	log.Printf("CPost Quote - Due: %v ServiceName: %v Days to Delivery: %v", Quote.Due, Quote.ServiceName, Quote.ExpectedTransitTime)
	// Check expected response?
	if len(Quote.ServiceName) < 1 {
		responseDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatalf("Could not dump HTTP response. Error: %v", err)
		}
		log.Printf("Unexpected response from Canada Post API.\nResponse:\n%v", string(responseDump))
		return cpRate, err
	}

	// Process and return rate
	cpRate.Cost = Quote.Due
	cpRate.Description = Quote.ServiceName
	cpRate.GuaranteedDaysToDelivery = Quote.ExpectedTransitTime
	return cpRate, err
}

type PriceQuotes struct {
	OptionPrice          float64     `xml:"price-quote>price-details>options>option>option-price"`
	ServiceCode          string      `xml:"price-quote>service-code"`
	Xmlns                string      `xml:"xmlns,attr"`
	Gst                  Gst         `xml:"price-quote>price-details>taxes>gst"`
	Percent              []string    `xml:"price-quote>price-details>adjustments>adjustment>qualifier>percent"`
	WeightDetails        string      `xml:"price-quote>weight-details"`
	AmDelivery           string      `xml:"price-quote>service-standard>am-delivery"`
	GuaranteedDelivery   string      `xml:"price-quote>service-standard>guaranteed-delivery"`
	ServiceName          string      `xml:"price-quote>service-name"`
	Pst                  float64     `xml:"price-quote>price-details>taxes>pst"`
	OptionCode           string      `xml:"price-quote>price-details>options>option>option-code"`
	AdjustmentCode       []string    `xml:"price-quote>price-details>adjustments>adjustment>adjustment-code"`
	AdjustmentName       []string    `xml:"price-quote>price-details>adjustments>adjustment>adjustment-name"`
	AdjustmentCost       []float64   `xml:"price-quote>price-details>adjustments>adjustment>adjustment-cost"`
	ServiceLink          ServiceLink `xml:"price-quote>service-link"`
	Due                  float64     `xml:"price-quote>price-details>due"`
	OptionName           string      `xml:"price-quote>price-details>options>option>option-name"`
	ExpectedTransitTime  int         `xml:"price-quote>service-standard>expected-transit-time"`
	ExpectedDeliveryDate string      `xml:"price-quote>service-standard>expected-delivery-date"`
	Base                 float64     `xml:"price-quote>price-details>base"`
	Hst                  float64     `xml:"price-quote>price-details>taxes>hst"`
}

type Gst struct {
	Text    string  `xml:",chardata"`
	Percent float64 `xml:"percent,attr"`
}
type ServiceLink struct {
	Href      string `xml:"href,attr"`
	MediaType string `xml:"media-type,attr"`
	Rel       string `xml:"rel,attr"`
}
