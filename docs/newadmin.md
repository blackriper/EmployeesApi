# Crear Nuevo Administrador 

Crear un nuevo administrador para un departamento este solo podrá hacer modificaciones del departamento del cual sea administrado. 

**URL** : `/auth/newadmin`

**Method** : `POST`

**Auth required** : YES


**Data constraints**
Es necesario proporcionar los siguientes campos para crear correctamente un nuevo administrador.


```json
{
    "email":"string",
    "user_name":" string",
    "password":"string",
    "departament_id":" uuid string"
}
```

**Data example** Todos los campos son requeridos.

```json
{
        "email":          "blackriper@gmail.com",
		"user_name":      "blackriper",
		"password":       "blackriper123",
		"departament_id": "c7f9f37a-72bb-46f0-bc43-09983adea143"
}
```

## Success Response

**Condition** : si el administrador no existe  y todos los campos son correctos

**Code** : `201 CREATED`

**Content example**

```json
{
   "message": "Admin created successfully"
}
```

## Error Responses

**Condition** : Si el administrador existe 

**Code** : ``400 BAD REQUEST``
**Content example**

```json
{
    "message": "admin already exists"
}
```


### Or

**Condition** : Si algún campo no está en el cuerpo de la solicitud. 


**Code** : `400 BAD REQUEST`

**Content example**

```json
{
        //error al no estar el campo email
		"user_name":      "blackriper",
		"password":      "blackriper123",
		"departament_id": "c7f9f37a-72bb-46f0-bc43-09983adea143"
	
}
```