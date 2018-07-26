package main

import (
	"database/sql"
	"flag"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	user := flag.String("user", "", "Username")
	password := flag.String("password", "", "Password")
	address := flag.String("address", "", "IP address")
	database := flag.String("database", "grafana", "Database")
	table := flag.String("table", "", "Table")
	metric := flag.String("metric", "generic", "Metric")
	flag.Parse()

	var val1 int
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	db, err := sql.Open("mysql", *user+":"+*password+"@tcp("+*address+")/"+*database)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	stm := "INSERT INTO " + *table + "(time, val1, metric1) VALUES(NOW(),?,?)"
	insForm, err := db.Prepare(stm)
	if err != nil {
		panic(err.Error())
	}

	rand.Seed(42) // Try changing this number!
	for {
		val1 = rand.Intn(100)
		insForm.Exec(val1, metric)
		log.Println("INSERT: time: ", time.Now(), " | val1: ", val1, " | metric: ", metric)
		time.Sleep(10 * 1000 * time.Millisecond)
	}
}
