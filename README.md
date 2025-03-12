# 🛒 **Idempotencia API - Gestión de Pedidos y Productos**  

📌 **Proyecto de API en Go con Idempotencia para Órdenes de Compra**  

---

## **🛠 Tecnologías utilizadas**
- **Golang (1.24-alpine)** 🐹  
- **Gorilla Mux** (Enrutamiento) 🌐  
- **GORM** (ORM para MySQL) 🗄️  
- **MySQL (latest)** (Base de datos) 💾  
- **Redis (latest)** (Manejo de Idempotencia) 🔁  
- **Docker & Docker Compose** 🐳  

---

## **📦 Instalación y Ejecución**
### **🔹 Clonar el repositorio**
```bash
git clone https://github.com/JoseAyala97/Idempotencia.git
cd Idempotencia
```

### **🔹 Configurar variables de entorno**
Crea un archivo **`.env`** en la raíz del proyecto y agrega:
```ini
# Redis
REDIS_HOST=redis:6379
REDIS_PASSWORD=

# MySQL
MYSQL_HOST=mysql
MYSQL_PORT=3306
MYSQL_ROOT_PASSWORD=secret
MYSQL_DB=order_management
```

### **🔹 Ejecutar con Docker Compose**
```bash
docker-compose up --build
```
> 🚀 **Esto iniciará la API, Redis y MySQL con seeders que crean datos de prueba.**  

Una vez que veas el mensaje **"Servidor corriendo en :8080"**, la API estará lista en:
```
http://localhost:8080
```

---

## **📌 Endpoints Disponibles**
### **🔹 Productos**
| Método | URL | Descripción |
|--------|-----|------------|
| `GET` | `/products` | Obtiene la lista de productos disponibles |
| `PUT` | `/products/{id}/stock` | Actualiza el stock de un producto |

#### **Ejemplo de petición `PUT` para actualizar stock**
```http
PUT http://localhost:8080/products/1/stock
Content-Type: application/json

{
    "stock": 20
}
```

---

### **🔹 Órdenes de compra (Idempotente)**
| Método | URL | Descripción |
|--------|-----|------------|
| `POST` | `/orders` | Crea un pedido (con idempotencia) |
| `GET` | `/orders/{id}` | Obtiene detalles de un pedido |

#### **📌 Para crear un pedido, es obligatorio enviar el `Idempotency-Key` en el header**
Ejemplo de petición `POST` para crear una orden:
```http
POST http://localhost:8080/orders
Content-Type: application/json
Idempotency-Key: 123e4567-e89b-12d3-a456-426614174000

{
    "customer_name": "Juan Pérez",
    "order_items": [
        {
            "product_id": 1,
            "quantity": 2
        }
    ]
}
```
🔹 **Si se envía la misma `Idempotency-Key` otra vez, la respuesta será la misma sin duplicar la orden.**

Ejemplo de respuesta:
```json
{
    "order_id": 1
}
```

---

## **🔍 Notas adicionales**
- **La API maneja errores correctamente con respuestas JSON estructuradas.**
- **La persistencia de datos en MySQL está asegurada gracias a los `seeders`.**
- **Redis almacena temporalmente las claves de idempotencia para evitar duplicados.**

---

## **💡 Contribuir**
Si deseas mejorar esta API, siéntete libre de hacer un **fork** y enviar un **pull request** 🚀.