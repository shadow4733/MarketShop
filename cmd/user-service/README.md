# Сервис предназначен для работы с пользователями

### Данный сервис будет использовать следующие технологии:


Зависимости которые будут использоваться
- gorm
- godotenv
- postgres
- uuid
- godoc
- goose
Сгенирировать файл миграции  goose create create_users_table sql
- Применить миграцию goose -dir ./migrations postgres "host=localhost port=5433 user=postgres password=secretpassw0rd dbname=user-service sslmode=disable" up
- откатить миграцию goose -dir ./migrations postgres "host=localhost port=5433 user=postgres password=secretpassw0rd dbname=user-service sslmode=disable" down
- gorm/postgresql
- go get -u github.com/swaggo/swag/cmd/swag
  go get -u github.com/swaggo/gin-swagger
  go get -u github.com/swaggo/files

Инициализация документации swag init -g main.go из того места где main(cmd)