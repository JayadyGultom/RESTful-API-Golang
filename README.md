# Application Programming Interface Project

## Description
This repository contains a RESTful API project built with Golang to manage student data. It helps students understand and implement basic REST API principles by providing CRUD operations with MySQL as the database. This project aims to enhance understanding of designing stateless and resource-based endpoints following REST architecture.

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
