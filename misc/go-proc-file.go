package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Print line and character positions for easy slice ref
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

	filename := flag.String("filename", "01.txt", "Filename")
	outfilename := flag.String("outfilename", "01.out.txt", "Output Filename")
	flag.Parse()

	// Open csv file
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileout, err := os.Create(*outfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer fileout.Close()

	// Loop index
	i := 1

	// Init time conversion
	// Time parse
	const Formatin = "01/02/2006 15:04:05.00"

	// Time Format for MySQL TIMEDATE(2)
	const formatout = "2006-01-02 15:04:05.00"
	// tz stuff
	// loc, _ := time.LoadLocation("Europe/Bucharest")

	// Read lines from file
	scanner := bufio.NewScanner(file)
	w := bufio.NewWriter(fileout)

	//Loop
	for scanner.Scan() {

		// Print a few lines for easy slices
		if i < 3 {
			printLineNr(scanner.Text())
		}

		// Process line into variables
		datetime := scanner.Text()[0:22]
		freq := scanner.Text()[24:32]
		power := scanner.Text()[35:42]

		// tz stuff
		// t, _ := time.ParseInLocation(longForm, datetime, loc)

		// Parse time
		t, _ := time.Parse(Formatin, datetime)

		// Prepare/convert variables
		// fmt.Println("freq:", freq)
		freq = strings.TrimSpace(freq)
		fr, err := strconv.ParseFloat(freq, 4)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// if err == nil {
		// 	fmt.Printf("%T, %v\n", fr, fr)
		// }
		// fmt.Println("power:", power)
		power = strings.TrimSpace(power)
		pw, err := strconv.ParseFloat(power, 3)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// if err == nil {
		// 	fmt.Printf("%T, %v\n", pw, pw)
		// }

		// Insert data into file

		// wo := fmt.Sprintf("%v,%v,%v", t.Format(formatout), fr, pw)
		fmt.Fprintln(w, fmt.Sprintf("%v,%6.4f,%5.3f", t.Format(formatout), fr, pw))
		w.Flush()
		fmt.Println("inserted:", i, t.Format(formatout), fr, "Hz", pw, "MW")

		i++
		// insert i == n (+1) lines. i == 0 for all
		if i == 0 {
			return
		}
		// defer w.Flush()

	} // Loop end

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
