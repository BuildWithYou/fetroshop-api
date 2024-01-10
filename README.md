FETROSHOP API
=================
![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)

Fetroshop API is a robust and efficient backend solution designed to power the online store app named Fetroshop. Developed using the Go programming language, this API serves as the backbone for managing the Content Management System (CMS) and handling various store-related functionalities.

Lorem ipsum dolor sit amet consectetur adipisicing elit. Eum quam aut at. Nostrum ut id mollitia at quidem debitis iste tempore culpa, tenetur perferendis porro impedit est molestias laborum accusamus? Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ipsum repellendus architecto, tempore eius incidunt nam ex recusandae! Sint officia saepe animi quis odio consequuntur expedita sed ipsa laudantium! Expedita, illum. Lorem ipsum dolor sit amet consectetur adipisicing elit. Aspernatur, ab. Eos cupiditate velit, minima voluptates quae enim delectus iste rem similique et autem est voluptas ipsum ratione ipsa magni nemo.

Generate Docs File
=================

For Web
```
   swag init -d "app/modules/web,app/model" -g web.go -o docs/web
```

For CMS
```
   swag init -d "app/modules/cms,app/model" -g cms.go -o docs/cms
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

## Docker
1. Pull this repository
2. Build docker image using this command
   ```
   docker builder prune
   docker-compose up --force-recreate
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