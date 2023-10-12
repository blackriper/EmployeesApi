# EmpleoyeesApi

Api escrito en go para la autenticación de administradores para poder consultar, actualizar, eliminar empleados del área el cual administran. 

## Objetivo del Proyecto 
Implementar Api en go con autenticación [JWT](https://jwt.io/introduction) y usando los [Orm](https://www.dreams.es/transformacion-digital/desarrolladores-paginas-web/que-es-un-orm) [Gorm](https://gorm.io/docs/) y [Mgm](https://github.com/Kamva/mgm) para la interacción con las bases de datos MYSQL en el caso de los empleados y usando MongoDB para los administradores. 

## Clean Arquitecture
Clean architecture es un conjunto de principios cuya finalidad principal es ocultar los detalles de implementación a la lógica de dominio de la aplicación.
De esta manera mantenemos aislada la lógica, consiguiendo tener una lógica mucho más mantenible y escalable en el tiempo. 

Este proyecto está estructurado siguiendo estos principios, este tipo de arquitectura está destinada a **proyectos grandes** con cambios constantes de tecnologías este repositorio solo tiene el fin de mostrar un ejemplo de cómo implementar estos principios. 

![Clean arquitecture](/docs/clean.webp "clean arquitecture schema")

## Estructura de proyecto 

En el proyecto puedes encontrar las siguientes carpetas a continuación una breve descripción de cada carpeta  

- Controllers: Carpeta que contiene las funciones que se ejecutan al llamar alguna ruta de la Api. 

- Docs: Contiene la documentación del proyecto así como la de la Api.

- Domain: Contiene los modelos o entidades de la aplicación los cuales usa el orm para trabajar. 

- Mock : Código para generar un mock de fakedata para propósitos de desarrollo y testing 

- Repository: Contiene los diferentes servicios externos a la aplicación todo el código para conectarse e interactuar con la base de datos vive aquí. 

- Routes: Contine las ruta de la aplicación

- Test: Pruebas unitarias de los controllers puedes crear pruebas unitarias para los repository o usecases 

- Usecases: Contiene la lógica de la aplicación diferentes acciones de la Api 

- Utilities: Contiene funciones que se pueden usar en cualquier parte de la aplicación.


## Variables de entorno
Para poder correr el servidor local es necesario crear las siguientes variables de entorno las variables requeridas son las siguientes.

- MONGO_URI: cadena de conexión a mongoDb 
- PORT: puerto local donde puede ejecutarse el servidor de desarrollo 
- DNS: cadena de conexión de la base de datos en este caso mysql   


# Definicion de la Api
Esta Api está construida usando el framework gin para go y utiliza json responses para recibir y envió de datos. 

## Open Endpoint
Estos enpoint no necesita de autenticación 

* [Login](/docs/login.md) : `POST /login`

## Endpoint que requiere autenticación 
Los siguientes endpoints requiere de un token validado el cual debe estar en la cabezera de la solicitud al endpoint. el token se obtiene al hacer una solicitud al endpoint anteriormente mencionado.

## Enpoint de administrador
Estos endpoint le permite al administrador crear otro administrador, así como generar un nuevo token cuando el anterior expira.  

* [Newadmin](/docs/newadmin.md) : `POST auth/newadmin`
* [Refresh Token](/docs/refresh.md) : `POST auth/refresh_token`


## Enpoint de Empleados y Proyectos
Estos endpoint pueden manipular las diferentes operaciones para empleados y proyectos.

* [New Employee](/docs/newemployee.md) : `POST api/newemployee`
* [Employees](/docs/employees.md) : `GET api/employees/:iddep`
* [Update Employee](/docs/updateemployee.md) : `PUT api/updateemployee/:idemp`
* [Delete Employee](/docs/deleteemployee.md) : `DELETE api/deleteemployee/:idemp`
* [New Project](/docs/newproject.md) : `POST api/newproject`
* [Delete Project](/docs/deleteproject.md) : `DELETE api/deleteproject/:idpro`