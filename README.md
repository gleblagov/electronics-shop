# Бэкенд для магазина электроники

# .env
```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=mopevm3737
POSTGRES_HOST=db
POSTGRES_DB=postgres
```

# как поднять

API доступен на 3737 порту.

1. `git clone git@github.com:gleblagov/electronics-shop.git`
2. `cd electronics-shop`
3. создать файл `.env`
4. `docker compose --env-file=.env up --build -d`

# скрипты для теста

Скрипты для быстрого тестирования API лежат в `util`. 

## Использование

`./util/create-user.sh mail@example.com mopevm` — создание нового пользователя с электронной почтой **mail@example.com** и паролем **mopevm**.
`./util/delete-user.sh 5` — удаление пользователя с ID 5.
`./util/get-user.sh 5` — получение информации о пользователе с ID 5.

# todo
- [ ] логи
- [ ] все TODO
