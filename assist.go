// // package main

// // import (
// // 	"fmt"
// // 	"log"
// // 	"os"

// // 	"github.com/urfave/cli/v2"
// // )

// // func main() {
// // 	app := &cli.App{
// // 		Name:  "boom",
// // 		Usage: "make an explosive entrance",
// // 		Action: func(args *cli.Context) error {
// // 			fmt.Println(args.Args().Get(0))
// // 			return nil
// // 		},
// // 	}

// // 	if err := app.Run(os.Args); err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"flag"
// 	"os"
// )

// func main() {
// 	testCmd := flag.NewFlagSet("test", flag.ExitOnError)
// 	testHell := testCmd.Bool("hello", false, "hello world")

// 	hiCmd := flag.NewFlagSet("hi", flag.ExitOnError)
// 	hiName := hiCmd.String("name", "", "Hello ${name}")
// }

package main

import (
	commandfuncs "assist/commandFuncs"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		os.Exit(1)
	}
	command := flag.Arg(0)
	runCommand(command, flag.Args()[1:])
}

func runCommand(command string, args []string) error {
	switch command {
	case "hello":
		fmt.Println("Hello", args[0])
	case "GPTepal":
		commandfuncs.ChatGpt()
	default:
		fmt.Println("Command not valid")
	}
	return nil
}
