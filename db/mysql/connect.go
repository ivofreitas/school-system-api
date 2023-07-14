package mysql

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leantech/school-system-api/config"
	"github.com/leantech/school-system-api/log"
)

var (
	once sync.Once
	conn *sql.DB
)

func GetConn() *sql.DB {
	once.Do(func() {
		mySQL := config.GetEnv().MySQL
		logger := log.NewEntry()

		var err error
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", mySQL.Username, mySQL.Password, mySQL.Host, mySQL.Database)
		conn, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			logger.WithError(err).Fatal()
		}

		if err = conn.Ping(); err != nil {
			logger.WithError(err).Fatal()
		}

		conn.SetMaxIdleConns(mySQL.PoolConn)
		conn.SetConnMaxLifetime(mySQL.ConnLifetime)
	})

	return conn
}
