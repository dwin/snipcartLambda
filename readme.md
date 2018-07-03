#### Snipcart Webhook - github.com/dsjr2006/snipcartWebhook
Created for originally AdoraStyle.ca for use with Snipcart webhook API running on AWS Lambda. 'POST' containing Snipcart JSON with "eventName":"shippingrates.fetch" to 
AWS endpoint, should return shipping rate quote as JSON, as specified by Snipcart API documentation.

Currently API queries Canada Post API with weight, origin postcode, and destination postcode for Expedited service, then returns that rate quote minus 
the discount set in OS Env.

Run `$ make` in directory to create package.zip file for upload to AWS Lambda Console.

__ToDo:__
- Unit Tests
- Improve Shipping Provider Error Handling
- Improve Snipcart Error Handling

_Endpoint:_ *(example only)  
https://xxxxxxxx.execute-api.us-east-1.amazonaws.com/prod/snipcartWebhook

_Accepts:_
* POST should return "Expected JSON Response to Snipcart" as shown below.
* GET should return "Nothing to see here. Your request works though!"

__AWS Lambda Deployment Config:__

* Runtime: Python 2.7
* Handler: handler.Handle
* Name: snipcartWebhook - by default the endpoint name is derived from this, see routes in `snipcartWebhook.go` to change
* API Gateway: Open
* Set Environmental Variables
  * `Ship_Discount`
  * `CAPost_CustNum`
  * `CAPost_USER`
  * `CAPost_PASS`
  * `CAPost_URL` Default:(https://soa-gw.canadapost.ca/rs/ship/price)
  * `ESHIP_USER`
  * `ESHIP_PASS`
  * `ESHIP_URL`





__Third-Party Libraries:__
* _aws-lambda-go-net_  
   Go shim for AWS Lambda which does not support Go as of 2017-02-15. https://github.com/eawsy/aws-lambda-go-net 
* _Gin_  
   Gin is a high-performance HTTP web framework for Go. https://github.com/gin-gonic/gin

__Expected JSON Response to Snipcart__
 ```
{"rates":[{"cost":10.1,"description":"Expedited Parcel","guaranteedDaysToDelivery":5}]}
 ```

__Error JSON Response to Snipcart__
 ```
{
  "errors": [{
    "key": "invalid_postal_code",
    "message": "The postal code is invalid."
    },
    ...
  ]
}
 ```

__Snipcart Shipping Fetch JSON:__
   ```
   {
  "eventName": "shippingrates.fetch",
  "mode": "Live",
  "createdOn": "2017-02-15T05:46:04.3278217Z",
  "content": {
    "token": "3ad41cba-43bc-42c1-8960-7162f922c6c3",
    "currency": "cad",
    "creationDate": "2017-02-15T05:43:52Z",
    "modificationDate": "2017-02-15T05:45:14Z",
    "status": "InProgress",
    "paymentStatus": null,
    "email": "test@testing.com",
    "willBePaidLater": false,
    "billingAddressFirstName": null,
    "billingAddressName": "John Doe",
    "billingAddressCompanyName": "Test",
    "billingAddressAddress1": "1293 Highridge Dr",
    "billingAddressAddress2": "",
    "billingAddressCity": "Kamloops",
    "billingAddressCountry": "CA",
    "billingAddressProvince": "BC",
    "billingAddressPostalCode": "V2C5G5",
    "billingAddressPhone": "5555555555",
    "billingAddress": {
      "fullName": "John Doe",
      "firstName": null,
      "name": "John Doe",
      "company": "Test",
      "address1": "1293 Highridge Dr",
      "address2": "",
      "fullAddress": "1293 Highridge Dr",
      "city": "Kamloops",
      "country": "CA",
      "postalCode": "V2C5G5",
      "province": "BC",
      "phone": "5555555555"
    },
    "shippingAddressFirstName": null,
    "shippingAddressName": "John Doe",
    "shippingAddressCompanyName": "Test",
    "shippingAddressAddress1": "1293 Highridge Dr",
    "shippingAddressAddress2": "",
    "shippingAddressCity": "Kamloops",
    "shippingAddressCountry": "CA",
    "shippingAddressProvince": "BC",
    "shippingAddressPostalCode": "V2C5G5",
    "shippingAddressPhone": "5555555555",
    "shippingAddress": {
      "fullName": "John Doe",
      "firstName": null,
      "name": "John Doe",
      "company": "Test",
      "address1": "1293 Highridge Dr",
      "address2": "",
      "fullAddress": "1293 Highridge Dr",
      "city": "Kamloops",
      "country": "CA",
      "postalCode": "V2C5G5",
      "province": "BC",
      "phone": "5555555555"
    },
    "shippingAddressSameAsBilling": true,
    "creditCardLast4Digits": null,
    "trackingNumber": null,
    "trackingUrl": null,
    "shippingFees": null,
    "shippingProvider": null,
    "shippingMethod": null,
    "cardHolderName": null,
    "paymentMethod": 0,
    "notes": null,
    "customFieldsJson": "[]",
    "userId": null,
    "completionDate": null,
    "paymentGatewayUsed": "None",
    "taxProvider": "Default",
    "discounts": [
      {
        "amountSaved": 24.5,
        "discountId": "46ea6470-8a0a-4904-ba7c-febf57aa82e3",
        "shippingDescription": null,
        "shippingCost": null,
        "shippingGuaranteedDaysToDelivery": null,
        "id": "44093ce1-35a6-4813-82ae-6876f8d3f269",
        "name": "test123456",
        "combinable": true,
        "trigger": "Code",
        "code": "test",
        "itemId": null,
        "totalToReach": null,
        "quantityOfAProduct": null,
        "quantityOfProductIds": null,
        "onlyOnSameProducts": false,
        "quantityInterval": false,
        "maxQuantityOfAProduct": null,
        "type": "FixedAmount",
        "rate": null,
        "amount": 24.5,
        "productIds": null,
        "alternatePrice": "",
        "numberOfItemsRequired": null,
        "numberOfFreeItems": null,
        "numberOfUsages": 0,
        "numberOfUsagesUncompleted": 0,
        "affectedItems": [],
        "dataAttribute": null,
        "hasSavedAmount": true,
        "products": []
      }
    ],
    "plans": [],
    "taxes": [],
    "user": null,
    "items": [
      {
        "token": "3ad41cba-43bc-42c1-8960-7162f922c6c3",
        "name": "Owl So Cute",
        "price": 12,
        "quantity": 1,
        "url": "https://www.adorastyle.ca/shop/kids",
        "id": "OWL01K",
        "initialData": "",
        "description": "Kids Leggings",
        "weight": null,
        "image": "/assets/img/products/cart/001k.jpg",
        "originalPrice": null,
        "uniqueId": "792876cf-368c-4571-bce1-c64a196644a4",
        "stackable": true,
        "minQuantity": null,
        "maxQuantity": null,
        "addedOn": "2017-02-15T05:43:52Z",
        "modificationDate": "2017-02-15T05:43:52Z",
        "shippable": true,
        "taxable": true,
        "duplicatable": false,
        "width": null,
        "height": null,
        "length": null,
        "totalPrice": 12,
        "totalWeight": 0,
        "taxes": [
          "GST (Kids)"
        ],
        "alternatePrices": {},
        "customFields": [
          {
            "name": "Size",
            "operation": null,
            "type": "dropdown",
            "options": "Small (Age 2-5)|Large (Age 6-8)",
            "required": false,
            "value": "Small (Age 2-5)",
            "optionsArray": [
              "Small (Age 2-5)",
              "Large (Age 6-8)"
            ]
          }
        ],
        "unitPrice": 12,
        "hasDimensions": false
      }
    ],
    "refunds": [],
    "lang": "en",
    "refundsAmount": 0,
    "adjustedAmount": 0,
    "finalGrandTotal": 0,
    "totalNumberOfItems": 0,
    "invoiceNumber": "3ad41cba-43bc-42c1-8960-7162f922c6c3",
    "billingAddressComplete": true,
    "shippingAddressComplete": true,
    "shippingMethodComplete": false,
    "rebateAmount": 24.5,
    "subtotal": 0,
    "itemsTotal": 12,
    "taxableTotal": 12,
    "grandTotal": 0,
    "total": 0,
    "totalWeight": 0,
    "totalRebateRate": 0,
    "customFields": [],
    "shippingEnabled": true,
    "numberOfItemsInOrder": 1,
    "paymentTransactionId": "",
    "metadata": {},
    "taxesTotal": 0,
    "itemsCount": 1,
    "summary": {
      "subtotal": 0,
      "taxableTotal": 12,
      "total": 0,
      "paymentMethod": 0,
      "taxes": [],
      "adjustedTotal": 0
    },
    "ipAddress": "68.00.150.101"
  }
}
```
