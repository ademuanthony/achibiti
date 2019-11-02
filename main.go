// Copyright (c) 2018-2019 The Achibiti developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	acl "github.com/ademuanthony/achibiti/acl/proto/acl"
	"github.com/ademuanthony/achibiti/app"
	"github.com/ademuanthony/achibiti/app/config"
	"github.com/ademuanthony/achibiti/app/help"
	"github.com/ademuanthony/achibiti/app/helpers"
	"github.com/ademuanthony/achibiti/postgres"
	"github.com/ademuanthony/achibiti/web"
	"github.com/jessevdk/go-flags"
)

func main() {
	// Create a context that is cancelled when a shutdown request is received
	// via requestShutdown.
	ctx := withShutdownCancel(context.Background())
	// Listen for both interrupt signals and shutdown requests.
	go shutdownListener()

	if err := _main(ctx); err != nil {
		if logRotator != nil {
			log.Error(err)
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func _main(ctx context.Context) error {
	cfg, args, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize log rotation.  After log rotation has been initialized, the
	// logger variables may be used.
	initLogRotator(cfg.ConfigFileOptions.LogFile)
	defer func() {
		if logRotator != nil {
			logRotator.Close()
		}
	}()

	// Special show command to list supported subsystems and exit.
	if cfg.DebugLevel == "show" {
		fmt.Println("Supported subsystems", supportedSubsystems())
		os.Exit(0)
	}

	// Parse, validate, and set debug log level(s).
	if cfg.Quiet {
		cfg.ConfigFileOptions.DebugLevel = "error"
	}

	// Parse, validate, and set debug log level(s).
	if err := parseAndSetDebugLevels(cfg.DebugLevel); err != nil {
		err := fmt.Errorf("loadConfig: %s", err.Error())
		return err
	}

	// if len(args) == 0, then there's nothing to execute as all command-line args were parsed as app options
	if len(args) > 0 {
		err := executeHelpCommand()
		if err != nil {
			return fmt.Errorf("%s: %s", err, config.Hint)
		}
		return nil
	}

	db, err := postgres.NewPgDb(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	if err != nil {
		return fmt.Errorf("error in establishing database connection: %s", err.Error())
	}

	defer func(db *postgres.PgDb) {
		err := db.Close()
		if err != nil {
			log.Error("Could not close database connection: %v", err)
		}
	}(db)

	if cfg.Reset {
		resetTables, err := helpers.RequestYesNoConfirmation("Are you sure you want to reset the dcrextdata db?", "")
		if err != nil {
			return fmt.Errorf("error reading your response: %s", err.Error())
		}

		if resetTables {
			err = db.DropAllTables()
			if err != nil {
				db.Close()
				log.Error("Could not drop tables: ", err)
				return err
			}

			fmt.Println("Done. You can restart the server now.")
			return nil
		}

		return nil
	}

	// Display app version.
	log.Infof("%s version %v (Go version %s)", app.AppName, app.Version(), runtime.Version())

	if err = createTablesAndIndex(db); err != nil {
		return err
	}

	// http server method
	go web.StartHTTPServer(cfg.HTTPHost, cfg.HTTPPort, db, acl.NewAclService("go.micro.srv.acl", nil))

	// wait for shutdown signal
	<-ctx.Done()

	return ctx.Err()
}

// executeHelpCommand checks if the operation requested by the user is -h, --help flags. If it not a help flag is throw an error.
func executeHelpCommand() (err error) {
	configWithCommands := &config.Config{}
	parser := flags.NewParser(configWithCommands, flags.HelpFlag|flags.PassDoubleDash)

	// re-parse command-line args to catch help flag or execute any commands passed
	_, err = parser.Parse()
	if err != nil {
		e, ok := err.(*flags.Error)
		if ok && e.Type == flags.ErrHelp {
			help.PrintGeneralHelp(os.Stdout, help.HelpParser())
			return nil
		}
		return err
	}

	return fmt.Errorf(config.Hint)
}

func createTablesAndIndex(db *postgres.PgDb) error {

	return nil
}
