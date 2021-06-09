package main

import (
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Akshit8/tdm/internal/env"
	"github.com/Akshit8/tdm/internal/env/vault"
	"github.com/Akshit8/tdm/internal/postgresql"
	"github.com/Akshit8/tdm/internal/rest"
	"github.com/Akshit8/tdm/internal/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

//go:embed static
var content embed.FS

func main() {
	var envFilePath, address string

	flag.StringVar(&envFilePath, "env", "", "Environment variable file path")
	flag.StringVar(&address, "address", "0.0.0.0:8000", "Server address")
	flag.Parse()

	err := env.Load(envFilePath)
	if err != nil {
		log.Fatalln("Couldn't load configuration", err)
	}

	config := env.NewConfiguration(newVaultProvider())

	db := newDB(config)
	defer db.Close()

	repo := postgresql.NewTask(db)
	svc := service.NewTask(repo)

	r := mux.NewRouter()

	rest.RegisterOpenAPI(r)
	rest.NewTaskHandler(svc).Register(r)

	fsys, _ := fs.Sub(content, "static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(fsys))))

	// cors enabled
	handler := cors.AllowAll().Handler(r)

	srv := &http.Server{
		Handler:           handler,
		Addr:              address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}

	log.Println("Starting server at", address)

	log.Fatal(srv.ListenAndServe())
}

func newDB(config *env.Configuration) *sql.DB {
	get := func(key string) string {
		val, err := config.Get(key)
		if err != nil {
			log.Fatalf("Couldn't get configuration vaules for %s: %s", key, err)
		}

		return val
	}

	databaseHost := get("DATABASE_HOST")
	databasePort := get("DATABASE_PORT")
	databaseUsername := get("DATABASE_USERNAME")
	databasePassword := get("DATABASE_PASSWORD")
	databaseName := get("DATABASE_NAME")
	databaseSSLMode := get("DATABASE_SSLMODE")

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%s", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		log.Fatalln("Couldn't open DB", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Couldn't ping DB", err)
	}

	return db
}

func newVaultProvider() env.Provider {
	vaultPath := os.Getenv("VAULT_PATH")
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddress := os.Getenv("VAULT_ADDRESS")

	provider, err := vault.NewVaultProvider(vaultToken, vaultAddress, vaultPath)
	if err != nil {
		log.Fatalln("Couldn't load vault provider", err)
	}

	return provider
}
