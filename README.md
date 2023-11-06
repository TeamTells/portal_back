Для запуска:
1. `go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen`
2. `go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen`
3. `oapi-codegen -generate gorilla,types api/frontendapi/frontendapi.yaml > api/frontendapi/frontendapi.gen.go`
4. `go mod tidy -v`
5. `go run main.go`

Configure GoLand run config with the following env vars: 
- DB_USER
- DB_PASSWORD
- DB_NAME
- DB_HOST
- BACKEND_PORT

![img.png](img/envVars.png)
