package command

import (
  f "fmt"
  "os"
  "github.com/pborman/getopt/v2"
)

func Run(args []string) {
  f.Println("Command")
  optName := getopt.StringLong("name", 'n', "", "Your name")
  optHelp := getopt.BoolLong("help", 0, "Help")
  getopt.Parse()

  if *optHelp {
    getopt.Usage()
    os.Exit(0)
  }
  /*if len(args) < 1 {
    f.Println("Mising flags")

  } else {
    f.Println(args)
  }*/
  f.Println("Hello " + *optName + "!")
}
