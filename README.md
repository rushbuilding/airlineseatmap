Install Postgres Using Docker

`docker-compose up`

Install Golang Migrate to create database relation

`go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

Create table in database
`migrate -database "mysql://seat:123456abc@tcp(localhost:3306)/seatmapdatabase" -path db/migrations up`

-database is database created from docker

This app hasn't finish so app cannot be run but to run it, it can use

`go run main.go`

ERD for seating database
![image](https://github.com/user-attachments/assets/e72a674f-dabf-4a0a-8de6-f59b2334d4ae)
