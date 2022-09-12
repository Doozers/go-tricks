package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Doozers/go-tricks/compileTimeVar/pkg/version"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	root := &ffcli.Command{
		ShortUsage: "testCmd ",
		Exec: func(_ context.Context, args []string) error {
			fmt.Println("version is:", version.Version)
			return nil
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
