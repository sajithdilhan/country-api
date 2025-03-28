# 🌍 CountryInfo API

CountryInfo API is a **Golang-based microservice** that provides country-related data such as **economy, population, history, and military information**. Built with **Gin, GORM, and PostgreSQL**, this service exposes RESTful APIs to fetch country details.

---

## 🚀 Features
- ✅ Fetch a list of all countries (`GET /countries`)
- ✅ Get detailed economic data for a country (`GET /countries/{id}/economy`)
- ✅ Designed for **scalability** with a microservices approach
- ✅ Uses **PostgreSQL** as the database with schema `country_infos`

---

## 🏗 Project Structure
country-api/
│── cmd/                     # Entry points for the application
│   ├── main.go              # Main application entry point
│
│── internal/                # Business logic (not exposed externally)
│   ├── country/             # Country-related business logic
│   │   ├── handler.go       # HTTP handlers for country APIs
│   │   ├── service.go       # Business logic for countries
│   │   ├── repository.go    # Database access layer for countries
│   │   ├── model.go         # Country entity struct
│   │   ├── dto.go           # Request/Response DTOs
│   │
│   ├── economy/             # Economy-related logic (similar structure)
│   │   ├── handler.go       
│   │   ├── service.go       
│   │   ├── repository.go    
│   │   ├── model.go         
│   │   ├── dto.go           
│
│── api/                     # API specifications (if using OpenAPI)
│   ├── openapi.yaml         
│
│── config/                  # Configuration files
│   ├── config.yaml          
│
│── db/                      # Database migrations, SQL scripts
│   ├── migrations/          
│
│── pkg/                     # Shared utilities/helpers
│   ├── database/            # Database connection handling
│   ├── middleware/          # Common middlewares
│   ├── logger/              # Logging setup
│   ├── httpserver/          # HTTP server setup
│
│── test/                    # Unit and integration tests
│── go.mod                   
│── go.sum                   
│── Dockerfile               
│── Makefile                 
│── README.md                

## ⚙️ Tech Stack
- **Language**: Go (Golang)
- **Framework**: [Gin](https://github.com/gin-gonic/gin) (HTTP framework)
- **Database**: PostgreSQL (Schema: `country_infos`)
- **ORM**: [GORM](https://gorm.io/)
- **Containerization**: Docker (Optional)

---

## 🔧 Setup & Installation

### **1️⃣ Prerequisites**
- Install **Go 1.21+**: [Download](https://go.dev/dl/)
- Install **PostgreSQL**: [Download](https://www.postgresql.org/download/)
- Install **Gin** (Go Web Framework):
  ```sh
  go get -u github.com/gin-gonic/gin
  Install GORM (Go ORM):
  go get -u gorm.io/gorm gorm.io/driver/postgres
  
2️⃣ Database Setup
CREATE DATABASE CountryInfo;
CREATE SCHEMA country_infos;
go run migrations/migrate.go

📜 License
This project is licensed under the MIT License.

3️⃣ Running the API
With Go : go run cmd/main.go 

With docker : docker-compose up --build
