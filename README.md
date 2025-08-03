# 🛒 Shopping Cart App

An end-to-end **Fullstack Shopping Cart Web Application** built with **React.js (Frontend)** and **Go (Gin, GORM, JWT) (Backend)** with PostgreSQL integration.

This project demonstrates modern web development practices including user authentication, protected APIs, state management, and responsive design — all built as part of a fullstack internship assignment.

---

## 🚀 Features

### 👤 User Panel
- 🔐 Secure login using JWT
- 🛍️ Browse and view product catalog
- 🛒 Add/remove items from cart
- 💳 Checkout and place orders
- 📜 View order history

### 🛠️ Backend System
- Built with **Go (Gin)**
- RESTful APIs for user, items, cart, and orders
- PostgreSQL database with GORM ORM
- Middleware for JWT verification

---

## 🧱 Tech Stack

| Layer     | Technology                     |
|-----------|--------------------------------|
| Frontend  | React.js, Axios, CSS Modules   |
| Backend   | Go, Gin, GORM, JWT             |
| Database  | PostgreSQL                     |
| Tools     | Postman, VS Code, Git & GitHub |

---

## 🗂️ Project Structure

shopping-cart/
├── backend/ # Gin backend server
│ ├── main.go
│ ├── models/
│ ├── routes/
│ └── database/
├── frontend/ # React UI
│ ├── src/
│ │ ├── components/
│ │ ├── App.js
│ │ └── index.js
├── README.md

yaml
Copy
Edit

---

## 🛠️ Getting Started

### ✅ Backend Setup

1. Navigate to backend directory:
   ```bash
   cd backend
Install Go modules:

bash
Copy
Edit
go mod tidy
Configure PostgreSQL in database/connection.go

Run server:

bash
Copy
Edit
go run main.go
Server starts at: http://localhost:8080

✅ Frontend Setup
Navigate to frontend directory:

bash
Copy
Edit
cd frontend
Install dependencies:

bash
Copy
Edit
npm install
Start app:

bash
Copy
Edit
npm start
App opens at: http://localhost:3000

📮 API Endpoints
Method	Endpoint	Description
POST	/users/login	User login
GET	/items	Fetch all items
POST	/carts	Add item to cart
GET	/carts	View current cart
POST	/orders	Checkout cart
GET	/orders	View order history

Protected routes require JWT in Authorization header.


📈 What I Learned
Structuring fullstack apps across frontend & backend

Handling JWT authentication securely

Interacting with relational databases using GORM

Building responsive UI components in React

Debugging and testing REST APIs using Postman

🤝 Acknowledgements
This project was built as part of a fullstack internship assignment to demonstrate understanding of real-world software architecture and development workflow.
