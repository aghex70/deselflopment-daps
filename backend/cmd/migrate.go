package cmd

import (
	"database/sql"
	"log"

	"github.com/aghex70/daps/persistence/database"
	"github.com/spf13/cobra"
)

func MigrateCommand(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Apply database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if err := database.Migrate(db); err != nil {
				log.Fatalf("error applying migrations %+v", err.Error())
			}
		},
	}
	return cmd
}
