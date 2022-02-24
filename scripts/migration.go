/*
 * @Author: Adrian Faisal
 * @Date: 02/09/21 8.33 PM
 */

package main

import (
	"database/sql"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MigrationConfig struct {
	Local struct {
		Dialect    string `yaml:"dialect"`
		Datasource string `yaml:"datasource"`
		Table      string `yaml:"table"`
		Dir        string `yaml:"dir"`
	} `yaml:"local"`
}

func main() {
	data, err := ioutil.ReadFile("dbconfig.yml")
	if err != nil {
		log.Fatalf("read configs file failed: %v", err)
	}

	var cfg MigrationConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("failed unmarshal yaml: %v", err)
	}

	migrations := &migrate.FileMigrationSource{Dir: cfg.Local.Dir}

	db, err := sql.Open("mysql", cfg.Local.Datasource)
	if err != nil {
		log.Fatalf("open sql connection error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("ping db error: %v", err)
	}

	_, err = migrate.Exec(db, cfg.Local.Dialect, migrations, migrate.Up)
	if err != nil {
		log.Fatalf("migration error: %v", err)
	}
}
