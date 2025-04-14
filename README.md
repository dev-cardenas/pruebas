# Documentación del Proyecto

## 🚀 Cómo ejecutar el proyecto

Para correr el proyecto Go, sigue estos pasos:

1. **Instalar dependencias:**

   Asegúrate de tener **Go** instalado en tu máquina. Si no lo tienes, puedes descargarlo desde [aquí](https://golang.org/doc/install).

2. **Cargar dependencias:**

   En la terminal, dentro del directorio raíz del proyecto, ejecuta:

   ```bash
   go mod tidy
   ```

## 📦 Estructura del Proyecto

La estructura de carpetas y archivos sigue una arquitectura modular y limpia, basada en los principios de Arquitectura Hexagonal y SOLID. A continuación, se desglosan los directorios y su propósito.

### 🏁 cmd/server/main.go (Punto de entrada)

Este es el punto de entrada de la aplicación, donde se orquestan todas las dependencias y servicios.

**_¿Por qué está aquí?_**

Siguiendo el patrón de arquitectura hexagonal, este archivo se mantiene limpio. Su función principal es iniciar las dependencias necesarias (como la base de datos, configuraciones, y routers) y arrancar el servidor.

Si en el futuro se necesitan otras entradas para la aplicación, como un worker o CLI, podemos agregar otros archivos bajo el directorio cmd/.

### 🧠 configs/config.go (Configuración centralizada)

Este archivo gestiona todas las configuraciones del sistema, como las variables de entorno y los valores por defecto.

**_¿Por qué está aquí?_**

Separa las preocupaciones de configuración de la lógica de negocio.

Centraliza las configuraciones, evitando tener valores dispersos en múltiples lugares, lo que facilita la modificación y el mantenimiento.

Hace más fácil la integración de pruebas (tests) al aislar las configuraciones de la lógica de negocio.

### 🧩 internal/user/domain (Dominio)

El dominio es el núcleo de la lógica de negocio. Aquí definimos todo lo relacionado con los "usuarios", sin ninguna dependencia de frameworks o tecnologías externas.

#### Archivos

- entity.go: Define el modelo de "User". Aquí se encuentran los datos fundamentales del dominio.

- repository.go: Define las interfaces que permiten la persistencia de los datos de usuario (por ejemplo, la búsqueda, creación o eliminación).

- service.go: Aquí se define la lógica del negocio relacionada con los usuarios, como validaciones o reglas complejas.

**_¿Por qué está aquí?_**

Este código es completamente independiente de cualquier framework o tecnología. Solo importa la lógica de negocio.

Sigue el principio de SOLID, particularmente el principio de responsabilidad única, asegurando que el dominio sea completamente testable y fácil de mantener.

🔌 internal/infrastructure/ (Infraestructura)
En esta capa se gestionan las dependencias externas como bases de datos, HTTP o servicios de terceros. Aquí se incluyen los adaptadores que conectan el dominio con estas tecnologías.

Archivos:

db/postgres/db.go: Conecta con la base de datos PostgreSQL usando pgxpool.

db/postgres/user_repository.go: Implementa las interfaces definidas en domain/user/repository.go para interactuar con la base de datos.

http/gin/router.go: Define las rutas y configuraciones del servidor HTTP (usando Gin).

http/gin/handler.go: Contiene las funciones que traducen entre el dominio y las peticiones HTTP, como crear, actualizar o eliminar usuarios.

¿Por qué está aquí?

Separamos la infraestructura del dominio para que el código del dominio no dependa de tecnologías específicas. Esto hace que la arquitectura sea flexible y fácil de modificar.

Si en algún momento decidimos cambiar de framework (por ejemplo, de Gin a Fiber), solo necesitaríamos modificar esta capa.

🎯 internal/application/user/ (Aplicación)
Esta capa orquesta los casos de uso de la aplicación. Aquí se definen las acciones que pueden realizar los usuarios, como crear, actualizar o eliminar un usuario.

¿Por qué está aquí?

Separa la lógica de negocio de la lógica de infraestructura. Esto permite que los casos de uso puedan evolucionar de manera independiente de las tecnologías externas.

Ideal para agregar validaciones, reglas de negocio complejas o coordinar diferentes servicios.

🧪 test/user/service_test.go (Pruebas unitarias)
Los tests unitarios se centran en la lógica de negocio pura, sin involucrar las dependencias externas (como bases de datos o servicios HTTP). Esto permite pruebas rápidas y confiables.

¿Por qué está aquí?

Se asegura que las reglas del negocio se comporten correctamente sin tener que depender de la infraestructura externa.

Sigue el principio de Testabilidad y Responsabilidad Única, permitiendo una rápida evolución del código de negocio con pruebas automatizadas.

🧩 Patrones aplicados

1. Arquitectura Hexagonal (Ports and Adapters)
   Separa el núcleo de la aplicación (dominio) de las implementaciones externas (como bases de datos, servicios externos y frameworks).

Facilita la testabilidad y permite cambiar fácilmente las dependencias sin afectar el dominio.

2. SOLID
   S (Single Responsibility Principle): Cada módulo, clase o función tiene una única responsabilidad.

O (Open/Closed Principle): El sistema está abierto para extensión, pero cerrado para modificación. Se pueden agregar nuevos casos de uso sin cambiar el código existente.

L (Liskov Substitution Principle): Las subclases deben ser intercambiables por sus clases base, lo que asegura una arquitectura flexible.

I (Interface Segregation Principle): Se utiliza interfaces pequeñas y específicas para cada caso de uso, evitando la sobrecarga de interfaces grandes.

D (Dependency Inversion Principle): Las dependencias de los módulos de alto nivel son abstraídas y no dependen de módulos de bajo nivel.

3. Inyección de Dependencias
   Las dependencias (como las conexiones de base de datos o los casos de uso) se pasan de manera explícita a través de constructores. Esto facilita la prueba de unidades de código y permite sustituir implementaciones sin modificar el código del negocio.
