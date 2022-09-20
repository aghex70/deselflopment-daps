package cmd

import (
	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/persistence/database"
	"github.com/spf13/cobra"
	"log"
)

func RootCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daps",
		Short: "Root command",
	}

	// Intialize database
	log.Println("Starting application database")
	db, err := database.NewSqlDB(*cfg.Database)
	if err != nil {
		log.Fatalf("error starting application database %+v", err.Error())
	}

	cmd.AddCommand(ServeCommand(cfg))
	cmd.AddCommand(MigrateCommand(db))
	cmd.AddCommand(MakeMigrationsCommand(db))
	//cmd.AddCommand(RollbackCommand(db))
	return cmd
}
