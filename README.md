# Golang microservices assignment

## Overview

The evaluation result of the test is not linked to how much time you spend on it. You should not spend more than 2h although there are no hard time limits.

This assignment is meant to evaluate the golang proficiency of full-time engineers.
Your code structure should follow microservices best practices and our evaluation will focus primarily on your ability to follow good design principles and less on the correctness and completeness of algorithms. During the face-to-face interview, you will have the opportunity to explain your design choices and provide justifications for the parts that you omitted.

## Evaluation criteria

- clean & self-explanatory code 
- goland idiomatic code
- testing 
- code quality checks linters & build tools
- git and useful commit messages
- you MUST use go modules 
- Bonus points: demonstrate knowledge of containers (Kubernetes or docker)
- Bonus points: documentation


Avoid using frameworks one of the goals is to see how you structure applications. 

Results: Please share a git repository with us containing your implementation.

## Technical test

- implement an http service that will listen for a webhook at the url `/webhooks/transaction`
- decode the payload in json and store it in a database 
- there are no guarantees that the same transaction will arrive only once
- the service needs to ensure that we will store the request exactly once
- the calling service will retry the operation several times until it receives back an 2xx
- the system must expose an endpoint `/account/{id}` which will return the balance of the account
- transaction types of SALE will debit the account
- transaction types of CREDIT or REFUND will credit the account

### Webhook Payload 

Schema

- **accountId**: is a unique identifier of the account
- **transactionId**: uniquely identifies the transaction
- **orderId**: identifies a group of transactions, a SALE and refund would have the same orderId
- **transactionType**: can be SALE, CREDIT or REFUND
- **amount**: is the transaction value in the currency specified
- **currency**: a 3 letter code like EUR, USD or GBP
- **description**: a short text description explaining the transaction


Example: 

```json
{
    "transactionId": "tqZi6QapS41zcEHy",
    "transactionType": "SALE",
    "orderId": "c66oxMaisTwJQXjD",
    "amount": "10.00",
    "currency": "EUR",
    "description": "Test transaction",
    "accountId": "001"
}
```


### Example request for testing

```bash
curl 'http://localhost:8080/webhooks/transaction' \
    -H 'content-type: application/json' \
    --data-raw '{"transactionId":"tqZi6QapS41zcEHy","orderId":"c66oxMaisTwJQXjD", "transactionType":"SALE", "amount": "10.00", "currency":"EUR", "description":"Test transaction", "accountId":"001"}'
```

