# Сервис транзакций

## Спринты

### Спринт 1

-   [x] Валидация токена (через secret)
-   [x] Подготовка инфраструктуры и технологии для реализации

### Спринт 2

-   [x] Начать делать сервис кошельков (настройка окружения, подключение к бд).
-   [x] Написать endpoint для создания кошелька у пользователя по ИИН. Изначальный баланс кошелька, понятное дело, равен 0. У каждого кошелька должен быть свой номер счета для совершения переводов.

### Спринт 3

-   [x] Создать endpoint для пополнения кошелька
-   [x] Добавить endpoint для получения информации о кошельках пользователя (номер счета, дата последней транзакции, баланс, дата создания кошелька)
-   [x] Создать endpoint для перевода денег. Пока что с минимальными проверками (авторизация юзера, принадлежность исходящего кошелька авторизованному юзеру и т.д.)

### Спринт 4

-   [x] Обернуть приложение в Docker
-   [x] Запускать оба сервиса через docker compose
-   [x] Написать unit тесты для сервиса транзакций

## Использованные технологии

-   [PostgreSQL](https://www.postgresql.org/)
-   [Redis](https://redis.io/)
-   [Docker](https://www.docker.com/)
-   [GORM](https://gorm.io/)
-   [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
-   [fasthttp](https://github.com/valyala/fasthttp)
-   [fasthttprouter](https://github.com/buaazp/fasthttprouter)
-   [sqlmock](https://github.com/DATA-DOG/go-sqlmock)
-   [redismock](https://github.com/go-redis/redismock)
-   [testify](https://github.com/stretchr/testify)
