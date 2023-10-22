Для запуска:
1. `go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen`
2. `oapi-codegen -generate gorilla,types api/frontendapi/frontendapi.yaml > api/frontendapi/frontendapi.gen.go`
3. `go mod tidy -v`
4. `go run main.go`
