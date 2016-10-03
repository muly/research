package main

import _ "github.com/denisenkom/go-mssqldb"
import "database/sql"
import "log"
import "fmt"
import "flag"

var debug = flag.Bool("debug", true, "enable debugging")
var password = flag.String("password", "srvreport", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "gocpntsqlsplt01", "the database server")
var user = flag.String("user", "srvreport", "the database user")

func main() {
	flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select top 1 Environment, FriendlyName from PMIDBA.dbo.ServerInventory")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}

	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber string
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Println("col1 :", somenumber)
	fmt.Println("col2 :", somechars)

	fmt.Printf("bye\n")

}
