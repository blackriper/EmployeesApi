# Registrar nuevo empleado

Crear un nuevo empleado del área a cargo del administrador de forma opcional se le puede agregar un proyecto al empleado. 

**URL** : `/api/newemployee`

**Method** : `POST`

**Auth required** : YES


**Data constraints**
Es necesario proporcionar los siguientes campos para crear correctamente un nuevo empleado.


```json
{
      "id_employee": "string uuid",
      "name": "string",
      "position": "string",
      "salary": "string",
      "contrat_date": "string",
      "id_dep": "string uuid",
      "project": {
        "id_project":"string uuid opcional",
        "name": "string opcional",
        "description": "string opcional",
       }
    }
```

**Data example** Todos los campos son requeridos.

```json
{
      "id_employee": "0fc4c3b5-f701-4067-86e9-a26b856c49d1",
      "name": "Hector",
      "position": "Programador",
      "salary": "$ 12,251",
      "contrat_date": "2023-09-20",
      "id_dep": "b5a5988c-ddf5-4ffc-bfbb-3cb8524c8475",     
}
```

## Success Response

**Condition** : si el id de empleado no existe y tiene los campos obligatorios

**Code** : `201 CREATED`

**Content example**

```json
{
   "message": "employee created successfully"
}
```

## Error Responses

**Condition** : Si algun campo no esta en la solicitud

**Code** : ``400 BAD REQUEST``
**Content example**

```json
{
    "message": "error"
}
```


### Or

**Condition** : Si el  id de empleado ya existe  


**Code** : `500 INTERNAL SERVER ERROR`

**Content example**

```json
{
     "message": "error"   
	
}
```