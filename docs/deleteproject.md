# Eliminar proyecto

Eliminar proyecto con el id del empleado que viene siendo el id del proyecto

**URL** : `/api/deleteproject/:idpro`

**URL Parameters** : `idpro=[string]` id del proyecto a eliminar.

**Method** : `DELETE`

**Auth required** : YES

**Data** : `{}`

## Success Response

**Condition** : si el  proyecto existe 

**Code** : `202 ACCEPTED`

**Content** : `{"message": "project delete successfully"}`

## Error Responses

**Condition** : Si el id del proyecto es incorrecto o no existe.

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : `{message:error}`