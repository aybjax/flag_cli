package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// create new app
	app := cli.NewApp()

	app.Version = "1.0"

	//add flags
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "name",
			Value: "default",
			Usage: "help documentation",
		},
		cli.IntFlag {
			Name: "num",
			Value: 0,
			Usage: "num",
		},
	}

	// app.Action = func(c *cli.Context) error {
	// 	fmt.Printf("Hello %s, number is %d\n", c.String("name"), c.Int("num"))

	// 	return nil
	// }
	app.Action = func(c *cli.Context) error {
		var args []string

		if c.NArg() > 0 {
			args = c.Args()
			personName := args[0]
			marks := args[1:]
			fmt.Println("Person: ", personName)
			fmt.Println("marks", marks)
		}

		if c.String("name") != "aybjax" {
			println("name is not aybjax")
		} else {
			fmt.Println("name is aybjax")
		}

		if c.Int("num") != 0 {
			println("num is not 0")
		} else {
			fmt.Println("num is 0")
		}

		return nil
	}

	app.Run(os.Args)
}