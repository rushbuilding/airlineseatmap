Install Postgres Using Docker

`docker-compose up`

Install Golang Migrate to create database relation

`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

Create table in database
`migrate -database "mysql://seat:123456abc@tcp(localhost:3306)/seatmapdatabase" -path db/migrations up`

-database is database created from docker

Change database url on https://github.com/rushbuilding/airlineseatmap/blob/main/db/connections/database.go
Use database created

This app hasn't finish so app cannot be run but to run it, it can use

`go run main.go`

ERD for seating database
![erd](https://github.com/user-attachments/assets/5ddebaf8-1ec6-499f-a237-9a27f5c58a30)

