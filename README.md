# Prëxis API

## Authentication

    POST https://prexis.io/api/v1/auth

Returns the JWT token.

**Parameters:**

| Name        | Value                                                          |
| ----------- | -------------------------------------------------------------- |
| `key`       | The client key provided by Prëxis                              |
| `password`  | Password provided by Prëxis                                    |


**Response example:**
```javascript
{
  "code": 200,
  "status": "success",
  "message": "Valid login credentials",
  "key": "<key>",
  "jwt": "<jwt token>",
  "cached": "true"
}
```

**Important:**

* The jwt returned string is used to make the API requests;
* The jwt has to be used after one second after the response to avoid service misuse;
* These token last for 60 seconds, after that, it's necessary to renew it;
* The error number uses an internal code.

