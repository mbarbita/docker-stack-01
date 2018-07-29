package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func printLineNr(st string) {
	fmt.Println("Line:", st)
	fmt.Println("Line Len:", len(st))
	for j := 0; j < len(st); j++ {
		fmt.Printf("%2v|", st[j:j+1])
	}
	fmt.Print("\n")
	for j := 0; j < len(st); j++ {
		fmt.Printf("%2v|", j)
	}
	fmt.Print("\n\n")
}

func main() {

	user := flag.String("user", "", "Username")
	password := flag.String("password", "", "Password")
	address := flag.String("address", "", "IP address")
	database := flag.String("database", "grafana", "Database")
	table := flag.String("table", "", "Table")
	filename := flag.String("filename", "", "Filename")
	// metric := flag.String("metric", "generic", "Metric")
	flag.Parse()

	// var val1 int
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
	stm := "INSERT INTO " + *table + "(time, val1, metric1, val2, metric2) VALUES(?,?,?,?,?)"
	insForm, err := db.Prepare(stm)
	if err != nil {
		panic(err.Error())
	}

	// rand.Seed(42) // Try changing this number!

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	i := 1
	const longForm = "01/02/2006 15:04:05.00"
	const format = "2006-01-02 15:04:05.00"
	// loc, _ := time.LoadLocation("Europe/Bucharest")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i < 3 {
			printLineNr(scanner.Text())
		}
		//02/01/2018 20:14:59.55,  50.0148,   90.933 11998

		// 2014-09-08 17:51:04.777
		datetime := scanner.Text()[0:22]
		// datetime := "02/01/2018 20:14:59.55"
		freq := scanner.Text()[24:32]
		power := scanner.Text()[35:42]
		// dmm := scanner.Text()[0:2]
		// ddd := scanner.Text()[3:5]
		// dyyyy := scanner.Text()[6:10]
		// thh := scanner.Text()[11:13]
		// tmm := scanner.Text()[14:16]
		// tss := scanner.Text()[17:19]
		// tms := scanner.Text()[20:22]
		// hzi := scanner.Text()[25:27]
		// hzr := scanner.Text()[28:32]
		// pwi := scanner.Text()[36:38]
		// pwr := scanner.Text()[39:42]
		// fmt.Println(scanner.Text(), i, ddd, dmm, dyyyy, thh, tmm, tss, tms, ":::", hzi, hzr, pwi, pwr)

		// t, _ := time.ParseInLocation(longForm, datetime, loc)
		t, _ := time.Parse(longForm, datetime)

		// fmt.Println("Parse:", t)
		// fmt.Println("Unix format:", t.Format(time.UnixDate))
		// fmt.Println("ANSIC format:", t.Format(time.ANSIC))
		// fmt.Println("RFC.. format:", t.Format(time.RFC3339Nano))
		// fmt.Println("mysql format:", t.Format("2006-01-02 15:04:05.00"))
		// fmt.Println("nano:", t.Nanosecond())

		// fmt.Println("freq:", freq)
		freq = strings.TrimSpace(freq)
		fr, _ := strconv.ParseFloat(freq, 4)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// if err == nil {
		// 	fmt.Printf("%T, %v\n", fr, fr)
		// }
		// fmt.Println("power:", power)
		power = strings.TrimSpace(power)
		pw, _ := strconv.ParseFloat(power, 3)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// if err == nil {
		// 	fmt.Printf("%T, %v\n", pw, pw)
		// }
		// stm := "INSERT INTO " + *table + "(time, val1, metric1,val2,metric2) VALUES(?,?,?,?,?)"
		insForm.Exec(t.Format(format), fr, "Hz", pw, "MW")
		fmt.Println("inserted:", i, t.Format(format), fr, "Hz", pw, "MW")

		i++
		// insert i == n (+1) lines. i == 0 for all
		if i == 0 {
			return
		}
		// return
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// val1 = rand.Intn(100)
	// log.Println("INSERT: time: ", time.Now(), " | val1: ", val1, " | metric: ", *metric)
	// time.Sleep(10 * 1000 * time.Millisecond)
}
