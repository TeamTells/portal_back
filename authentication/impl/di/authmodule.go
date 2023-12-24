package di

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
	"portal_back/authentication/api/frontend"
	"portal_back/authentication/api/internalapi"
	"portal_back/authentication/impl/app/auth"
	"portal_back/authentication/impl/app/authrequest"
	"portal_back/authentication/impl/app/token"
	"portal_back/authentication/impl/infrastructure/sql"
	"portal_back/authentication/impl/infrastructure/transport"
)

func InitAuthModule() (internalapi.AuthRequestService, *pgx.Conn) {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "app"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)

	conn, _ := pgx.Connect(context.Background(), connStr)

	repo := sql.NewTokenStorage(conn)
	tokenService := token.NewService(repo)

	authRepo := sql.NewAuthRepository(conn)
	authService := auth.NewService(authRepo, tokenService)
	server := transport.NewServer(authService, tokenService)
	authRequestService := authrequest.NewService()

	router := mux.NewRouter()
	router.MethodNotAllowedHandler = methodNotAllowedHandler()

	options := frontendapi.GorillaServerOptions{
		BaseRouter: router,
		Middlewares: []frontendapi.MiddlewareFunc{func(handler http.Handler) http.Handler {
			return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				println("1111111")
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				handler.ServeHTTP(w, r)

				//if r.Method == http.MethodOptions {
				//	println("333333")
				//
				//	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
				//	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				//	w.Header().Set("Access-Control-Allow-Credentials", "true")
				//	_, _ = w.Write([]byte("OK"))
				//} else {
				//	handler.ServeHTTP(w, r)
				//}
			}))
		}},
	}
	r := frontendapi.HandlerWithOptions(server, options)
	http.Handle("/authorization/", r)

	return authRequestService, conn
}

func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("4444444")

		if r.Header.Get("Access-Control-Request-Method") != "" {
			println("2222222")

			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			//w.Header().Set("Access-Control-Max-Age", "86400")
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func methodNotAllowedHandler2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("4444444")

		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			println("333333")

			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			_, _ = w.Write([]byte("OK"))
		}
	})
}
