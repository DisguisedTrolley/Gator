package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DisguisedTrolley/gator/internal/config"
	"github.com/DisguisedTrolley/gator/internal/database"
	_ "github.com/lib/pq"
)

const dbUrl = "postgres://samarthbhat:@localhost:5432/gator?sslmode=disable"

type state struct {
	config *config.Config
	db     *database.Queries
}

func newState() (*state, error) {
	cfgFile, err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database")
	}

	dbQueries := database.New(db)

	return &state{
		config: &cfgFile,
		db:     dbQueries,
	}, nil
}

func main() {
	progState, err := newState()
	if err != nil {
		log.Fatal(err)
	}

	// Register commands
	cmd := commands{
		cmd: make(map[string]func(*state, command) error),
	}
	cmd.register("login", handlerLogin)
	cmd.register("register", handlerRegister)
	cmd.register("reset", handlerReset)
	cmd.register("users", handlerListUsers)

	// Deal with arguments
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := args[1]
	arguments := args[2:]

	err = cmd.run(progState, command{name: cmdName, args: arguments})
	if err != nil {
		log.Fatal(err)
	}
}
