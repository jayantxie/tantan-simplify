package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"tantan-simplify/config"
)

// TODO: connect goroutine numbers ?
var db *pg.DB

func MustSetDB(postgreSQLConfig *config.PostgreSQL) error {
	db = pg.Connect(&pg.Options{
		Addr:     postgreSQLConfig.Addr,
		User:     postgreSQLConfig.User,
		Password: postgreSQLConfig.Password,
		Database: postgreSQLConfig.DBName,
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func GetDB() *pg.DB {
	return db
}

func CreateTable(db *pg.DB) error {
	for _, model := range []interface{}{&Users{}, &Relationships{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
