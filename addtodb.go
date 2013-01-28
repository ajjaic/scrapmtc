package main

import sq "github.com/mattn/go-sqlite3"
import "database/sql/driver"
import "os"

var busdb string = "bus.db"
var bustableschema string = "CREATE TABLE buses(ROUTE TEXT, SERVT TEXT, ORIGIN TEXT, DEST TEXT, DURATION INTEGER);"


func addBusToDB() error {
  
  if _, e := os.Open(busdb); os.IsNotExist(e) {
    dbdrv := new(sq.SQLiteDriver)
    dbcon, _ := dbdrv.Open(busdb)
    commitTrans(dbcon, bustableschema)
  }

  dbdrv := new(sq.SQLiteDriver)
  dbcon, _ := dbdrv.Open(busdb)
  op1 := `INSERT INTO buses VALUES("01A", "ORDINARY", "ANNA", "PAADI", 100)`
  op2 := `INSERT INTO buses VALUES("ZM70", "AIR CONDITIONER", "GUINDY", "ADYAR", 65)`
  commitTrans(dbcon, op1)
  commitTrans(dbcon, op2)
  dbcon.Close()
  return nil 
}

func commitTrans(dbcon driver.Conn, cmd string) error {
  dbtrans, e := dbcon.Begin()
  if e != nil {
    return e
  }

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

  e = dbtrans.Commit()
  if e != nil {
    return e
  }
  

  return nil
}