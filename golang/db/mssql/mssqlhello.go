package main

import (
	"fmt"

	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	//constr := "server=gocpntsqlsplt01;user id=srvreport;password=srvreport;port=1433" // Working

	constr := "server=gocpntsqlsplt01;port=1433" // Working // if no credentials are provided, current logfed-in windows users authentication

	//constr := "server=db7.erpqa.pmi.org;port=1433" // ERROR: sometimes "[Scan] sql: no rows in result set". sometimes "[Scan] driver: bad connection"
	//constr := "server=erpqadb7.pmienvs.pmihq.org;port=1433" // ERROR: sometimes "[Scan] sql: no rows in result set". sometimes "[Scan] driver: bad connection"
	//constr := "server=erpqadb7;port=1433" // ERROR: unknown host

	//constr := "server=localhost;port=1433" // if no credentials are provided, current logfed-in windows users authentication
	//constr := "server=localhost;user id=ru;password=ru=1433"

	db, err := sql.Open("mssql", constr)
	if err != nil {
		fmt.Println("[Open]", err.Error())
		return
	}
	defer db.Close()

	qry, err := db.Prepare("select FriendlyName from PMIDBA.dbo.ServerInventory")
	if err != nil {
		fmt.Println("[Prepare]", err.Error())
		return
	}
	row := qry.QueryRow()

	name := ""

	err = row.Scan(&name)
	if err != nil {
		fmt.Println("[Scan]", err.Error())
		return
	}

	fmt.Println(name)

}
