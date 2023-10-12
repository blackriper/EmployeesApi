# Eliminar empleado

Eliminar empleados de acuerdo al id del departamento que tenga el administrador 

**URL** : `/api/deleteemployee/:idemp`

**URL Parameters** : `idemp=[string]` id del empleado a eliminar.

**Method** : `DELETE`

**Auth required** : YES

**Data** : `{}`

## Success Response

**Condition** : si el empleado existe

**Code** : `202 ACCEPTED`

**Content** : `{"message": "employee delete successfully"}`

## Error Responses

**Condition** : Si el id del empleado es incorrecto o no existe.

**Code** : `500 INTERNAL SERVER ERROR`

**Content** : `{message:error}`


