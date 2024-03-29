package cmd

import (
	"database/sql"
	"log"

	"github.com/aghex70/daps/persistence/database"
	"github.com/spf13/cobra"
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
			if err := database.MakeMigrations(db, filename); err != nil {
				log.Fatalf("foooooooooooooo migrations")
			}
		},
	}
	return cmd

}
