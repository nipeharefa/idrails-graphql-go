package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nipeharefa/idrails-graphql-go/graph"
	"github.com/nipeharefa/idrails-graphql-go/graph/generated"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const defaultPort = "8080"

func init() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sqlx.Open("postgres", viper.GetString("application.db.url"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(db)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
