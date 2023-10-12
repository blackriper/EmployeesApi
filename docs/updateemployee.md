# Actualizar empleado y proyeco

Actualizar datos del empleado asi como del proyecto(opcional) que puede tener a cargo 

**URL** : `/api/updateemployee/:idemp`

**Method** : `PUT`

**Auth required** : YES

**URL Parameters** : `iddemp=[uuid string]` id del empleado

**Data constraints**

```json
{
      "id_employee": " string uuid",
      "name": "string",
      "position": "string",
      "salary": "string ",
      "contrat_date": "string",
      "id_dep": "string uuid",
      "project": {
        "id_project":"string uuid id de empleado",
        "name": "string",
        "description": "string"
       }
    
}
```

**Data example** 
Puedes incluir un proyecto si hay que actulizar los datos del mismo.

```json
{
      "id_employee": "0fc4c3b5-f701-4067-86e9-a26b856c49d1",
      "name": "Hector",
      "position": "Programador",
      "salary": "$ 14,251",
      "contrat_date": "2023-09-20",
      "id_dep": "b5a5988c-ddf5-4ffc-bfbb-3cb8524c8475",
      "project": {
        "id_project":"0fc4c3b5-f701-4067-86e9-a26b856c49d1",
        "name": "App IOS ",
        "description": "app de IOS para renta de departamentos"
       }
    }
```

## Success Responses

**Condition** : Datos correctamente proporcionados 

**Code** : `202 ACEPTED`

**Content example** 

```json
{
  "message": "employee updated successfully"
   
}
```

## Error Response

**Condition** : Datos de solicitud incorrectos 

**Code** : `400 BAD REQUEST`

**Content** : `{message:error}`

### Or

**Condition** : Usuario no existe o algun error al actualizar

**Code** : `500 ITERNAL SERVER ERROR`

**Content** : `{message:error}`

