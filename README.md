# GoRestFullApi

## Создание бд на основе postgresql:

#### docker run --name=restfull-api -e POSTGRESS_USER=postgres -e POSTGRES_PASSWORD=root -e POSTGRES_DB=restfull_api -p 5436:5432 -d postgres:latest

## Установка миграций, если необходимо:

#### go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.0
#### migrate create -ext sql -dir migrations -seq init_schema