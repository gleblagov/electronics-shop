# initial commit =)

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
