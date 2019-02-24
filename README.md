# shopee
shopee
API Calculator Documentation
----------------------------

Table structure :
-----------------
transaction(
	billing_id SERIAL primary key,
	name varchar(125),
	tax_code smallint,
	price float
);
-> this api will save the four field above
--> name used to save the name of thing that you create.
--> tax_code used to categorize your creating thing.
--> price the number you have to paid.

API : 
-----
1) getmybill : 
URL : localhost:8080/getmybill
Param : -
Method : GET
Description :
-> this endpoint is aim to list all the product that you have created that consist of :
|
 --> detail transaction -> name, tax_code, price, and type.
|
 --> transaction calculation -> refundable, tax, and amount.
|
 --> the summary of the transaction -> price subtotal, tax subtotal, grand total.
    |
    ---->price subtotal => the sum of all price
    |
    ---->tax subtotla => the sum of all tax
    |
    ----> grand total => the sum of all amount

Example use : 
URL : localhost:8080/getmybill
Response:
{
    "StatusCode": 200,
    "data": {
        "detail_bills": [
            {
                "transaction": {
                    "name": "Lucky Stretch",
                    "tax_code": 2,
                    "type": "tobacco",
                    "price": 1000
                },
                "calc_detail": {
                    "refundable": "NO",
                    "tax": 30,
                    "amount": 1030
                }
            },
            {
                "transaction": {
                    "name": "Big Mac",
                    "tax_code": 1,
                    "type": "food",
                    "price": 1000
                },
                "calc_detail": {
                    "refundable": "YES",
                    "tax": 100,
                    "amount": 1100
                }
            },
            {
                "transaction": {
                    "name": "Movie",
                    "tax_code": 3,
                    "type": "entertainment",
                    "price": 150
                },
                "calc_detail": {
                    "refundable": "NO",
                    "tax": 0.5,
                    "amount": 150.5
                }
            }
        ],
        "summary": {
            "price_sub_total": 2150,
            "tax_sub_total": 130.5,
            "grand_total": 2280.5
        }
    },
    "Error": ""
}


2) addbill : 
URL : localhost:8080/addbill
Param : 
{
	"name":"xxx",
	"tax_code" : 1111,
	"price" : 111
}
Description : 
-> this enpoint aim to add transaction.

Example Use : 
	URL : localhost:8080/addbill
	{
		"name":"oreo",
		"tax_code" : 1,
		"price" : 1000
	}
	Response:
	{
	    "StatusCode": 200,
	    "data": "Success to Add Bill",
	    "Error": ""
	}
