package main

import (
	_ "bufio"
	"fmt"
	_ "os"
	"strings"
	_ "strings"

	"github.com/fatih/color"
	"github.com/lunux2008/xulu"
)
func scancode(code []string, line int) string {
  var output string
  var serror = color.New(color.FgRed, color.Bold) 
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
  } else {
    output = ""
    serror.Printf(fmt.Sprint("netsync: syntax-error[not-defined]: the function is not defined - line ", line, "\n"))
  }
  return output
}