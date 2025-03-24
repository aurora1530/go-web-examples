package internal

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Server struct {
	DB *sql.DB
	Store *sessions.CookieStore
}

func openDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := "localhost"
	port := "54321"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createSessionStore() *sessions.CookieStore {
	key := os.Getenv("SESSION_KEY")
	if key == "" {
		panic("SESSION_KEY is not set")
	}

	store := sessions.NewCookieStore([]byte(key))

	return store
}

func NewServer() Server {
	db, err := openDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database")

	store := createSessionStore()

	return Server{DB: db, Store: store}
}

func CreateRouter(server Server) *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/api/auth/create", server.PostApiCreateUser).Methods("POST")
	r.HandleFunc("/api/auth/login", server.PostApiLoginUser).Methods("POST")

	return r
}