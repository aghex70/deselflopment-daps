package cmd

import (
	"database/sql"
	"github.com/aghex70/daps/persistence/database"
	"github.com/spf13/cobra"
	"log"
)

func MakeMigrationsCommand(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "makemigrations [filename]",
		Short: "Generate database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatalf("fooooooooooo")
			}
			filename := args[0]
			err := database.MakeMigrations(db, filename)
			if err != nil {
				log.Fatalf("foooooooooooooo migrations")
			}
		},
	}
	return cmd

}
