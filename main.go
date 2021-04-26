package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/willianSteffler/libsrv/data"
	"github.com/willianSteffler/libsrv/socket"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DbConf struct {
	driver   string
	connStr  string
	connIdle int
	connLife int
	connOpen int
	logLevel logger.LogLevel
}

type Conf struct {
	socketPort int
	DbConf
}

var conf Conf

func main() {

	flag.IntVar(&conf.socketPort, "socket", 3000, "porta socket")
	flag.StringVar(&conf.driver, "driver", "sqlite", "driver do banco de dados")
	flag.IntVar(&conf.connLife, "connLife", 60, "tempo de vida da conexão com o banco em minutos")
	flag.IntVar(&conf.connIdle, "connIdle", 5, "maximo de conexões com o banco em estado idle")
	flag.IntVar(&conf.connOpen, "connOpen", 10, "maximo de conexões com o banco")
	flag.IntVar((*int)(&conf.logLevel), "dblog", (int)(logger.Silent), "nivel de log do banco "+
		"\n\tSilent = 1"+
		"\n\tError  = 2"+
		"\n\tWarn   = 3"+
		"\n\tInfo   = 4")

	Usage := func() {
		fmt.Fprintf(os.Stderr, "Sintaxe do comando %s: [opções] dbConnStr\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "** dbConnStr: string de conexão do banco de dados ou arquivo .db\n")
		flag.PrintDefaults()
	}

	flag.Usage = Usage
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	conf.DbConf.connStr = args[0]

	db, err := connectDb()
	if err != nil {
		log.Fatalf("não foi possível conectar ao banco de dados %v!!", err)
	}

	if err = socket.Listen(conf.socketPort,db); err != nil {
		log.Fatalf("erro no socket %v!!", err)
	}

}

func connectDb() (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	cfg := conf.DbConf
	gormCfg := gorm.Config{
		Logger: logger.Default.LogMode(cfg.logLevel),
	}

	if cfg.driver == "sqlite" {
		db, err = gorm.Open(sqlite.Open(cfg.connStr), &gormCfg)
	} else if cfg.driver == "postgres" {
		db, err = gorm.Open(postgres.Open(cfg.connStr), &gormCfg)
	} else {
		err = errors.New("driver deve ser sqlite ou postgres")
	}

	if err == nil {
		var sqlDb *sql.DB
		sqlDb, err = db.DB()
		if err == nil {
			sqlDb.SetMaxOpenConns(cfg.connOpen)
			sqlDb.SetMaxIdleConns(cfg.connIdle)
			sqlDb.SetConnMaxLifetime(time.Duration(cfg.connLife) * time.Minute)
		}
	}


	err = db.AutoMigrate(
		&data.Livro{},
		&data.Autor{},
		&data.Edicao{},
		)

	return db, err

}
