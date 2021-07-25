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

#### Viper

Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

[More about viper](https://github.com/spf13/viper)

#### Swag

Swag converts Go annotations to Swagger Documentation.

[How to use swag with Gin](https://github.com/swaggo/swag)

[Example of using swag](https://github.com/swaggo/swag/blob/master/example/celler/controller/examples.go)