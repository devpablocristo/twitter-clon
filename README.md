# Twitter Clon

Una versión simplificada de una plataforma de microblogging similar a Twitter, que permite a los usuarios publicar tweets, seguir a otros usuarios y ver su timeline de tweets.  
Este proyecto fue desarrollado como parte del proceso de selección para Ualá.

---

## 📌 Instalación y Ejecución

### 🔹 Entorno de Desarrollo

1. **Clona el repositorio y accede al directorio del proyecto:**
   ```bash
   git clone https://github.com/tu_usuario/twitter-clon.git
   cd twitter-clon/projects/qh
   ```

2. **Construye la aplicación para el entorno de desarrollo (recomendado para debug):**  
   ```bash
   make qh-dev-build
   ```

3. **Depuración en Visual Studio Code:**
   - Abre el proyecto en VS Code.
   - Adjunta el contenedor (Attach Visual Studio Code).
   - Ejecuta el debugger (`F5`).

### 🔹 Entorno de Prueba / Staging

1. **Asegúrate de que los servicios de Docker Compose estén en ejecución.**

2. **Construye la aplicación para el entorno de prueba (recomendado para pruebas):**
   ```bash
   make qh-stg-build
   ```

📍 **Por defecto, la aplicación se ejecuta en el puerto 8080.**  
📍 **Se incluye el archivo `.env` para simplificar la configuración de variables de entorno.**

---

### 🔹 Docker Compose

Si prefieres usar Docker Compose para levantar los servicios, ejecuta:

   ```bash
   docker-compose up -d
   ```

🔹 **Nota:** Aunque se puede utilizar Docker Compose, se recomienda usar `make` para una mejor gestión de la aplicación.

---

## 📂 Estructura del Proyecto

```
pkg                         # Librerías de infraestructura

Proyecto
├── cmd
│   └── api
│       └── main.go         # Punto de entrada de la aplicación
├── internal
│   ├── auth               # Lógica de autenticación
│   ├── config             # Configuración general
│   ├── person             # Lógica de personas (PostgreSQL)
│   ├── tweet              # Lógica de tweets (Cassandra, Redis, RabbitMQ)
│   └── user               # Lógica de usuarios (GORM)
├── integration-tests       # Tests de integración para tweets
├── mocks                   # Mocks generados con GoMock
└── wire                    # Inyección de dependencias con Wire
```

---

## ⚡ Consideraciones y Asunciones de Negocio

✅ **Autenticación:**  
No se implementa un sistema de *sign in* o manejo de sesiones. Se asume que todos los usuarios son válidos y se identifican mediante un parámetro en las solicitudes.

✅ **Escalabilidad:**  
La solución está diseñada para soportar millones de usuarios utilizando bases de datos distribuidas (*Cassandra*), caché distribuido (*Redis*) y mensajería asíncrona (*RabbitMQ*).

✅ **Optimización para Lecturas:**  
Se prioriza la eficiencia en consultas utilizando estrategias de cache y estructuras de datos desnormalizadas para el timeline.

✅ **Testing:**  
Se incluyen pruebas unitarias y de integración para validar los casos de uso principales.

✅ **Infraestructura:**  
Se despliega en contenedores mediante *Docker Compose* para facilitar la replicación en distintos entornos.

---

## 🚀 Pruebas con `curl`

### 🏷️ Crear una Persona
```bash
curl --location 'localhost:8080/api/v1/person/public' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "Homero",
    "last_name": "Simpson",
    "age": 40,
    "gender": "male",
    "national_id": 1234567890,
    "phone": "+11234567890",
    "interests": ["beer", "doughnuts", "TV"],
    "hobbies": ["fishing", "sleeping"]
}'
```

### 🏷️ Crear un Usuario
```bash
curl --location 'localhost:8080/api/v1/users/public' \
--header 'Content-Type: application/json' \
--data '{
  "user_type": "regular",
  "email_validated": false,
  "person_id": "3d1130b1-2dff-4e8c-bbbc-e56e5c960460",
  "credentials": {
    "email": "homero.simpson@example.com",
    "password": "doh123"
  },
  "roles": [
    {
      "name": "user",
      "permissions": [
        {
          "name": "tweet:read",
          "description": "Permite leer tweets"
        },
        {
          "name": "tweet:write",
          "description": "Permite escribir tweets"
        }
      ]
    }
  ]
}'
```

### 🏷️ Seguir a un Usuario
```bash
curl --location 'localhost:8080/api/v1/users/public/follow' \
--header 'Content-Type: application/json' \
--data '{
    "follower_id": "fb312abc-d3ee-4450-8c3c-ac4ff1f3ed86",
    "followee_id": "8aa26deb-ab89-42df-a757-59633b35731b"
}'
```

### 🏷️ Publicar un Tweet
```bash
curl --location 'localhost:8080/api/v1/tweets/public' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": "8aa26deb-ab89-42df-a757-59633b35731b",
    "content": "¡Me gustan las rosquillas! 🍩"
}'
```

### 🏷️ Consultar el Timeline de un Usuario
```bash
curl --location 'localhost:8080/api/v1/tweets/public/fb312abc-d3ee-4450-8c3c-ac4ff1f3ed86/timeline'
```
