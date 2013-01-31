package main

import (
	"database/sql/driver"
	sq "github.com/mattn/go-sqlite3"
	"strconv"
)

func createDB(dbname string) (driver.Conn, error) {
	d, e := new(sq.SQLiteDriver).Open(dbname)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func addBus(b *busdet, table string, dbcon driver.Conn) error {
	jmin := strconv.Itoa(b.jmin)
	insertstr := "INSERT INTO " + table + " VALUES(" +
		strconv.Quote(b.routenum) + "," +
		strconv.Quote(b.servtype) + "," +
		strconv.Quote(b.origin) + "," +
		strconv.Quote(b.dest) + "," + jmin + ");"
	e := commitTrans(dbcon, insertstr)
	if e != nil {
		return e
	}
	return nil
}

func addBusPath(b *busdet, c driver.Conn) error {
	colmns := "id INTEGER, stagename TEXT"
	tablename := "p" + b.routenum
	e := createTable(c, tablename, colmns)
	if e != nil {
		return e
	}
	for i, p := range b.stname {
		pq := "INSERT INTO " + tablename + " VALUES(" + strconv.Itoa(i+1) + "," + strconv.Quote(p) + ");"
		e = commitTrans(c, pq)
		if e != nil {
			return e
		}
	}
	return nil
}
func createTable(c driver.Conn, t, colmns string) error {
	tq := "CREATE TABLE IF NOT EXISTS " + t + "(" + colmns + ");"
	e := commitTrans(c, tq)
	if e != nil {
		return e
	}

	return nil
}

func commitTrans(dbcon driver.Conn, cmd string) error {
	dbstmt, e := dbcon.Prepare(cmd)
	if e != nil {
		return e
	}
	_, e = dbstmt.Exec(nil)
	if e != nil {
		return e
	}
	e = dbstmt.Close()
	if e != nil {
		return e
	}
	return nil
}
