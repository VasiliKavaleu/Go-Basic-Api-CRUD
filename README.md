## Golang: Gin, Gorm; PostgreSQL

Basic web App implementing CRUD

Clone project, run Database, and install dependency
```bash
# Run database in container
docker-compose -f docker-compose.dev.yml up

# Install dependency
go mod download
```

Run server
```bash
go run main.go
```

[After running server swager docs are available](http://localhost:8080/swagger/index.html#)

### Additional info

[How to use swag with Gin](https://github.com/swaggo/swag)