package provider

import (
	"encoding/json"
	"log"
	"time"
)

func OpenOrder(data []byte) (Order, error) {
	// Decode Request Body JSON
	var o Order
	err := json.Unmarshal(data, &o)
	if err != nil {
		log.Printf("Could not decode/unmarshall order. Error: %v", err)
	}
	return o, err
}

type SnipcartQuote struct {
	Rates []Rate `json:"rates"`
}
type Rate struct {
	Cost                     float64 `json:"cost"`
	Description              string  `json:"description"`
	GuaranteedDaysToDelivery int     `json:"guaranteedDaysToDelivery,omitempty"`
}

type Order struct {
	EventName interface{} `json:"eventName"`
	Mode      string      `json:"mode"`
	CreatedOn time.Time   `json:"createdOn"`
	Content   struct {
		Token                     string      `json:"token"`
		Currency                  string      `json:"currency"`
		CreationDate              time.Time   `json:"creationDate"`
		ModificationDate          time.Time   `json:"modificationDate"`
		Status                    string      `json:"status"`
		PaymentStatus             interface{} `json:"paymentStatus"`
		Email                     string      `json:"email"`
		WillBePaidLater           bool        `json:"willBePaidLater"`
		BillingAddressFirstName   interface{} `json:"billingAddressFirstName"`
		BillingAddressName        string      `json:"billingAddressName"`
		BillingAddressCompanyName string      `json:"billingAddressCompanyName"`
		BillingAddressAddress1    string      `json:"billingAddressAddress1"`
		BillingAddressAddress2    string      `json:"billingAddressAddress2"`
		BillingAddressCity        string      `json:"billingAddressCity"`
		BillingAddressCountry     string      `json:"billingAddressCountry"`
		BillingAddressProvince    string      `json:"billingAddressProvince"`
		BillingAddressPostalCode  string      `json:"billingAddressPostalCode"`
		BillingAddressPhone       string      `json:"billingAddressPhone"`
		BillingAddress            struct {
			FullName    string      `json:"fullName"`
			FirstName   interface{} `json:"firstName"`
			Name        string      `json:"name"`
			Company     string      `json:"company"`
			Address1    string      `json:"address1"`
			Address2    string      `json:"address2"`
			FullAddress string      `json:"fullAddress"`
			City        string      `json:"city"`
			Country     string      `json:"country"`
			PostalCode  string      `json:"postalCode"`
			Province    string      `json:"province"`
			Phone       string      `json:"phone"`
		} `json:"billingAddress"`
		ShippingAddressFirstName   interface{} `json:"shippingAddressFirstName"`
		ShippingAddressName        string      `json:"shippingAddressName"`
		ShippingAddressCompanyName string      `json:"shippingAddressCompanyName"`
		ShippingAddressAddress1    string      `json:"shippingAddressAddress1"`
		ShippingAddressAddress2    string      `json:"shippingAddressAddress2"`
		ShippingAddressCity        string      `json:"shippingAddressCity"`
		ShippingAddressCountry     string      `json:"shippingAddressCountry"`
		ShippingAddressProvince    string      `json:"shippingAddressProvince"`
		ShippingAddressPostalCode  string      `json:"shippingAddressPostalCode"`
		ShippingAddressPhone       string      `json:"shippingAddressPhone"`
		ShippingAddress            struct {
			FullName    string      `json:"fullName"`
			FirstName   interface{} `json:"firstName"`
			Name        string      `json:"name"`
			Company     string      `json:"company"`
			Address1    string      `json:"address1"`
			Address2    string      `json:"address2"`
			FullAddress string      `json:"fullAddress"`
			City        string      `json:"city"`
			Country     string      `json:"country"`
			PostalCode  string      `json:"postalCode"`
			Province    string      `json:"province"`
			Phone       string      `json:"phone"`
		} `json:"shippingAddress"`
		ShippingAddressSameAsBilling bool        `json:"shippingAddressSameAsBilling"`
		CreditCardLast4Digits        interface{} `json:"creditCardLast4Digits"`
		TrackingNumber               interface{} `json:"trackingNumber"`
		TrackingURL                  interface{} `json:"trackingUrl"`
		ShippingFees                 interface{} `json:"shippingFees"`
		ShippingProvider             interface{} `json:"shippingProvider"`
		ShippingMethod               interface{} `json:"shippingMethod"`
		CardHolderName               interface{} `json:"cardHolderName"`
		PaymentMethod                interface{} `json:"paymentMethod"`
		Notes                        interface{} `json:"notes"`
		CustomFieldsJSON             string      `json:"customFieldsJson"`
		UserID                       interface{} `json:"userId"`
		CompletionDate               interface{} `json:"completionDate"`
		PaymentGatewayUsed           string      `json:"paymentGatewayUsed"`
		TaxProvider                  string      `json:"taxProvider"`
		Discounts                    []struct {
			AmountSaved                      float64       `json:"amountSaved"`
			DiscountID                       string        `json:"discountId"`
			ShippingDescription              interface{}   `json:"shippingDescription"`
			ShippingCost                     interface{}   `json:"shippingCost"`
			ShippingGuaranteedDaysToDelivery interface{}   `json:"shippingGuaranteedDaysToDelivery"`
			ID                               string        `json:"id"`
			Name                             string        `json:"name"`
			Combinable                       bool          `json:"combinable"`
			Trigger                          string        `json:"trigger"`
			Code                             string        `json:"code"`
			ItemID                           interface{}   `json:"itemId"`
			TotalToReach                     interface{}   `json:"totalToReach"`
			QuantityOfAProduct               interface{}   `json:"quantityOfAProduct"`
			QuantityOfProductIds             interface{}   `json:"quantityOfProductIds"`
			OnlyOnSameProducts               interface{}   `json:"onlyOnSameProducts"`
			QuantityInterval                 interface{}   `json:"quantityInterval"`
			MaxQuantityOfAProduct            interface{}   `json:"maxQuantityOfAProduct"`
			Type                             string        `json:"type"`
			Rate                             interface{}   `json:"rate"`
			Amount                           float64       `json:"amount"`
			ProductIds                       interface{}   `json:"productIds"`
			AlternatePrice                   interface{}   `json:"alternatePrice"`
			NumberOfItemsRequired            interface{}   `json:"numberOfItemsRequired"`
			NumberOfFreeItems                interface{}   `json:"numberOfFreeItems"`
			NumberOfUsages                   interface{}   `json:"numberOfUsages"`
			NumberOfUsagesUncompleted        interface{}   `json:"numberOfUsagesUncompleted"`
			AffectedItems                    []interface{} `json:"affectedItems"`
			DataAttribute                    interface{}   `json:"dataAttribute"`
			HasSavedAmount                   interface{}   `json:"hasSavedAmount"`
			Products                         []interface{} `json:"products"`
		} `json:"discounts"`
		Plans []interface{} `json:"plans"`
		Taxes []interface{} `json:"taxes"`
		User  interface{}   `json:"user"`
		Items []struct {
			Token            string      `json:"token"`
			Name             string      `json:"name"`
			Price            float64     `json:"price"`
			Quantity         int         `json:"quantity"`
			URL              string      `json:"url"`
			ID               string      `json:"id"`
			InitialData      string      `json:"initialData"`
			Description      string      `json:"description"`
			Weight           interface{} `json:"weight"`
			Image            string      `json:"image"`
			OriginalPrice    interface{} `json:"originalPrice"`
			UniqueID         string      `json:"uniqueId"`
			Stackable        bool        `json:"stackable"`
			MinQuantity      interface{} `json:"minQuantity"`
			MaxQuantity      interface{} `json:"maxQuantity"`
			AddedOn          time.Time   `json:"addedOn"`
			ModificationDate time.Time   `json:"modificationDate"`
			Shippable        bool        `json:"shippable"`
			Taxable          bool        `json:"taxable"`
			Duplicatable     bool        `json:"duplicatable"`
			Width            interface{} `json:"width"`
			Height           interface{} `json:"height"`
			Length           interface{} `json:"length"`
			TotalPrice       float64     `json:"totalPrice"`
			TotalWeight      float64     `json:"totalWeight"`
			Taxes            []string    `json:"taxes"`
			AlternatePrices  struct {
			} `json:"alternatePrices"`
			CustomFields []struct {
				Name         string      `json:"name"`
				Operation    interface{} `json:"operation"`
				Type         string      `json:"type"`
				Options      string      `json:"options"`
				Required     bool        `json:"required"`
				Value        string      `json:"value"`
				OptionsArray []string    `json:"optionsArray"`
			} `json:"customFields"`
			UnitPrice     float64 `json:"unitPrice"`
			HasDimensions bool    `json:"hasDimensions"`
		} `json:"items"`
		Refunds                 []interface{} `json:"refunds"`
		Lang                    string        `json:"lang"`
		RefundsAmount           float64       `json:"refundsAmount"`
		AdjustedAmount          float64       `json:"adjustedAmount"`
		FinalGrandTotal         float64       `json:"finalGrandTotal"`
		TotalNumberOfItems      int           `json:"totalNumberOfItems"`
		InvoiceNumber           string        `json:"invoiceNumber"`
		BillingAddressComplete  bool          `json:"billingAddressComplete"`
		ShippingAddressComplete bool          `json:"shippingAddressComplete"`
		ShippingMethodComplete  bool          `json:"shippingMethodComplete"`
		RebateAmount            float64       `json:"rebateAmount"`
		Subtotal                float64       `json:"subtotal"`
		ItemsTotal              float64       `json:"itemsTotal"`
		TaxableTotal            float64       `json:"taxableTotal"`
		GrandTotal              float64       `json:"grandTotal"`
		Total                   float64       `json:"total"`
		TotalWeight             float64       `json:"totalWeight"`
		TotalRebateRate         float64       `json:"totalRebateRate"`
		CustomFields            []interface{} `json:"customFields"`
		ShippingEnabled         bool          `json:"shippingEnabled"`
		NumberOfItemsInOrder    int           `json:"numberOfItemsInOrder"`
		PaymentTransactionID    string        `json:"paymentTransactionId"`
		Metadata                struct {
		} `json:"metadata"`
		TaxesTotal float64 `json:"taxesTotal"`
		ItemsCount int     `json:"itemsCount"`
		Summary    struct {
			Subtotal             float64       `json:"subtotal"`
			TaxableTotal         float64       `json:"taxableTotal"`
			Total                float64       `json:"total"`
			SummaryPaymentMethod interface{}   `json:"paymentMethod"`
			Taxes                []interface{} `json:"taxes"`
			AdjustedTotal        float64       `json:"adjustedTotal"`
		} `json:"summary"`
		IPAddress string `json:"ipAddress"`
	} `json:"content"`
}

