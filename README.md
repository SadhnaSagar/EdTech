# 🛡️ Golang RBAC System with Gin, GORM, and PostgreSQL

This is a working prototype of a **Role-Based Access Control (RBAC)** system built using:

- 🧠 Go (Golang)
- ⚙️ Gin web framework
- 🗃️ GORM ORM
- 🐘 PostgreSQL
- 🔐 JWT-based authentication

The RBAC system supports user-role-permission mapping, JWT login, and middleware-based access control.

---

## 🧩 Features

- User login with JWT token generation
- Role and permission based authorization
- Middleware to protect routes
- GORM auto migrations
- MVC project structure

---

## 📁 Folder Structure
```bash
├── database/ # DB config and connection
├── controllers/ # Request handlers
├── middleware/ # Authorization middleware
├── models/ # GORM models for User, Role, Permission
├── routes/ # Route initialization
├── utils/ # JWT utilities
├── .env # Environment variables
├── main.go # Application entry point
├── go.mod
```

---

## 🔧 Tech Stack

- **Go** 1.21+
- **Gin** for HTTP server
- **GORM** as ORM
- **PostgreSQL** as database
- **JWT** for authentication

---

## ⚙️ Getting Started

### 1. Clone the Repository

```bash
git clone git@github.com:SadhnaSagar/EdTech.git
cd EdTech
```

### 2. Install Dependecy
```bash
go mod tidy
```

### 3. Setup env
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=rbac_demo
JWT_SECRET=supersecretkey
```

### 4. Start PostgreSQL
Make sure PostgreSQL is running and the database rbac_demo exists.
Create the DB if not already:
```bash
CREATE DATABASE rbac_demo;
```

### 5.Run the App
```bash
go run main.go
```

## Authentication & Authorization
- Login generates a JWT token
- Protected routes require the Authorization: Bearer <token> header
- Middleware checks user roles and permissions

## 🔑 Example Requests
Login and Get Token
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}'
```

Access Protected Endpoint
```bash
curl http://localhost:8080/api/dashboard \
  -H "Authorization: Bearer <your_jwt_token>"
```

## 🔧 Extending the Project
You can extend this base with:
Classroom and course models
User-course enrollments
Assignment and submission features
Service-to-service authentication using API tokens with RBAC

## 🧪 Testing the System
Create roles like admin, teacher, student
Assign permissions like read:course, write:assignment
Assign roles to users and test access via JWT-protected endpoints

## 📜 License
This project is provided for educational purposes only.

## 🙌 Author
Built with ❤️ using Go & PostgreSQL for EdTech Platform by Sadhna Sagar.
