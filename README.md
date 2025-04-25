# Application Programming Interface Project

## Description
This project is a simple API implementation using the Go programming language (Golang). The API provides CRUD (Create, Read, Update, Delete) operations for managing student data in a MySQL database.

## Technologies Used
- **Golang**
- **MySQL**
- **Gorilla Mux** (for routing)

## Project Structure
```
├── controllers
│   ├── StudentsController.go
├── models
│   ├── StudentsModels.go
├── routes
│   ├── StudentsRoutes.go
├── main.go
```

## Installation and Configuration
1. **Clone the repository**
   ```bash
   git clone <repository_url>
   cd <repository_folder>
   ```
2. **Install dependencies**
   ```bash
   go mod tidy
   ```
3. **Create MySQL database**
   ```sql
   CREATE DATABASE learn_;
   ```
4. **Adjust database connection in `main.go`**
   ```go
   dsn := "root:@tcp(localhost:3306)/learn_"
   ```
5. **Run the application**
   ```bash
   go run main.go
   ```

## API Endpoints
| Method | Endpoint         | Description |
|--------|-----------------|-------------|
| GET    | `/students`     | Retrieve all students |
| GET    | `/students/{id}` | Retrieve a student by ID |
| POST   | `/students`     | Add a new student |
| PUT    | `/students/{id}` | Update student data by ID |
| DELETE | `/students/{id}` | Delete a student by ID |

## Team Members
- **Jayady Managam Gultom** (1252----)
- **Firman Sulaiman** (1252----)
- **Adrian Wahyuda Aditya P** (1252----)
- **Agung Dwi Pradipta** (1252----)
- **Ahmad Hardiansyah** (1252----)

## License
This project is created for academic purposes as part of the **Integrative Programming and Technology** course and can be used for further learning.

---
