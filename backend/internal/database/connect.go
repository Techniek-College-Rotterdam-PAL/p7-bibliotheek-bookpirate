package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"server/internal/util"
)

func RunDriver() error {
	config := util.LoadConfigFile()
	fmt.Println(config)
	db, err := sql.Open(config.Database.Driver, config.Database.Dsn)
	if err != nil {
		return err
	}
	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	fmt.Println(conn)
	return nil
}
