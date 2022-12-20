package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	_ "strings"
  _ "time"
	"github.com/fatih/color"
	"github.com/lunux2008/xulu"
)
func scancode(code []string, line int) string {
  var output string
  var serror = color.New(color.FgRed, color.Bold) 
  var swarning = color.New(color.FgYellow, color.Bold)
  xulu.Use(output)
  if code[0] == "print"{  
    if len(code) > 1 && code[1] != "" {
      output = code[1]
    } else {
      output = ""
      serror.Printf(fmt.Sprint("netsync: syntax-error[argument-not-specified]: print expected 1 arg but got 0 - line ", line, "\n"))
    }
  } else if strings.HasPrefix(code[0], "//") {
    output = ""
  } else if code[0] == "pause"{
    fmt.Println("Press Enter to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n') 
  } else if code[0] == "" {} else if code[0] == "throw"{
    if len(code) > 2 && code[1] != "" && code[2] != "" {
      if code[1] == "error" {
        serror.Printf(fmt.Sprint(code[2], "  - line ", line, "\n"))
      } else if code[1] == "warning" {
        swarning.Printf(fmt.Sprint(code[2], " - line ", line, "\n"))
      }
    } else {
      output = ""
      serror.Printf(fmt.Sprint("netsync: syntax-error[argument-not-specified]: throw expected 2 args but got less - line ", line, "\n"))
    }
  } else {
    output = ""
    serror.Printf(fmt.Sprint("netsync: syntax-error[not-defined]: the function is not defined - line ", line, "\n"))
  }
  return output
}