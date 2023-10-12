# Registrar nuevo proyecto

Crear un proyecto a un empleado por efectos de aprendizaje un empleado solo puede tener un proyecto y el id de este tiene que ser el id del empleado a cuál es asignado. 

**URL** : `/api/project`

**Method** : `POST`

**Auth required** : YES


**Data constraints**
Es necesario proporcionar los siguientes campos para crear correctamente un nuevo proyecto.


```json
   {     
        "id_project": "uuid string",
        "name": "string",
        "description": "string"
     }
```

**Data example** Todos los campos son requeridos.

```json
{
    "id_project": "1657524e-d38c-499e-99d6-410503ec9153",
    "name": "Backend go",
    "description": "backend go para consultas"
}
```

## Success Response

**Condition** : si el id de projecto no existe y tiene los campos obligatorios

**Code** : `201 CREATED`

**Content example**

```json
{
   "message": "proyect created successfully"
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

**Condition** : Si el  id de proyecto ya esta asignado a un empleado 


**Code** : `500 INTERNAL SERVER ERROR`

**Content example**

```json
{
     "message": "error"   
	
}
```