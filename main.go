package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	//fmt.Println("hello world")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		textToAddOntoLine := util.GetWhatIShouldAddOntoTheLine()

		fmt.Printf("%s%s\n", line, textToAddOntoLine)

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
