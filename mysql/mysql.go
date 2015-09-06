package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//
// conn defines a connection container that we can use to check for an active connection.
type conn struct {
	open bool
	obj  *sql.DB
}

//
// myconn is our shared connection object.
var myconn conn

//
// Open returns the shared connection object. If the object has not yet been initialized,
// it is initialized before being returned.
func Open() (*sql.DB, error) {

	//
	// Check for an open connection
	if !myconn.open {

		//
		// Open a new connection
		newconn, err := sql.Open("mysql", os.ExpandEnv("${MYSQL_ENV_MYSQL_USER}:${MYSQL_ENV_MYSQL_PASSWORD}@tcp(${MYSQL_PORT_3306_TCP_ADDR}:${MYSQL_PORT_3306_TCP_PORT})/${MYSQL_ENV_MYSQL_DATABASE}?autocommit=true&parseTime=true"))
		if err != nil {
			return nil, err
		}
		myconn.obj = newconn

		//
		// Verify the connection
		if err = myconn.obj.Ping(); err != nil {
			return nil, err
		}

		//
		// Save the connection state
		myconn.open = true

	}

	//
	// Return the connection
	return myconn.obj, nil

}

//
// Close closes the shared connection, if the connection is currently open.
func Close() {
	if myconn.open {
		myconn.obj.Close()
		myconn.open = false
	}
}
