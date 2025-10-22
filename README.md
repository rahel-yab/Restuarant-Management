# 🍽️ Restaurant Management System

A RESTful API for managing a restaurant, built with Go, Gin, and MongoDB.

## ✨ Features

- User authentication (signup, login)
- Food, Menu, Order, Invoice, Table, and Order Item management
- Pagination and basic filtering for list endpoints
- JWT-based authentication middleware (access + refresh tokens)
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
- MongoDB (local or remote)
- PowerShell / Terminal for running commands

### 🔑 Required environment variables

Create a `config.env` (or `.env`) in the project root with at least the following:

- `MONGO_URI` - MongoDB connection string (example: `mongodb://localhost:27017`)
- `DB_NAME` - database name used by the project (example: `RestaurantDB`)
- `ACCESS_TOKEN_SECRET` - secret for signing access tokens
- `REFRESH_TOKEN_SECRET` - secret for signing refresh tokens
- `PORT` - optional; default `8000`

Note: Some code defaults may differ (DB name or default port). Verify `database/dbConnection.go` and `main.go` if you see connection errors.

### 📦 Installation & Quick start (PowerShell)

1. Install dependencies and tidy modules:

   ```sh
   go mod tidy
   ```

2. Run the server:

   ```sh
   $env:ACCESS_TOKEN_SECRET = "your-access-secret"; $env:REFRESH_TOKEN_SECRET = "your-refresh-secret"; go run main.go
   ```

Or set variables in `config.env` and load them before running.

The server listens on port `8000` by default. Visit `http://localhost:8000` (or your configured port).


---

### API Endpoints

- **Authentication:**

  - `POST /users/signup` - Create a new user
  - `POST /users/login` - Authenticate user and receive tokens

- **Food Management:**

  - `GET /foods` - List foods (pagination, filtering)
  - `POST /foods` - Add a new food item
  - `PUT /foods/{id}` - Update a food item
  - `DELETE /foods/{id}` - Remove a food item

- **Menu Management:**

  - `GET /menus` - List menus (pagination, filtering)
  - `POST /menus` - Create a new menu
  - `PUT /menus/{id}` - Update a menu
  - `DELETE /menus/{id}` - Delete a menu

- **Order Management:**

  - `GET /orders` - List orders (pagination, filtering)
  - `POST /orders` - Create a new order
  - `PUT /orders/{id}` - Update an order
  - `DELETE /orders/{id}` - Cancel an order

- **Order Item Management:**

  - `GET /order_items` - List order items (pagination, filtering)
  - `POST /order_items` - Add an item to an order
  - `PUT /order_items/{id}` - Update an order item
  - `DELETE /order_items/{id}` - Remove an item from an order

- **Invoice Management:**

  - `GET /invoices` - List invoices (pagination, filtering)
  - `POST /invoices` - Create a new invoice
  - `PUT /invoices/{id}` - Update an invoice
  - `DELETE /invoices/{id}` - Delete an invoice

- **Table Management:**
  - `GET /tables` - List tables (pagination, filtering)
  - `POST /tables` - Add a new table
  - `PUT /tables/{id}` - Update a table
  - `DELETE /tables/{id}` - Remove a table


## 🤝 Contributing

Contributions are welcome! If you find a bug or want to add a feature, open an issue or submit a pull request.