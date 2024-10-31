# fetch_receipt_processor
Fetch Take Home Assessment

# dependencies
- go
- docker

# running the api
1. `go mod tidy`
2. `cd /source`
3. `go run .`

# POST
`curl -X POST http://localhost:8080/receipts/process    -H 'Content-Type: application/json'    -d '{
     "retailer": "Target",
     "purchaseDate": "2022-01-01",
     "purchaseTime": "13:01",
     "items": [
       {
         "shortDescription": "Mountain Dew 12PK",
         "price": "6.49"
       },
       {
         "shortDescription": "Emils Cheese Pizza",
         "price": "12.25"
       },
       {
         "shortDescription": "Knorr Creamy Chicken",
         "price": "1.26"
       },
       {
         "shortDescription": "Doritos Nacho Cheese",
         "price": "3.35"
       },
       {
         "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
         "price": "12.00"
       }
     ],
     "total": "35.35"
   }'`

   `curl -X POST http://localhost:8080/receipts/process \
   -H 'Content-Type: application/json' \
   -d '{
     "retailer": "Target",
     "purchaseDate": "2022-01-01",
     "purchaseTime": "13:01",
     "items": [
       {
         "shortDescription": "Mountain Dew 12PK",
         "price": "6.49"
       }
     ],
     "total": "6.49"
   }'
`

# GET
`curl http://localhost:8080/receipts/<id>/points`