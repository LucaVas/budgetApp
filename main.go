package main

import (
	"fmt"
	"os"

	"github.com/LucaVas/budgetApp/internal/io/cli"
	"github.com/LucaVas/budgetApp/internal/io/file"
)

func main() {
	args := os.Args[1:]
	m, err := cli.GetArgs(args)
	if err != nil {
		panic(err)
	}

	for _, p := range m {
		content, err := file.ReadCsv(p)
		if err != nil {
			panic(err)
		}

		for _, row := range content {
			fmt.Printf("Row: %v\n", row)
		}
	}
}
