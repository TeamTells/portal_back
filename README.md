Для запуска:
1. `go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest`
2. `oapi-codegen -generate gorilla,types authentication/impl/generated/frontendapi/frontendapi.yaml > authentication/impl/generated/frontendapi/frontendapi.gen.go`
3. `go mod tidy -v`
4. `go run main.go`

Configure GoLand run config with the following env vars: 
- DB_USER
- DB_PASSWORD
- DB_NAME
- DB_HOST
- BACKEND_PORT

![img.png](img/envVars.png)

Для работы с бд
1. установить postgres
2. создать две таблицы

```
CREATE TABLE auth_user
(
id serial primary key,
login character varying(256) NOT NULL,
salt character varying(256) NOT NULL,
password character varying(256) NOT NULL
);

CREATE TABLE tokens
(
id serial primary key,
user_id integer NOT NULL,
token character varying(256) NOT NULL
);