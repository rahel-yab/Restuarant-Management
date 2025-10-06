# 🍽️ Restaurant Management System

A RESTful API for managing a restaurant, built with Go, Gin, and MongoDB.

## ✨ Features

- User authentication (signup, login)
- Food, Menu, Order, Invoice, Table, and Order Item management
- Pagination and filtering for food and menu items
- JWT-based authentication middleware
- MongoDB integration

## 📁 Project Structure

- `main.go` - Entry point, sets up routes and middleware
- `controllers/` - Request handlers for each resource (food, menu, order, invoice, user, etc.)
- `models/` - Data models for MongoDB collections
- `routes/` - Route definitions for each resource
- `middleware/` - Authentication and other middleware
- `database/` - MongoDB connection logic
- `helpers/` - Utility functions (e.g., JWT token helpers)

## 🚀 Getting Started

### 🛠️ Prerequisites

- Go 1.24+
- MongoDB (default URI: `mongodb://localhost:20717`, DB name: `Restuarant_managment`)
- (Optional) Postman for API testing

### 📦 Installation

1. Clone the repository
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Set up your `.env` file (see `config.env` for example)
4. Start MongoDB on your machine

### ▶️ Running the Server

```sh
go run main.go
```

The server will start on port `8000` by default (or the value of the `PORT` environment variable).

---

🤝 Contributing
contributions are welcome! If you discover a bug or have a feature request, feel free to open an issue or submit a pull request.
