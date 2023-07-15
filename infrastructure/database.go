package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"university-api/domain/repository"

	_ "github.com/lib/pq"
)

const DATABASE_DRIVER_PGSQL string = "pgsql"

type Database struct {
	PgSqlInstance       string
	PgSqlDns            string
	PgSqlUser           string
	PgSqlPass           string
	PgSqlDbname         string
	GeneralDatabase     *sql.DB
	ApiFacultyInterface repository.ApiFacultyInterface
}

func NewDatabase(
	pgSqlInstance string,
	pgSqlDns string,
	pgSqlUser string,
	pgSqlPass string,
	pgSqlDbname string,
) repository.DatabaseInterface {
	db := &Database{
		PgSqlInstance: pgSqlInstance,
		PgSqlDns:      pgSqlDns,
		PgSqlUser:     pgSqlUser,
		PgSqlPass:     pgSqlPass,
		PgSqlDbname:   pgSqlDbname,
	}
	db.ApiFacultyInterface = NewFacultyRepository(db)
	return db
}

type Transaction struct {
	tx *sql.Tx
}

func NewTransaction(
	tx *sql.Tx,
) repository.TransactionInterface {
	return &Transaction{
		tx: tx,
	}
}

func (t *Transaction) Get() any {
	return t.tx
}

func (t *Transaction) Commit() error {
	return t.tx.Commit()
}

func (t *Transaction) Rollback() error {
	return t.tx.Rollback()
}

func (db *Database) Begin() (tx repository.TransactionInterface, err []error) {
	if db.GeneralDatabase != nil {
		fmt.Println("Database general exist")
	} else {
		errConnectSQL := db.ConnectToPGSQL()
		if errConnectSQL != nil {
			return nil, errConnectSQL
		}
	}
	txDb, begin_err := db.GeneralDatabase.Begin()
	if begin_err != nil {
		return nil, append(err, errors.New(begin_err.Error()))
	}
	tx = NewTransaction(txDb)
	return tx, nil
}

func (db *Database) Close() {
	if db.GeneralDatabase != nil {
		fmt.Println("Close general database connection")
		db.GeneralDatabase.Close()
		db.GeneralDatabase = nil
	}
}

func (db *Database) Commit(tx repository.TransactionInterface) (err []error) {
	err = []error{}
	t := tx.Get().(*sql.Tx)
	commit_err := t.Commit()
	if commit_err != nil {
		return append(err, errors.New(commit_err.Error()))
	}
	return nil
}

func (db *Database) Rollback(tx repository.TransactionInterface) (err []error) {
	err = []error{}
	t := tx.Get().(*sql.Tx)
	rollback_err := t.Rollback()
	if rollback_err != nil {
		return append(err, errors.New(rollback_err.Error()))
	}
	return nil
}
func (db *Database) ConnectToPGSQL() (err []error) {
	var (
		pgsql_user     string
		pgsql_pass     string
		pgsql_dbname   string
		pgsql_dns      string
		errGeneralOpen error
		errGeneralPing error
	)
	// pgsql_user = "teywcjxi"
	// pgsql_pass = "dAy_tVmBHE3$qtW"
	// pgsql_dbname = "postgres"
	// pgsql_host := "db.yyleioknuokgmzoqyyjc.supabase.co"
	pgsql_user = db.PgSqlUser
	pgsql_pass = db.PgSqlPass
	pgsql_dbname = db.PgSqlDbname
	pgsql_host := "db.yyleioknuokgmzoqyyjc.supabase.co"
	if db.GeneralDatabase == nil {
		fmt.Println(pgsql_user, pgsql_pass, pgsql_dbname)
		pgsql_dns = fmt.Sprintf("host=%s dbname=%s", pgsql_host, pgsql_dbname)
		dns := fmt.Sprintf("%s port=5432 user=%s password=%s sslmode=disable", pgsql_dns, pgsql_user, pgsql_pass)
		db.GeneralDatabase, errGeneralOpen = sql.Open("postgres", dns)
		fmt.Println(errGeneralOpen)
		if errGeneralOpen != nil {
			return append(err, errors.New(errGeneralOpen.Error()))
		}
		errGeneralPing = db.GeneralDatabase.Ping()
		if errGeneralPing != nil {
			return append(err, errors.New(errGeneralPing.Error()))
		}
		fmt.Printf("GeneralDatabase:%v", db.GeneralDatabase)
	}
	return nil
}

func (db *Database) ApiFaculty() repository.ApiFacultyInterface {
	return db.ApiFacultyInterface
}
