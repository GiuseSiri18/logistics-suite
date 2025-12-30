# 🚢 Enterprise Logistics Multi-Service Suite

![Architecture](https://img.shields.io/badge/Architecture-Microservices-orange)
![Traefik](https://img.shields.io/badge/Gateway-Traefik_2.10-blue)
![Laravel](https://img.shields.io/badge/Core-Laravel_11-ff2d20)
![Go](https://img.shields.io/badge/Analytics-Golang_1.21-00add8)

Una piattaforma logistica enterprise che dimostra l'integrazione di diverse tecnologie (PHP, Go, JavaScript) in un ecosistema orchestrato con Docker e instradato tramite un API Gateway unico.

## 🏛️ Architettura del Sistema
Il progetto è suddiviso in tre servizi principali accessibili tramite la porta standard HTTP (80):

* **Frontend (Vue.js + Vuetify):** Una dashboard Material Design che monitora lo stato dei servizi.
* **Core API (Laravel):** Gestisce la logica di business e l'integrazione con il database PostgreSQL.
* **Analytics Engine (Go):** Un microservizio ultra-veloce per il calcolo delle statistiche.

---

## 🚀 Come avviare il progetto

1.  **Avvia i container:**
    ```bash
    docker-compose up -d --build
    ```
2.  **Inizializza il database Core:**
    ```bash
    docker-compose exec laravel php artisan migrate
    ```

---

## ✅ Verifica del Funzionamento (Browser)

Puoi verificare che l'orchestrazione funzioni correttamente visitando questi URL:

| Componente | Indirizzo Browser | Cosa dovresti vedere |
| :--- | :--- | :--- |
| **Dashboard UI** | [http://localhost/](http://localhost/) | La web app con le card colorate che mostrano "online". |
| **Laravel API** | [http://localhost/api/core/status](http://localhost/api/core/status) | Un JSON di conferma dal core PHP. |
| **Go API** | [http://localhost/api/stats](http://localhost/api/stats) | Un JSON di conferma dal motore Go. |
| **Traefik Health** | [http://localhost:8082](http://localhost:8082) | La dashboard tecnica con i router tutti "verdi". |

---

## 🛠️ Highlights Tecnici
- **API Gateway Pattern:** Tutte le richieste passano per Traefik, eliminando la necessità di esporre porte multiple.
- **Service Priority:** Gestione intelligente dei percorsi per evitare conflitti tra frontend e API.
- **Container Isolation:** I servizi comunicano in una rete interna sicura, protetti dal gateway.