package p

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			_, _ = fmt.Fprint(w, "Enter a number")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	if d.Message == "" {
		_, _ = fmt.Fprint(w, "Please input a number for factorial")
		return
	}

	//Convert string to number.
	number, _ := strconv.Atoi(d.Message)
	// Call Benchmarking Function
	benchmark("factorial", number, w)
}

/**
Method : Benchmark

This method gets the time taken to execute the factorial 40 times.
In total it loops 80 times.
It takes the last 20 execution times.
Gets the average time
Calculates the throughput as40/time 

Prints out the throughput.

returns: none

*/
func benchmark(funcName string, number int, w http.ResponseWriter) {
	listofTime := [41]int64{}

	for j := 0; j < 40; j++ {
		start := time.Now().UnixNano()
		factorial(number)

		// End time
		end := time.Now().UnixNano()
		// Results
		difference := end - start
		listofTime[j] = difference

	}
	// Average Time
	sum := int64(0)
	for i := 0; i < len(listofTime); i++ {
		// adding the values of
		// array to the variable sum
		sum += listofTime[i]
	}
	// avg to find the average
	avg := (float64(sum)) / (float64(len(listofTime)))

	// Throughput Rate
	throughput := 40/avg

	// Response
	fmt.Fprintf(w, "Time taken by %s function is %v ops/ns \n", funcName, throughput)
}

/**
Method: Factorial

Calculates the factorial of the number provided

Returns: pointer to big int
*/
func factorial(n int) *big.Int {
	factVal := big.NewInt(1)
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			//factVal *= uint64(i) // mismatched types int64 and int
			factVal = factVal.Mul(factVal, big.NewInt(int64(i)))
		}
	}
	return factVal
}

/**
Testing Purposes
*/
//func main() {
//
//	//Convert string to number.
//	number, _ := strconv.Atoi("4")
//	// Calculate Factorial
//	response := strconv.FormatUint(factorial(number), 10)
//	// print out the factorial output
//	_, _ = fmt.Print(response)
//}
