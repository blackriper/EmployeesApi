# Refresh token 

Generar un nuevo token cuando el anterior ha expirado

**URL** : `/auth/refresh_token`


**Method** : `GET`

**Auth required** : YES

**Header Request**: `"Authorization:Bearer xxxxxxxxx"`

## Success Response

**Condition** : Si el usuario este logeado y tiene un token expirado valido 

**Code** : `200 OK`

**Content example**

```json
{
     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5NzQ0MzIsIm9yaWdfaWF0IjoxNjk2OTcwODMyfQ.I18e0P6nr8gQrKChr9lqtYz0dZIECF2UYAAH3Jnyyns",
    "expire": "2023-10-10T15:47:12-06:00"
}
```

## Error Responses

**Condition** : Si el token no es valido 

**Code** : ``401 UNAUTHORIZED``

**Content** :

```json
{"message": "token not valid"}
```
