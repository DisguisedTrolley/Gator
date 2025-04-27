package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/DisguisedTrolley/gator/internal/config"
	"github.com/DisguisedTrolley/gator/internal/database"
	_ "github.com/lib/pq"
)

const dbUrl = "postgres://samarthbhat:@localhost:5432/gator"

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfgFile, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("unable to connect to database")
	}

	dbQueries := database.New(db)

	s := state{
		config: &cfgFile,
		db:     dbQueries,
	}

	cmd := commands{
		cmd: make(map[string]func(*state, command) error),
	}
	cmd.register("login", handlerLogin)

	// Read cmd args
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Insufficient arguments")
	}

	cmdName := args[1]
	arguments := args[2:]

	err = cmd.run(&s, command{name: cmdName, args: arguments})
	if err != nil {
		log.Fatal(err)
	}
}
