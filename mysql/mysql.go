package mysql

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"gopkg.in/yaml.v2"
)

type (
	dbconfig struct {
		Dialect    string
		DataSource string
	}

	config struct {
		Development dbconfig
	}
)

var db *gorp.DbMap

func DB() *gorp.DbMap {
	return db
}

func NewDB() error {
	rawConf, err := ioutil.ReadFile("dbconfig.yml")
	if err != nil {
		return err
	}

	conf := &config{}
	if err = yaml.Unmarshal(rawConf, conf); err != nil {
		return err
	}

	pdb, err := sql.Open(conf.Development.Dialect, conf.Development.DataSource)
	if err != nil {
		return err
	}

	db = &gorp.DbMap{
		Db:      pdb,
		Dialect: &gorp.MySQLDialect{},
	}
	defer db.Db.Close()

	return nil
}
