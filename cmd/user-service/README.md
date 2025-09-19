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
- Применить миграцию goose -dir ./migrations postgres "host=localhost port=5433 user=user password=secretpassw0rd dbname=user-service sslmode=disable" up
- откатить миграцию goose -dir ./migrations postgres "host=localhost port=5433 user=user password=secretpassw0rd dbname=user-service sslmode=disable" down
