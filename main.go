package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ksabanty/gator/internal/config"
	"github.com/ksabanty/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	// store config in state struct
	programState := &State{cfg: &cfg, db: dbQueries}

	cmds := Commands{cmdMap: make(map[string]func(*State, Command) error)}
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handlerRegister)
	cmds.Register("reset", handlerReset)
	cmds.Register("users", handlerUsers)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(programState, Command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
