# ecommerce-api
## Description:
- This is a simple backend API design for a C2C (consumer to consumer) e-commerce website based on Go, using the Gin framework and GORM ORM.


##  Project Structure：
ecommerce-api/
├── main.go              # Program entry
├── config/              # Configuration management
│   └── config.go
├── controllers/         # Controllers
│   ├── user.go
│   ├── product.go
│   ├── order.go
│   └── review.go
├── models/              # Data model
│   ├── user.go
│   ├── product.go
│   ├── order.go
│   ├── review.go
│   └── db.go
├── routes/              # Route Management
│   └── routes.go
├── middlewares/         # Middleware
│   └── jwt.go
├── tests/               # Testing
│   └── api_test.go
├── go.mod               # Dependency management
└── .env                 # Configuration file


## Core API functions:
- User registration and login (with password encryption).
- Product management (add products, get product lists, search for products).
- Order management (create orders, view orders).
- Comment function (add comments to order products)


## Database table design:
- User table: users
- Product table: products
- Order table: orders
- Order item table: order_items
- Review table: reviews


## API Security:
- JWT Authentication


## Setup
1. Install dependencies:
   ```bash
   go mod init ecommerce-api
   go mod tidy


2. Setup environment variables:
   ```bash
   cp .env.example .env

3. Run the application:
   ```bash
   go run main.go
   

## Run test case:
   ```bash
   go test ./...




