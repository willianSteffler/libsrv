package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"strconv"
	"time"
)

type DbConf struct {
	driver string
	connStr string
	connIdle int
	connLife int
	connOpen int
	logMode bool
}

type Conf struct{
	socketPort int
	DbConf
}

var conf Conf

func main (){

	flag.IntVar(&conf.socketPort,"soquet",3000,"porta socket")

	flag.StringVar(&conf.driver,"driver","sqlite","driver do banco de dados")
	flag.IntVar(&conf.connLife,"connLife",60,"tempo de vida da conexão com o banco em minutos")
	flag.IntVar(&conf.connIdle,"connIdle",5,"maximo de conexões com o banco em estado idle")
	flag.IntVar(&conf.connOpen,"connOpen",10,"maximo de conexões com o banco")

	Usage := func() {
		fmt.Fprintf(os.Stderr, "Sintaxe do comando %s: dbConnStr [opções]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "dbConnStr: string de conexão do banco de dados ou arquivo .dbzn")
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

	db,err := connectDb()
	if err != nil {
		log.Fatalf("não foi possível conectar ao banco de dados %v!!",err)
	}

}

func connectDb() (*gorm.DB,error) {

	var db *gorm.DB
	var err error

	cfg := conf.DbConf

	if cfg.driver == "sqlite" {
		db, err = gorm.Open("sqlite3", cfg.connStr)
	} else if cfg.driver == "postgres" {
		db, err = gorm.Open("postgres", cfg.connStr)
	} else {
		err = errors.New("driver deve ser sqlite ou postgres")
	}

	if err == nil{
		db.DB().SetMaxIdleConns(cfg.connIdle)
		db.DB().SetMaxOpenConns(cfg.connOpen)
		db.DB().SetConnMaxLifetime(time.Duration(cfg.connLife) * time.Minute)
		db.LogMode(cfg.logMode)
	}

	return db,err

}