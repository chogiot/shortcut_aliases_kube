package main

import (
	
	"os"
	"fmt"
	"bufio"
	"strings"
	

)

func main() {

	fi, err := os.Open("aliases")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)
    for scanner.Scan() {
		// fmt.Println(scanner.Text())
		if strings.Contains(scanner.Text(), "delete") {
			fmt.Println(scanner.Text())
		}
    }

    if err := scanner.Err(); err != nil {
		panic(err)
		
		// if strings.Contains(lineOfFile, "delete") {
			// fmt.Println(lineOfFile)
		// }
	}
	// }
}
		