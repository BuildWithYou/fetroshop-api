FETROSHOP API
=================
![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)

Fetroshop API is a robust and efficient backend solution designed to power the online store app named Fetroshop. Developed using the Go programming language, this API serves as the backbone for managing the Content Management System (CMS) and handling various store-related functionalities.

Lorem ipsum dolor sit amet consectetur adipisicing elit. Eum quam aut at. Nostrum ut id mollitia at quidem debitis iste tempore culpa, tenetur perferendis porro impedit est molestias laborum accusamus? Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ipsum repellendus architecto, tempore eius incidunt nam ex recusandae! Sint officia saepe animi quis odio consequuntur expedita sed ipsa laudantium! Expedita, illum. Lorem ipsum dolor sit amet consectetur adipisicing elit. Aspernatur, ab. Eos cupiditate velit, minima voluptates quae enim delectus iste rem similique et autem est voluptas ipsum ratione ipsa magni nemo.

API Documentation
=================
We use [swag](https://github.com/swaggo/swag) to generate API documentation. Here is commands you can use to generate documentation.
You can access the documentation in `http://localhost:3000/documentation` for web api module documentation and `http://localhost:3001/documentation` for cms api module documentation.

For Web
```
   swag init -d "app/modules/web,app/model" -g web.go -o docs/openapi2/web
```

For CMS
```
   swag init -d "app/modules/cms,app/model" -g cms.go -o docs/openapi2/cms
```

Migration
=================
## A. Install Migration
```
   go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
## B. Create Migration
```
   migrate create -ext sql -dir db/migrations migration_file_name
```
## C. Runing Migration
We use [golang-migrate](https://github.com/golang-migrate/migrate) to sync database.

Migration Up
```
   migrate -database "postgresql://{username}:{password}@{host}:{port}/{database}" -path db/migrations up
```
Migration down
```
   migrate -database "postgresql://{username}:{password}@{host}:{port}/{database}" -path db/migrations down
```

Installation
=================
Lorem ipsum dolor sit amet consectetur adipisicing elit. Eum quam aut at. Nostrum ut id mollitia at quidem debitis iste tempore culpa, tenetur perferendis porro impedit est molestias laborum accusamus? Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ipsum repellendus architecto, tempore eius incidunt nam ex recusandae! Sint officia saepe animi quis odio consequuntur expedita sed ipsa laudantium! Expedita, illum. Lorem ipsum dolor sit amet consectetur adipisicing elit. Aspernatur, ab. Eos cupiditate velit, minima voluptates quae enim delectus iste rem similique et autem est voluptas ipsum ratione ipsa magni nemo.

## Local
1. Pull this repository
2. Copy `config.yaml.example` then rename it to `config.yaml`
3. Setup postgre database and minio and update config.yaml file
4. Execute `go run .`

## Docker
1. Pull this repository
2. Copy `.env.example` file and rename it to `.env`
3. Build docker image using this command
   ```
   docker-compose up
   ```

Testing
=================
```
go test ./... -coverprofile cover.out
```

Usage and documentation
=================
Lorem ipsum dolor sit amet consectetur adipisicing elit. Eum quam aut at. Nostrum ut id mollitia at quidem debitis iste tempore culpa, tenetur perferendis porro impedit est molestias laborum accusamus? Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ipsum repellendus architecto, tempore eius incidunt nam ex recusandae! Sint officia saepe animi quis odio consequuntur expedita sed ipsa laudantium! Expedita, illum. Lorem ipsum dolor sit amet consectetur adipisicing elit. Aspernatur, ab. Eos cupiditate velit, minima voluptates quae enim delectus iste rem similique et autem est voluptas ipsum ratione ipsa magni nemo.

How to Contribute
=================
Make a pull request...