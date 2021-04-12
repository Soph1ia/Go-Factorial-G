package p

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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
	// Calculate Factorial
	response := strconv.FormatUint(factorial(number), 16)
	// print out the factorial output
	_, _ = fmt.Fprint(w, response)
}

/* Variable Declaration */
var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.

/*     function declaration        */
func factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	}else{
		for i:=1; i<=n; i++ {
			factVal *= uint64(i)  // mismatched types int64 and int
		}
		return factVal
	}
	return 0  /* return from function*/
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

