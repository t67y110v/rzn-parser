# REST API 

fiber - swagger - logrus - postgreSQL/pq - testify - gomail

## Swagger documentaion on http://localhost:4000/swagger/index.html#/

## Configuration setup

#### configs/apiserver.toml

```toml
[server]
bind_addr=":8080"
log_level="debug"

[database]
database_url="user=postgres password=p02tgre2 dbname=restapi sslmode=disable"

[smtp]
email_sender="email"
smtp_email="smtp.yandex.ru"
password_sender="password"
smtp_port=456
```


