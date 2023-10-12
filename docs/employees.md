# Mostrar todos los empleados

Mostrar todos los empleados de acuerdo al id de departamento asignado al administrador.

**URL** : `/api/accounts/:iddep/`

**URL Parameters** : `iddep=[uuid string]` id del departamento

**Method** : `GET`

**Auth required** : YES

## Success Response

**Condition** : Si el id del departamento es correcto

**Code** : `200 OK`

**Content example**

```json
{
  "body": [
    {
      "id_employee": "01fca326-f618-4a12-826a-9ff4c5adfbd9",
      "name": "Audrey",
      "position": "Maquetador",
      "salary": "$ 4882",
      "contrat_date": "2003-01-11",
      "id_dep": "428e976b-726c-4226-a77c-2ea601024ef8",
      "project": {
        "id_project": "",
        "name": "",
        "description": ""
      }
   
    }
  ]
}
```

## Error Responses

**Condition** : Si el parametro idep noes proporcionado.

**Code** : `404 NOT FOUND`

**Content** : `{"message": "parameter iddep incorrect"}`

### Or

**Condition** : iddep incorrecto o no existe 

**Code** : `200 OK`

**Content** :

```json
{
  "body": []
}
```

