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
				setCorsHeaders(w)
				handler.ServeHTTP(w, r)
			}))
		}},
	}
	r := frontendapi.HandlerWithOptions(server, options)
	http.Handle("/authorization/", r)

	return authRequestService, conn
}

func methodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			setCorsHeaders(w)
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-user-id, X-organization-id")
	w.Header().Set("Access-Control-Allow-Origin", "https://dev4.env.teamtells.ru")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}