package main

import (
	"fmt"
)

/*
slice
*/

/*

 */
const (
	parseObjectKey   = iota // parsing object key (before colon)
	parseObjectValue        // parsing object value (after colon)
	parseArrayValue         // parsing array value
)

const (
	// Continue.
	scanContinue     = iota // uninteresting byte
	scanBeginLiteral        // end implied by next result != scanContinue
	scanBeginObject         // begin object
	scanObjectKey           // just finished object key (string)
	scanObjectValue         // just finished non-last object value
	scanEndObject           // end object (implies scanObjectValue if possible)
	scanBeginArray          // begin array
	scanArrayValue          // just finished array value
	scanEndArray            // end array (implies scanArrayValue if possible)
	scanSkipSpace           // space byte; can skip; known to be last "continue" result

	// Stop.
	scanEnd   // top-level value ended *before* this byte; known to be first "stop" result
	scanError // hit an error, scanner.err.
)

func main() {
	fmt.Println(parseObjectKey, parseObjectValue, parseArrayValue)
	fmt.Println("scan state\n",
		"scanContinue: ", scanContinue, "\n",
		"scanBeginLiteral: ", scanBeginLiteral, "\n",
		"scanBeginObject: ", scanBeginObject, "\n",
		"scanObjectKey: ", scanObjectKey, "\n",
		"scanObjectValue: ", scanObjectValue, "\n",
		"scanEndObject: ", scanEndObject, "\n",
		"scanBeginArray: ", scanBeginArray, "\n",
		"scanArrayValue: ", scanArrayValue, "\n",
		"scanEndArray: ", scanEndArray, "\n",
		"scanSkipSpace: ", scanSkipSpace,
	)
}
