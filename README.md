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
