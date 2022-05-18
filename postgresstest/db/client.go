package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

type DBConfig struct{
	ConnString     string
	Debug    bool
}

type DBClient struct{
	debugMode  bool
	connString string
	db *pgx.Conn
}

// EXPORTED

func NewDBClient(config DBConfig) *DBClient {
	cl := DBClient{
		connString: config.ConnString,
		debugMode:  config.Debug,
	}

	cl.debug("debug enabled")

	return &cl
}

func (cl *DBClient) Ping() error{
	err := cl.open()
	if err != nil{
		return cl.error("Ping", err)
	}

	/*defer cl.db.Close()

	err = cl.db.Ping()
	if err == nil{
		cl.debug("ping successful")
	}*/
	return err
}


// PRIVATE

func (cl *DBClient) open() error{
	conn, err := pgx.Connect(context.Background(), cl.connString)
	if err != nil {
		return err
	}

	cl.db = conn
	return nil
}

func (cl *DBClient) close(){

	err := cl.db.Close(context.Background())
	if err != nil {
		cl.error("close")
	}
}

func (cl *DBClient) debug(msg string, a ...interface{}){
	if cl.debugMode{
		args := ""
		for i, b := range a {
			if i == 0{
				args += " { "
			}
			if i != 0 {
				args += ", "
			}
			args += fmt.Sprintf("%+v", b)		// use #v or +v for more details
			if i == len(a)-1{
				args += " }"
			}
		}

		log.Printf("[ DBG ] [ DB ] %s%s\n",
			msg,
			args,
		)
	}
}

func (cl *DBClient) error(msg string, a ...interface{}) error{
	if cl.debugMode{
		args := ""
		for i, b := range a {
			if i == 0{
				args += " { "
			}
			if i != 0 {
				args += ", "
			}
			args += fmt.Sprintf("%+v", b)		// use #v or +v for more details
			if i == len(a)-1{
				args += " }"
			}
		}

		log.Printf("[ ERR ] [ DB ] %s%s\n",
			msg,
			args,
		)
	}

	return fmt.Errorf("ERR.DB{ %s }",msg)
}
