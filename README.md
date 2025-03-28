# 🌍 CountryInfo API

CountryInfo API is a **Golang-based microservice** that provides country-related data such as **economy, population, history, and military information**. Built with **Gin, GORM, and PostgreSQL**, this service exposes RESTful APIs to fetch country details.

---

## 🚀 Features
- ✅ Fetch a list of all countries (`GET /countries`)
- ✅ Get detailed economic data for a country (`GET /countries/{id}/economy`)
- ✅ Designed for **scalability** with a microservices approach
- ✅ Uses **PostgreSQL** as the database with schema `country_infos`

---

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
 - Install GORM (Go ORM):
  ```sh
  go get -u gorm.io/gorm gorm.io/driver/postgres
  ```
  
###  2️⃣ Database Setup
 - Create DB, Schema and Tables script is available under db/initial directory

###  3️⃣ Running the API
 - With Go :
 ```sh
   go run cmd/main.go
 ```
 - With docker :
```sh
   docker-compose up --build
```
## 📜 License
This project is licensed under the MIT License.
