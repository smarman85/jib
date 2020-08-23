package main

import (
  "os"
  "github.com/smarman85/jib/command"
)

func main() {
  command.Run(os.Args[1:])
}
