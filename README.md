# Prëxis API

## Authentication

    POST https://prexis.io/api/v1/auth

Returns the JWT token.

**Parameters:**

| Parameter Name | Parameter Value                                                |
| -------------- | -------------------------------------------------------------- |
| `key`          | The client key provided by Prëxis                              |
| `password`     | Password provided by Prëxis                                    |


**Response example:**
```javascript
{
  "code": 200,
  "status": "success",
  "message": "Valid login credentials",
  "key": "<client key>",
  "jwt": "<jwt token>",
  "cached": "true"
}
```

**Important:**

* The jwt returned string is used to make the API requests;
* The jwt has to be used after one second after the response to avoid service misuse;
* These token last for 60 seconds, after that, it's necessary to renew it;
* The error number uses an internal code.


### How to use JWT token

The token must be in the header of every request.

**Example:**

    curl --header "Authorization: Bearer <jwt token>" http://prexis.io/api/v1/credits
    
Please notice the space between the `Bearer` string and the `jwt token`.


## Credit Balance

    GET https://prexis.io/api/v1/credits

Returns the credit balance.

**Parameters:**

None.

**Header:**

| Header Name     | Header Value                                                   |
| --------------- | -------------------------------------------------------------- |
| `Authorization` | `Bearer <jwt token>`                                           |


**Response example:**
```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535980547,
  "result": {
    "rescode": 200,
    "message": "Operation successfull",
    "credits": 10
  }
}
```

## Registering a Hash

    POST https://prexis.io/api/v1/register

Registers a valid SHA256 hash in Ethereum blockchain.

**Parameters:**

| Parameter Name | Parameter Value                                                |
| -------------- | -------------------------------------------------------------- |
| `hash`         | A valid SHA256 hash string without the `0x` prefix             |


**Header:**

| Header Name     | Header Value                                                   |
| --------------- | -------------------------------------------------------------- |
| `Authorization` | `Bearer <jwt token>`                                           |


**Response examples:**

Successfull request:

```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535980723,
  "result": {
    "credits": 9,
    "rescode": 200,
    "message": "Operation successfull",
    "hash": "6aaf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387"
  }
}
```

Hash already registered:

```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535980752,
  "result": {
    "rescode": 402,
    "message": "Hash already registered",
    "hash": "6aaf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387"
  }
}
```

Invalid SHA256 hash format:

```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535980784,
  "result": {
    "rescode": 401,
    "message": "Invalid SHA256 hash format",
    "hash": "<invalid hash>"
  }
}
```

## Get Blockchain Transaction

    POST https://prexis.io/api/v1/transaction

Registers a valid SHA256 hash in Ethereum blockchain.

**Parameters:**

| Parameter Name | Parameter Value                                                |
| -------------- | -------------------------------------------------------------- |
| `hash`         | A valid SHA256 hash string without the `0x` prefix             |


**Header:**

| Header Name     | Header Value                                                   |
| --------------- | -------------------------------------------------------------- |
| `Authorization` | `Bearer <jwt token>`                                           |


**Response examples:**

Transaction already mined in a block:

```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535981729,
  "result": {
    "transaction": {
      "blockHash": "0xa74ef377e2b8f889390c6c2168c5dbd26533741123d3f0ad097607538fb6a14e",
      "blockNumber": 3965065,
      "blockchainUrl": "https://ropsten.etherscan.io/tx/0x3d0da45eeea6e5acf29a79b16d34dd565709a3874646f50ec60339345aead806",
      "cumulativeGasUsed": 10538243,
      "from": "0xd8f7e3adb3c93739b2600e7180d2dfaecfa3b4e9",
      "gasUsed": 23176,
      "timestamp": 1535980696,
      "transactionHash": "0x3d0da45eeea6e5acf29a79b16d34dd565709a3874646f50ec60339345aead806",
      "transactionIndex": 26
    },
    "rescode": 200,
    "message": "Operation successfull",
    "hash": "6aaf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387"
  }
}
```

Transaction pending (waiting for be mined):

```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535981729,
  "result": {
    "transaction": {
      "blockchainUrl": "https://ropsten.etherscan.io/tx/0x3d0da45eeea6e5acf29a79b16d34dd565709a3874646f50ec60339345aead806",
      "transactionHash": "0x3d0da45eeea6e5acf29a79b16d34dd565709a3874646f50ec60339345aead806",
    },
    "rescode": 200,
    "message": "Operation successfull",
    "hash": "6aaf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387"
  }
}
```

Transaction not yet submitted:

```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535981729,
  "result": {
    "transaction": {
      "transactionHash": "wait",
    },
    "rescode": 200,
    "message": "Operation successfull",
    "hash": "6aaf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387"
  }
}
```

Invalid SHA256 hash format:

```javascript
{
  "code": 200,
  "status": "success",
  "key": "<client key>",
  "current_time": 1535981729,
  "result": {
    "rescode": 401,
    "message": "Invalid SHA256 hash format",
    "hash": "<invalid hash>"
  }
}
```

Hash not registered with Prëxis:

```javascript
{
  "code": 200,
  "status": "success",
  "key": "cK19vqzgqB2",
  "current_time": 1535981729,
  "result": {
    "rescode": 403,
    "message": "Hash not registered",
    "hash": "6baf8f326d9d27d212e8647cbd1306dc1687f90595d0e6821e274d4d6312c387"
  }
}
```
