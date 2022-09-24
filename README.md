# REST API 

gorilla - logrus - postgreSQL/pq - testify - gomail

# Installation

## Supported Versions

This library supports the following Go implementations:

* Go 1.14
* Go 1.15
* Go 1.16
* Go 1.17
* Go 1.18
* Go 1.19

## Install Package

```bash
go get github.com/t67y110v/REST
```

## Configuration setup

#### configs/apiserver.toml

```toml
[server]
bind_addr=":8081"
log_level="debug"

[database]
database_url="user=postgres password=p02tgre2 dbname=restapi sslmode=disable"

[smtp]
email_sender="email"
smtp_email="smtp.yandex.ru"
password_sender="password"
smtp_port=456
```

## Endpoints

| Name | Description |
|------|-------------|
| **/userCreate** | Creating a new user |
| **/userUpdate** | Updates user data |
| **/userDelete** | Deleting a user |
| **/sessionsr** | User authorization |
| **/makeAdmin** | Makes the user role admin |
| **/makeManager** | Makes the user role manager |
| **/changePassword** | Changes user password |
| **/departmentCondition** | User department status |
| **/sendEmail** | Sends email |

## Request | Responds
### Json models in jsons/requests | jsons/responds
