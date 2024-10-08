# Backend de Autenticación de Usuarios (Golang)

Este repositorio contiene la implementación de un microservicio de autenticación de usuarios utilizando Firebase basado en el codigo Moises, con los cambios necesarios para guardarse en la base de datos supabase. 

## Funcionalidades Implementadas

- **/login**: Permite a los usuarios iniciar sesión utilizando Firebase Auth y devuelve un token JWT válido.
- **/register**: Registra nuevos usuarios utilizando Firebase Auth y guarda información adicional en supabase.
- **/verify**: Verifica códigos de verificación enviados por correo electrónico durante el registro.
- **/forgotPassword**: Permite a los usuarios solicitar recuperación de contraseña mediante correo electrónico.
- **/updateUser**: Permite a los usuarios actualizar su información personal.
- **/uploadPhoto**: Permite a los usuarios cargar y actualizar su foto de perfil.
- **/validateToken**: Valida tokens JWT y devuelve la información del usuario.

## Requisitos

- Implementación utilizando Firebase Auth para la autenticación.
- Almacenamiento de datos adicionales en supabase.
- Envío de correos electrónicos utilizando SMTP para verificaciones y recuperación de contraseña.
- Documentación de los endpoints utilizando Swagger.
- Manejo de errores y respuestas consistentes.

## Estado Actual

- Implementación inicial completada con integración de Firebase Auth y SMTP.
- Funcionalidades básicas como registro, verificación, actualización de usuario y carga de fotos de perfil están completamente funcionales.
- Documentación Swagger disponible localmente para consulta.

## Próximos Pasos

- Terminar la conexion con la base de datos y comprobar buen funcinamiento.
