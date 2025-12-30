# 🚢 Enterprise Logistics Multi-Service Suite

![Architecture](https://img.shields.io/badge/Architecture-Polyglot_Microservices-orange)
![Traefik](https://img.shields.io/badge/Gateway-Traefik_2.10-blue?logo=traefik)
![Laravel](https://img.shields.io/badge/Core-Laravel_11-ff2d20?logo=laravel)
![Go](https://img.shields.io/badge/Analytics-Golang_1.21-00add8?logo=go)
![Vuetify](https://img.shields.io/badge/UI-Vuetify_2-1867c0?logo=vuetify)

A professional-grade logistics ecosystem designed to showcase **Service Decoupling**, the **API Gateway Pattern**, and **High-Performance Microservices**. This suite integrates three different technology stacks (PHP, Go, and Vue.js) into a unified infrastructure orchestrated via Docker.

---

## 🏛️ System Architecture

The project is designed with a "separation of concerns" approach, where each language is chosen for its specific strengths:

* **🎨 Frontend (Vue.js 2 + Vuetify):** An enterprise-grade dashboard using Material Design to monitor real-time service health and logistics data.
* **🛡️ Core API (Laravel 11):** Handles complex business logic, user authentication, and relational data management via PostgreSQL.
* **⚡ Analytics Engine (Go):** A lightweight, high-concurrency microservice optimized for rapid data processing and statistics.
* **🚦 Edge Router (Traefik):** Acts as the single entry point (Port 80) for the entire ecosystem, handling dynamic routing and service discovery.



---

## 🚀 Getting Started

### Prerequisites
- Docker and Docker Compose installed on your machine.

### Installation & Launch
1.  **Clone the repository and enter the directory:**
    ```bash
    git clone [https://github.com/your-username/logistics-suite.git](https://github.com/your-username/logistics-suite.git)
    cd logistics-suite
    ```

2.  **Build and start the fleet:**
    ```bash
    docker-compose up -d --build
    ```

3.  **Initialize the Core Database:**
    ```bash
    docker-compose exec laravel php artisan migrate
    ```

---

## ✅ Live Verification (Browser Testing)

Once the stack is running, you can verify the orchestration is working by visiting these URLs:

| Component | URL Endpoint | Expected Result |
| :--- | :--- | :--- |
| **Dashboard UI** | [http://localhost/](http://localhost/) | The main Web App showing service status cards. |
| **Laravel Core API** | [http://localhost/api/core/status](http://localhost/api/core/status) | A JSON response directly from the PHP backend. |
| **Go Analytics** | [http://localhost/api/stats](http://localhost/api/stats) | A JSON response directly from the Go engine. |
| **Traefik Admin** | [http://localhost:8082](http://localhost:8082) | Technical dashboard showing "Healthy" routers. |

---

## 🛠️ Key Technical Features

-   **API Gateway Pattern:** Using Traefik to abstract internal ports and provide a unified API interface.
-   **Priority Routing:** Intelligent path management to prevent conflicts between the Frontend and API routes.
-   **Container Isolation:** All services communicate within a secure internal Docker network, shielded from direct external access.
-   **Polyglot Strategy:** Demonstrating the ability to choose the right tool for the job (Go for speed, Laravel for stability).

---
