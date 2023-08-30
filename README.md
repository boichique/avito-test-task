# Тестовое задание Avito

## Задание
Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

Полное описание по ссылке - [тут](https://github.com/boichique/avito-test-task/blob/main/AvitoTask.md "тут")

## Необходимые инструменты для запуска сервиса
На компьютере должны быть установлены:
- Docker (с возможностью использования docker compose)
- go

## Команды Makefile
Запуск сервиса:
- `make service-up`

Остановка сервиса:
- `make service-down`

Форматирование, проверка линтерами и прогон тестов:
- `make before-push`

## Работа с сервисом
Сервис стартует без данных, так что сначала необходимо заполнить базу пользователями и сегментами. Для запуска запросов можно использовать postman, swagger или curl. Ниже приведены примеры запросов для postman и swagger:


#### Postman:
Создание пользователя
![CreateUserPostman](https://github.com/boichique/avito-test-task/assets/87061629/5ef82422-b2c1-460b-90b0-a791c8b345bb)


Удаление пользователя
![DeleteUserPostman](https://github.com/boichique/avito-test-task/assets/87061629/6b5560a6-8c9c-43e2-ae63-2ff91891fc69)


Создание сегмента
![CreateSegmentPostman](https://github.com/boichique/avito-test-task/assets/87061629/bc947a69-5a76-4bfa-b835-aade19708490)


Удаление сегмента
![DeleteSegmentPostman](https://github.com/boichique/avito-test-task/assets/87061629/00c13755-09a3-4f8c-92d4-c1c7fc35ef0e)


Изменение сегментов пользователя
![UpdateUserSegmentsPostman](https://github.com/boichique/avito-test-task/assets/87061629/944b2029-7e7b-4ee8-a338-8bca71593620)


Получение сегментов пользователя
![GetUserSegmentsPostman](https://github.com/boichique/avito-test-task/assets/87061629/92126163-e7fa-49bd-b6d7-fc8138ba4225)


#### Swagger:
Также все запросы можно прогнать и через Swagger
![Swagger](https://github.com/boichique/avito-test-task/assets/87061629/5fd822cb-5dff-4647-9459-d98294a4b25d)


URL для подключения после запуска сервиса - [тут](http://localhost:8080/swagger/index.html "тут")
