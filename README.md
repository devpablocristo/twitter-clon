# Twitter Clon

Una versión simplificada de una plataforma de microblogging similar a Twitter, que permite a los usuarios publicar tweets, seguir a otros usuarios y ver su timeline de tweets.  
Este proyecto fue desarrollado como parte del proceso de selección para Ualá.

---

## Instalación y Ejecución

### Entorno de Desarrollo

1. Clona el repositorio y accede al directorio del proyecto:
   ```bash
   git clone https://github.com/tu_usuario/twitter-clon.git
   cd twitter-clon/projects/qh
   ```

2. Levanta las dependencias (Cassandra, PostgreSQL, Redis, RabbitMQ, etc.) utilizando Docker Compose:
   ```bash
   docker-compose up -d
   ```

3. Construye la aplicación para el entorno de desarrollo:
   ```bash
   make qh-dev-build
   ```

4. Para depurar en Visual Studio Code:
   - Abre el proyecto en VS Code.
   - Levantar el container (visual code - attach visual studio code-) 
   - Correr el debugger (F5).

### Entorno de Prueba / Staging

1. Asegúrate de tener levantados los servicios de Docker Compose (como en el entorno de desarrollo).

2. Construye la aplicación para el entorno de prueba:
   ```bash
   make qh-stg-build
   ```

En ambos entornos, la aplicación se iniciará en el puerto configurado (por defecto, 8080).

Se incluye el archvo de variables de entorno .env por simplicidad.

---

## Estructura del Proyecto

La estructura del proyecto es la siguiente (se muestran las carpetas principales):

```
pkg                         # liberias de infrastrutura

Proyecto
├── cmd
│   └── api
│       └── main.go         # Punto de entrada de la aplicación
├── internal
│   ├── authe               # Lógica de Autenticación
│   ├── config              # Configuración general
│   ├── person              # Lógica de Person (PostgreSQL)
│   ├── tweet               # Lógica de Tweets (Cassandra, Redis, RabbitMQ)
│   └── user                # Lógica de Users (GORM)
├── integration-tests       # Tests de integración para tweets
├── mocks                   # Mocks generados con GoMock
└── wire                    # Inyección de dependencias con Wire
```

---

## Consideraciones y Asunciones de Negocio

- **Autenticación:** No se implementa un sistema de *sign in* o manejo de sesiones, ya que se asume que todos los usuarios son válidos. El identificador del usuario se envía como parámetro.
- **Escalabilidad:** La solución está pensada para escalar a millones de usuarios mediante el uso de bases de datos distribuidas (Cassandra), cache distribuido (Redis) y mensajería asíncrona (RabbitMQ).
- **Optimización para Lecturas:** Se prioriza la optimización en operaciones de lectura, utilizando técnicas de cache y una estructura de datos denormalizada para el timeline.
- **Testing:** Se incluyen tests unitarios y de integración para cubrir los casos de uso principales.
- **Infraestructura:** La aplicación se despliega en contenedores utilizando Docker Compose, lo que facilita la replicación del entorno de producción en desarrollo y pruebas.

---

## Notas Adicionales

- **Docker Compose:** El archivo `docker-compose.yml` incluye la configuración para levantar servicios como Cassandra, PostgreSQL, Redis y RabbitMQ.
- **Wire:** Se utiliza para la inyección de dependencias, lo que facilita la escalabilidad y el mantenimiento del código.
- **Make:** Se utilizan targets en el Makefile para simplificar la construcción de la aplicación en diferentes entornos.