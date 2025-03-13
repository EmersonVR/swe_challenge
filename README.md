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
