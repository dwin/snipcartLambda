package main

// /* snipcartWebhook.go - Darwin Smith 2017 */

// /* Required, but no C code needed. */
import "C"

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/dsjr2006/snipcartLambda/provider"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

func init() {

	// Set Proxy Event Listener
	listener := net.Listen()

	// Amazon API Gateway Binary support out of the box.
	Handle = apigatewayproxy.New(listener, nil).Handle

	// New Gin Router
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// Gin Routes, all others return '404' Not Found
	r.POST("/snipcartWebhook", handle)
	r.GET("/snipcartWebhook", show)

	go http.Serve(listener, r)
}
func main() {

}
func handle(c *gin.Context) {
	// Read Request Body
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("Cannot read request body. Error: %v", err)
	}
	// Open/Unpack Order JSON
	order, err := provider.OpenOrder(reqBody)
	if err != nil {
		log.Fatalf("Unable to unpack JSON received as Snipcart Order. Error: %v", err)
		c.String(400, "Not in Snipcart Order Format")
		return
	}

	// Check Order Event Type
	if order.EventName.(string) != "shippingrates.fetch" {
		log.Println("Request does not contain 'eventName:shippingrates.Fetch'")
		c.String(400, "Accepts Shipping Order Events Only.")
		return
	}
	log.Printf("Request Received for Invoice number %v", order.Content.InvoiceNumber)
	// Remove spaces and hyphen? from shippingAddressPostalCode
	originPostcode := spacemap("K1J9H7") // Origin Postcode set here
	destinationPostcode := spacemap(order.Content.ShippingAddressPostalCode)
	weight := float64(order.Content.TotalWeight)
	// Weight must be at least 0.25
	if weight > 0.25 {
		weight = 0.25
	}
	// Check Destination Country
	if order.Content.ShippingAddressCountry != "CA" {
		log.Printf("\nAttempting Canada Post Request, country is not Canada. Country: %v", order.Content.ShippingAddressCountry)
		snipErr := new(provider.SnipcartShipError)
		e := provider.ShipError{Key: "invalid_destination_country", Message: "Contact Support, Canada Post setup for Canada destinations only."}
		snipErr.Errors = append(snipErr.Errors, e)
		c.IndentedJSON(200, snipErr)
	}
	// Get Shipping Quote from Canada Post API
	log.Printf("Received- Weight: %v Origin: %v Dest: %v", weight, originPostcode, destinationPostcode)
	cPostRateQuote, err := provider.GetCanadaPostRate(weight, originPostcode, destinationPostcode)
	if err != nil {
		log.Printf("Invalid Response from Canada Post API. Error: %v", err)
	}

	shipDiscStr := os.Getenv("Ship_Discount") // Shipping discount must be set as OS Env Var ex. $ export Ship_Discount=5.00
	if len(shipDiscStr) == 0 {
		log.Fatal("Shipping discount cannot be empty. Check OS Env.")
	}
	shipDisc, err := strconv.ParseFloat(shipDiscStr, 64)
	if err != nil {
		log.Fatal("Could not parse shippping discount from Env. ex. 5.00")
	}
	// Check discount not greater than total shipping rate
	if cPostRateQuote.Cost > shipDisc {
		log.Printf("Discount given: Original Cost %v - Disc Cost %v", cPostRateQuote.Cost, (cPostRateQuote.Cost - shipDisc))
		cPostRateQuote.Cost = cPostRateQuote.Cost - shipDisc
	} else {
		log.Printf("Discount not applied, would exceed shipping cost of %v", cPostRateQuote.Cost)
	}
	q := new(provider.SnipcartQuote)
	q.Rates = append(q.Rates, cPostRateQuote)
	if err != nil {
		log.Printf("Error: %v", err)
		c.Status(500)
	}
	c.JSON(200, q)
	return
}
func show(c *gin.Context) {
	c.String(200, "Nothing to see here. Your request works though!")
	return
}

// Removes spaces from strings
func spacemap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
