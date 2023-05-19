package main

import (
	"bufio"
	"fmt"
	"input_method/loader"
	"os"
)

func main() {
	rootNode := loader.StartLoader()
	scanner := bufio.NewScanner((os.Stdin))

	for {
		//get user input
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		//search in the trie tree
		res := rootNode.Search(input)
		for i, r := range res {
			fmt.Print(string(r), " ")
			if i == 10 {
				break
			}
		}
		fmt.Print("\n")
	}

}
