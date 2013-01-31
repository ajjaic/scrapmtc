package main

import (
	"errors"
	"fmt"
	"os"
)

const buswebsite = `http://www.mtcbus.org/Routes.asp`
const prebusurl = `http://www.mtcbus.org/Routes.asp?cboRouteCode=`
const postbusurl = `&submit=Search`
const busdb = "bus.db"
const bustable = "buss"
const buscolmns = "route TEXT, servtype TEXT, origin TEXT, dest TEXT, jtime INTEGER"

var (
	ErrBusListAccess = errors.New("Cannot download the buslist ")
	ErrCreateDbase   = errors.New("Cannot open the database ")
	ErrTableCreate   = errors.New("Cannot create table ")
	ErrAddBuss       = errors.New("Unable to add bus data to database ")
	ErrUnableClose   = errors.New("Unable to close the connection to the database ")
	ErrGetBus        = errors.New("Unable to get bus data from remote server ")
	ErrAddBusPath    = errors.New("Unable to add bus path data to database ")
)

func main() {
	bl, e := getBusListfrmMTC(buswebsite)
	if e != nil {
		fmt.Println(ErrBusListAccess)
		os.Exit(-1)
	}
	dbc, e := createDB(busdb)
	if e != nil {
		fmt.Println(ErrCreateDbase, busdb)
		os.Exit(-1)
	}
	e = createTable(dbc, bustable, buscolmns)
	if e != nil {
		fmt.Println(ErrTableCreate, bustable)
		os.Exit(-1)
	}

	for _, v := range bl[:5] {
		b, e := newBus(prebusurl, v, postbusurl)
		if e != nil {
			fmt.Println(ErrGetBus, v)
			continue
		}
		e = addBus(b, bustable, dbc)
		if e != nil {
			fmt.Println(ErrAddBuss, v)
			continue
		}
		e = addBusPath(b, dbc)
		if e != nil {
			fmt.Print(ErrAddBusPath, v)
			continue
		}
		fmt.Println("added: ", v)

	}

	e = dbc.Close()
	if e != nil {
		fmt.Println(ErrUnableClose)
	}
}
