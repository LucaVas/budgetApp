package main

import (
	"fmt"
	"os"

	"github.com/LucaVas/budgetApp/internal/io/cli"
)

func main() {
	args := os.Args[1:]
	m, err  := cli.GetArgs(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}