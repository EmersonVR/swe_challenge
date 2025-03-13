Descripción del Proyecto
Este proyecto consiste en un sistema que obtiene datos de acciones desde una API, los almacena en CockroachDB y proporciona una interfaz para visualizar la información y recomendar las mejores opciones de inversión.

El sistema se compone de:
Backend en Golang: API REST para gestionar las acciones y realizar recomendaciones.
Base de datos en CockroachDB: Almacena la información de las acciones obtenidas desde la API.
Frontend en Vue 3 con TypeScript y Pinia: UI amigable para la visualización de datos.
Estilos con Tailwind CSS: Diseño moderno y optimizado para mejorar la experiencia del usuario.

Características Principales

Conexión a la API de Acciones

Obtiene datos de acciones desde la API externa de manera paginada.
Maneja errores y validaciones en la respuesta de la API.
Almacena la información en CockroachDB.

API REST en Golang

Endpoint para listar todas las acciones almacenadas.
Endpoint para obtener las mejores recomendaciones de inversión según un algoritmo de puntuación.

Interfaz de Usuario en Vue 3

Tabla interactiva con filtros y paginación.
Visualización clara de la información de las acciones.

Algoritmo de Recomendación

Analiza el crecimiento de la acción y la calificación del corretaje.
Asigna un puntaje basado en tendencias del mercado.
Retorna las 5 mejores opciones de inversión.

Pruebas Unitarias

Pruebas de integración con base de datos usando sqlmock.
Pruebas de endpoints con httptest.
Validación de lógica de recomendación.

### Despliegue del Proyecto con Terraform  

Para desplegar el proyecto en AWS, se utiliza Terraform para crear la infraestructura necesaria, incluyendo una instancia EC2 para el backend y una base de datos en RDS con CockroachDB.  

Pasos para desplegar la infraestructura

1. Inicializar Terraform
   
   terraform init

2. Aplicar la configuración y desplegar los recursos

   terraform apply -auto-approve

3. Obtener la dirección IP pública de la instancia EC2
   
   terraform output ec2_instance_public_ip


4. Acceder a la instancia EC2 vía SSH

   ssh -i ~/.ssh/my-key1.pem ec2-user@<IP_PUBLICA>
   

Manejo de instancias para optimización de costos
Las instancias se detienen para evitar costos innecesarios. Para reactivarlas cuando sea necesario
comandos:

1. Iniciar la instancia EC2
  
   aws ec2 start-instances --instance-ids <INSTANCE_ID>
  

2. Iniciar la base de datos en RDS

   aws rds start-db-instance --db-instance-identifier stockdb
  

3. Obtener la nueva IP pública de la instancia EC2
   aws ec2 describe-instances --instance-ids <INSTANCE_ID> --query 'Reservations[*].Instances[*].PublicIpAddress' --output text
   

4. Conectarse nuevamente a la instancia EC2
   ssh -i ~/.ssh/my-key1.pem ec2-user@<NEW_PUBLIC_IP>

Verificación del backend en la instancia EC2

Para comprobar que el backend está en ejecución:  

sudo systemctl status backend.service

Si el servicio no está corriendo, reiniciarlo manualmente:  
sudo systemctl restart backend.service

Pruebas de API

Una vez que el backend esté activo, se pueden realizar pruebas con los siguientes comandos:  

1. Obtener recomendaciones de inversión
   curl -X GET http://<NEW_PUBLIC_IP>:9090/recomendaciones

2. Listar todas las acciones almacenadas: 

   curl -X GET http://<NEW_PUBLIC_IP>:9090/acciones

Estos pasos aseguran que la infraestructura esté correctamente desplegada y funcional cuando se requiera demostrar el sistema.
