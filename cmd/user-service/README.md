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
- Применить миграцию goose -dir ./migrations postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable" up
- откатить миграцию goose -dir ./migrations postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable" down
