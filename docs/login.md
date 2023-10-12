# Login

Verificar credenciales del usuario para autenticación 

**URL** : `/api/login/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "username": "[valid email address]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "email":"blackriper@gmail.com",
    "password":"blackriper123"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
  "message": {
    "user_name": "blackriper",
    "departament_id": "b5a5988c-ddf5-4ffc-bfbb-3cb8524c8475",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5Njk2MjMsIm9yaWdfaWF0IjoxNjk2OTY2MDIzfQ.gEh4p_LLhWlUcmnWtHAK1f6fhr9cb2yuoMH38RBsTZU",
    "expire": "2023-10-10T14:27:03-06:00"
  }
```

## Error Responses

**Condition** : si faltan el usuario y contraseña.

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "code": 401,
  "message": "missing Username or Password"
}
```

**Condition** : Cuando el Admin no existe o credenciales incorrectas

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "code": 401,
  "message": "admin user not exist"
}
```