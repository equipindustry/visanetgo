package test

// TokenResponse string.
var TokenResponse = `eyJraWQiOiJmWk1tV3pZR0RBckxHektvalNCK2w3SjFhMnNPXC9zQnNwOTlNNmNuM3F5MD0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJjNTAxNjYyOS04Zjc2LTQ1M2QtYjhlNC01MGJjZDI5YjI2NTAiLCJldmVudF9pZCI6IjYyYTM2MjBmLWVlNzItNDdkMy1hODdjLWRmMmQ3ZjI0YWU0MSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1OTQ4MzQ3MjksImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy1lYXN0LTEuYW1hem9uYXdzLmNvbVwvdXMtZWFzdC0xXzJjSjFTZTFmSSIsImV4cCI6MTU5NDgzODMyOSwiaWF0IjoxNTk0ODM0NzI5LCJqdGkiOiIyZDA2ZDU5Ni0xMGRmLTRmYzQtODA1Yi0zZjEyNzk4NTRlNGIiLCJjbGllbnRfaWQiOiIxMGx2MDYxN281ZGljNTFlYnNucWVpaWpiNyIsInVzZXJuYW1lIjoiaW50ZWdyYWNpb25lcy52aXNhbmV0QG5lY29tcGx1cy5jb20ifQ.HKj00LRJVtUv3kUHu83JWCTdHjMcAtzzJEMEX7aXNkr-0cQ5d3ML_RLn5bqhK44S8VKCRBUZzY-eCiBllXVPicTxdmhHIg4GbkQwpKGhHlhGpkQpRKsNmTO1xQ3IkaSHKEkl1GdngPdtet0rYHTefy16xJXrluREizpNepI-BYY4-KVVcdZpDIbNg0r5xiXlzaQ4dPagNfqT6XpeJz3dHcRhQ74NKOtl0HqkMZAWx_Qj6zZjddqQXi9-HcLy9Q3FG3PghlQCv-qNHKki0FpT1nVH6FMeQtkGNSa5AJ_SSqbohrWs3-ZtK-AijqyVQqRPRvkCfRWcKEEuj-qOgyag3w`

// TokenFailResponse string.
var TokenFailResponse = `Unauthorized access`

// SessionResponse string.
var SessionResponse = `{
  "sessionKey": "e242ccbf2e4ab1bbd705e20cbebf9bd3df739e7a0239a14294b300d83995b092",
  "expirationTime": 1594910096318
}`

// SessionResponseFailTokenUsed string.
var SessionResponseFailTokenUsed = `{
  "errorCode": 400,
  "errorMessage": "Token has been used before",
  "data": {}
}`

// SessionResponseFailSendData string.
var SessionResponseFailSendData = `{
  "errorCode": 400,
  "errorMessage": "%s is not valid.",
  "data": {}
}`

// AuthorizationResponse string.
var AuthorizationResponse = `{
  "header": {
    "ecoreTransactionUUID": "3bfe5a37-570c-49cf-bebd-588d7066a33a",
    "ecoreTransactionDate": 1522528900672,
    "millis": 3867
  },
  "order": {
    "tokenId": "44BDC4D1500F4927BDC4D1500F7927D6",
    "purchaseNumber": "8078",
    "productId": "321",
    "amount": 147.02,
    "currency": "PEN",
    "authorizedAmount": 147.02,
    "authorizationCode": "153831",
    "actionCode": "000",
    "traceNumber": "3713",
    "transactionDate": "180331154140",
    "transactionId": "991180900182558"
  },
  "dataMap": {
    "CURRENCY": "0604",
    "TRANSACTION_DATE": "180331154140",
    "TERMINAL": "00000001",
    "ACTION_CODE": "000",
    "TRACE_NUMBER": "3713",
    "ECI_DESCRIPTION": "Transaccion no autenticada pero enviada en canal seguro",
    "ECI": "07",
    "BRAND": "visa",
    "CARD": "402160******4513",
    "MERCHANT": "341198210",
    "STATUS": "Verified",
    "ADQUIRENTE": "570002",
    "ACTION_DESCRIPTION": "Aprobado y completado con exito",
    "ID_UNICO": "991180900182558",
    "AMOUNT": "147.02",
    "PROCESS_CODE": "000000",
    "RECURRENCE_STATUS": "Registered",
    "TRANSACTION_ID": "991180900182558",
    "AUTHORIZATION_CODE": "153831"
  }
}`

// AuthorizationFailResponse string.
var AuthorizationFailResponse = `{
  "errorCode": 400,
  "errorMessage": "Not Authorized",
  "header": {
    "ecoreTransactionUUID": "3bfe5a37-570c-49cf-bebd-588d7066a33a",
    "ecoreTransactionDate": 1522528900672,
    "millis": 3867
  },
  "data": {
    "CURRENCY": "0604",
    "TRANSACTION_DATE": "180331152444",
    "TERMINAL": "00000001",
    "ACTION_CODE": "180",
    "TRACE_NUMBER": "3711",
    "ECI_DESCRIPTION": "Transaccion no autenticada pero enviada en canal seguro",
    "ECI": "07",
    "CARD": "400310****6160",
    "BRAND": "visa",
    "MERCHANT": "341198210",
    "STATUS": "Not Authorized",
    "ADQUIRENTE": "570002",
    "ACTION_DESCRIPTION": "Tarjeta no valida",
    "ID_UNICO": "991180900182558",
    "AMOUNT": "125.34",
    "PROCESS_CODE": "000000",
    "TRANSACTION_ID": "991180900182558"
  }
}`