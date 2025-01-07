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

### util/user

`./create-user.sh mail@example.com mopevm $роль` — создание нового пользователя с электронной почтой **mail@example.com** и паролем **mopevm**. $роль = {admin/client/seller}.

`./delete-user.sh 5` — удаление пользователя с ID 5.

`./get-user.sh 5` — получение информации о пользователе с ID 5.

`./update-user.sh 5 new-mail@example.com new-password admin` — обновление существующего пользователя с ID **5**. Задаётся новая почта, новый пароль, роль **администратор**.


### util/product

`./create-product.sh мышка 1200 100 периферия 4` — создание нового товара с названием **мышка**, ценой **1200**, кол-вом на складе **100**, категорией **периферия** и рейтингом **4**.

`./delete-product.sh 5` — удаление товара с ID 5.

`./get-product.sh 5` — получение информации о товаре с ID 5.
