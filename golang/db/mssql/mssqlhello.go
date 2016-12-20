package main

import (
	"fmt"

	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	//constr := "server=gocpntsqlsplt01;user id=srvreport;password=srvreport;port=1433" // Working
	//constr := "server=gocpntsqlsplt01;port=1433" // Working // if no credentials are provided, current logfed-in windows users authentication
	//constr := "server=gocpntsqlsplt01;port=1433;database=PMIDBA" // Worked

	//constr := "server=172.19.205.226;port=1433"  			  // ERROR: [Scan] driver: bad connection
	//constr := "server=db9.erpint.pmi.org;port=1433" // ERROR: [Scan] sql: no rows in result set
	//constr := "server=ERPQADB9.pmienvs.pmihq.org;port=1433" // ERROR: [Scan] sql: no rows in result set

	//constr := "server=db7.erpqa.pmi.org;port=1433" // ERROR: sometimes "[Scan] sql: no rows in result set". sometimes "[Scan] driver: bad connection"
	//constr := "server=erpqadb7.pmienvs.pmihq.org;port=1433" // ERROR: sometimes "[Scan] sql: no rows in result set". sometimes "[Scan] driver: bad connection"
	//constr := "server=erpqadb7;port=1433" // ERROR: unknown host

	//constr := "server=localhost;port=1433" 				 	// ERROR: [Prepare] Login error: read tcp 127.0.0.1:58288->127.0.0.1:1433: wsarecv: An existing connection was forcibly closed by the remote host.
	//constr := "server=localhost;user id=ru;password=ru=1433" 	// ERROR: [Prepare] Login error: read tcp [::1]:58295->[::1]:1433: 		   wsarecv: An existing connection was forcibly closed by the remote host.
	//constr := "server=127.0.0.1;port=1433"					// ERROR: [Prepare] Login error: read tcp 127.0.0.1:58765->127.0.0.1:1433: wsarecv: An existing connection was forcibly closed by the remote host.
	//constr := `server=PC_7929P72;port=1433` // ERROR: [Prepare] Login error: read tcp [fe80::fc8d:c5c6:8ead:d8bc%Wireless Network Connection]:58016->[fe80::fc8d:c5c6:8ead:d8bc]:1433: wsarecv: An existing connection was forcibly closed by the remote host.
	constr := `server=PC_7929P72;port=1433;encrypt=disable`
	//constr := `server=PC_7929P72.pmint.pmihq.org\SQL2012;port=1433`// ERROR
	//constr := `server=PC_7929P72\SQL2012;port=1433` // ERROR
	//constr := `server=PC_B7FHL32;port=1433` // ERROR

	//constr := "server=SQL2012TEST01V;port=1433"

	db, err := sql.Open("mssql", constr)
	if err != nil {
		fmt.Println("[Open]", err.Error())
		return
	}
	defer db.Close()

	qry, err := db.Prepare("select top 1 name from sys.tables")
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
