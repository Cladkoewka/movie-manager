# 🎬 Movie Manager API

A simple and extensible RESTful API for managing movies, posters, and user reviews. Built with Go, Gin, GORM, and Redis. Swagger is available for easy documentation and testing. Works with the [Movie Manager Frontend](https://github.com/Cladkoewka/movie-manager-frontend).

#### [🚀 Try Movie Manager In Work](https://cladkoewka.github.io/movie-manager-frontend/)


## Main Page
![main](https://github.com/user-attachments/assets/b7b73a34-76d4-4aae-af84-92286d5032a9)
## Movie Details Page
![Desc](https://github.com/user-attachments/assets/5db1da36-9295-44fa-bb37-899ff56c74bc)
## Search & Sorting
![search](https://github.com/user-attachments/assets/790a6090-f100-493c-bea3-1c164c93fc42)

## 🚀 Features

- Full CRUD for movies
- Upload and retrieve movie posters
- Filtering, sorting, pagination
- Basic reviews: add, get, delete
- Swagger UI documentation (`/swagger/index.html`)
- Redis-based caching for better performance
- JSON data loader for initial seeding

## 🛠️ Tech Stack

- **Backend:** Go, Gin
- **Database:** PostgreSQL + GORM
- **Cache:** Redis
- **Docs:** Swagger (`swaggo/gin-swagger`)
- **Storage:** File uploads for posters (videos with B2 optional)
- **Deployment:** Easily deployable via Docker (optional)

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/Cladkoewka/movie-manager.git
cd movie-manager
```

Install dependencies:

```bash
go mod tidy
```

Generate Swagger docs:

```bash
swag init
```

## 🔧 Configuration

Create a `.env` file or export environment variables manually:

```bash
DB_USER=
DB_NAME=
DB_PASSWORD=
DB_HOST=
DB_PORT=
```

## 🗄️ Migrate & Seed Database

Run database migrations:

```bash
go run cmd/movie-manager.go -migrate
```

Load initial data from JSON dumps:

```bash
go run cmd/movie-manager.go -load
```

## ▶️ Run the API

```bash
go run cmd/movie-manager.go
```

By default, server runs on `http://localhost:8080`.

## 📚 API Documentation

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

## 📂 Available Endpoints

![swag](https://github.com/user-attachments/assets/2b2a5da6-1666-41cc-a3d0-cda1698f4b87)

### 🎥 Movies

- `GET /movies`: Get all movies (supports filters, sorting, pagination)
- `GET /movies/:id`: Get movie by ID
- `POST /movies`: Create a movie
- `PUT /movies/:id`: Update a movie
- `DELETE /movies/:id`: Delete a movie
- `POST /movies/:id/poster`: Upload movie poster
- `GET /movies/:id/poster`: Get movie poster

### 📝 Reviews

- `GET /reviews/movie/:movie_id`: Get reviews for a movie
- `POST /reviews`: Create a review
- `DELETE /reviews/:id`: Delete a review

## 🛠️ Future Improvements

- Add authentication and user system (JWT)
- Rating system
- Video trailer uploads (Backblaze B2)
- Caching strategies per endpoint
- Docker & CI/CD

## 📄 License

MIT — free to use, share, and modify.
