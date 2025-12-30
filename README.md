# 🚢 Enterprise Logistics & Analytics Suite

![Architecture](https://img.shields.io/badge/Architecture-Microservices-orange)
![Traefik](https://img.shields.io/badge/Proxy-Traefik-blue?logo=traefik)
![Laravel](https://img.shields.io/badge/Core-Laravel_10-ff2d20?logo=laravel)
![Go](https://img.shields.io/badge/Analytics-Golang-00add8?logo=go)
![Vuetify](https://img.shields.io/badge/UI-Vuetify_2-1867c0?logo=vuetify)

A professional-grade logistics management system built with a polyglot microservices approach. This suite handles inventory, order processing, and real-time data analytics through a unified entry point.



---

### 🏛️ System Architecture

This project showcases a decoupled infrastructure where each service is chosen for its specific strengths:

* **🛡️ Core API (Laravel 10):** Handles complex business rules, user authentication, and relational data management for inventory and orders.
* **⚡ Analytics Engine (Go):** A high-performance microservice dedicated to real-time data processing and logistics forecasting.
* **🎨 Admin Dashboard (Vue 2 + Vuetify):** A Material Design frontend providing a reactive and intuitive user experience for warehouse managers.
* **🚦 Edge Router (Traefik):** Acts as the unified API Gateway and Reverse Proxy, managing service discovery and routing.

---

### 🛠️ Technology Stack

| Service | Technology | Role |
| :--- | :--- | :--- |
| **Gateway** | Traefik | Reverse Proxy & Load Balancing |
| **Frontend** | Vue.js 2 + Vuetify | Enterprise UI/UX |
| **Backend A** | Laravel (PHP 8.2) | Business Logic & PostgreSQL ORM |
| **Backend B** | Golang 1.21 | High-speed Analytics API |
| **Database** | PostgreSQL | Primary Data Store |

---

### 🚀 Getting Started

#### Prerequisites
- Docker & Docker Compose
- Git

#### Installation

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/logistics-suite.git](https://github.com/your-username/logistics-suite.git)
    cd logistics-suite
    ```

2.  **Initialize Services:**
    The project uses a unified `docker-compose.yml` to orchestrate all containers.
    ```bash
    docker-compose up --build
    ```

3.  **Setup Database (Laravel):**
    ```bash
    docker-compose exec laravel php artisan migrate --seed
    ```

#### Routing Table
Once the stack is running, Traefik routes the traffic as follows:
- 🏠 **Frontend:** `http://localhost/`
- ⚙️ **Core API:** `http://localhost/api/core`
- 📊 **Analytics:** `http://localhost/api/stats`
- 📈 **Traefik Dashboard:** `http://localhost:8081`

---

### ⚓ DevOps & Architectural Decisions

- **Service Independence:** Each microservice has its own `Dockerfile`, allowing for independent scaling and maintenance.
- **Polyglot Strategy:** Leveraging Go for CPU-intensive tasks while using Laravel's robust ecosystem for rapid business logic development.
- **Material Design:** Implementing Vuetify to ensure a consistent, professional, and accessible interface for enterprise users.
- **Unified Gateway:** Using Traefik labels for dynamic routing, removing the need to expose multiple ports to the end user.

---