var SnipcartTestJSON = []byte(`{
  "eventName": "shippingrates.fetch",
  "mode": "Live",
  "createdOn": "2015-02-21T14:58:02.6738454Z",
  "content": {
    "token": "22808196-0eff-4a6e-b136-3e4d628b3cf5",
    "creationDate": "2015-02-21T14:58:02.6738454Z",
    "modificationDate": "2015-02-21T14:58:02.6738454Z",
    "status": "Processed",
    "paymentMethod": "CreditCard",
    "email": "customer@snipcart.com",
    "cardHolderName": "Nicolas Cage",
    "billingAddressName": "Nicolas Cage",
    "billingAddressCompanyName": "Company name",
    "billingAddressAddress1": "888 The street",
    "billingAddressAddress2": "",
    "billingAddressCity": "QuÃ©bec",
    "billingAddressCountry": "CA",
    "billingAddressProvince": "QC",
    "billingAddressPostalCode": "G1G 1G1",
    "billingAddressPhone": "(888) 888-8888",
    "shippingAddressName": "Nicolas Cage",
    "shippingAddressCompanyName": "Company name",
    "shippingAddressAddress1": "888 The street",
    "shippingAddressAddress2": "",
    "shippingAddressCity": "QuÃ©bec",
    "shippingAddressCountry": "CA",
    "shippingAddressProvince": "QC",
    "shippingAddressPostalCode": "G1G 1G1",
    "shippingAddressPhone": "(888) 888-8888",
    "shippingAddressSameAsBilling": true,
    "finalGrandTotal": 310.00,
    "shippingAddressComplete": true,
    "creditCardLast4Digits": "4242",
    "shippingFees": 10.00,
    "shippingMethod": "Livraison",
    "items": [{
      "uniqueId": "eb4c9dae-e725-4dad-b7ae-a5e48097c831",
      "token": "22808196-0eff-4a6e-b136-3e4d628b3cf5",
      "id": "1",
      "name": "Movie",
      "price": 300.00,
      "originalPrice": 300.00,
      "quantity": 1,
      "url": "https://snipcart.com",
      "weight": 10.00,
      "description": "Something",
      "image": "http://placecage.com/50/50",
      "customFieldsJson": "[]",
      "stackable": true,
      "maxQuantity": null,
      "totalPrice": 300.0000,
      "totalWeight": 10.00
    }],
    "subtotal": 610.0000,
    "totalWeight": 20.00,
    "discounts": [],
    "willBePaidLater": false
  }
}`)
