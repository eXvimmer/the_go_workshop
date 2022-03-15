package main

import (
	"fmt"
	"time"
)

/* NOTE:
 * ----------------------------------------------------------------------------
 * |  Type                         | Zero Value
 * |-------------------------------|-------------------------------------------
 * |  bool                         | false
 * |-------------------------------|-------------------------------------------
 * |  numbers                      | 0
 * |-------------------------------|-------------------------------------------
 * |  strings                      | ""
 * |-------------------------------|-------------------------------------------
 * |  pointers, functions, maps,   | nil
 * |  interfaces, slices, channels |
 * ----------------------------------------------------------------------------
 */

/* NOTE:
* -----------------------------------------------------------------------------
* |  Substitution    | Formatting
* |------------------|---------------------------------------------------------
* |  %v              | Any value. You don't care about type you're using
* |------------------|---------------------------------------------------------
* |  %+v             | Values with extra information, such as struct filed names
* |------------------|---------------------------------------------------------
* |  %#v             | Go syntax, such as %+v with the addition of the name of
* |                  | the type of the variable
* |------------------|---------------------------------------------------------
* |  %T              | Print the variable's type
* |------------------|---------------------------------------------------------
* |  %d              | Decimal (base 10)
* |------------------|---------------------------------------------------------
* |  %s              | String
* -----------------------------------------------------------------------------
 */

func main() {
	var count int
	fmt.Printf("Count: %#v \n", count)

	var discount float64
	fmt.Printf("Discount: %#v \n", discount)

	var debug bool
	fmt.Printf("Debug: %#v \n", debug)
	fmt.Printf("Debug: %t \n", debug)

	var message string
	fmt.Printf("Message: %#v \n", message)

	var emails []string
	fmt.Printf("Emails: %#v \n", emails)

	var startTime time.Time
	fmt.Printf("Start time: %#v \n", startTime)

}
