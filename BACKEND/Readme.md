### 🖥️ Backend README
Place this file inside your backend folder (e.g., inside the `BACKEND/` folder).

**Filename:** `BACKEND/README.md`

```markdown
# Backend Service (Golang)

This directory contains the backend REST API service written in **Go**. It handles user authentication (Login/Register) and manages the database with **Auto-Migration**.

## ⚙️ Configuration (.env)

Before running the application, you must create a `.env` file in this directory with the following configuration:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=amass-test
PORT=8080
```

⚠️ Important: Please ensure that a MySQL database named amass-test (or the name defined in DB_NAME) is created in your database server before starting the application.

## How to Run
Open your terminal at the root of the project and run the following command:
```
go run api/main.go
```
