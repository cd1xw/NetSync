package main

import ( 
  "fmt"
  "os"
  "bufio"
  "strings"
  "github.com/fatih/color"
  _ "github.com/lunux2008/xulu"
)
func main() {
  if (len(os.Args) == 1){
    color.New(color.FgRed, color.Bold).Printf(fmt.Sprint("\033[31m", "netsync: program-error[cmd-not-specified]: for more commands run `netsync help` \n"))
    return
  }
  if (os.Args[1] == "open" || len(strings.Split(os.Args[1], ".")) == 2){
    if (len(os.Args) == 2 && len(strings.Split(os.Args[1], ".")) == 1){
      color.New(color.FgRed, color.Bold).Printf(fmt.Sprint("\033[31m", "netsync: program-error[file-not-specified]: specify a script to run \n"))
      return
    }
    var thingtoopen string 
    if len(strings.Split(os.Args[1], ".")) == 2 {
      thingtoopen = os.Args[1]
    } else { 
      thingtoopen = os.Args[2]
    }
   file, err := os.Open(thingtoopen) 
	 if err != nil {
		fmt.Println(err)
		return
	 }
	 defer file.Close()

	 scanner := bufio.NewScanner(file)
   var thing int
	 for scanner.Scan(){
     thing = thing + 1
     if (scancode(strings.Split(scanner.Text(), ", "), thing) != ""){
       fmt.Println(scancode(strings.Split(scanner.Text(), ", "), thing))
     }
   }

	 if err := scanner.Err(); err != nil {
		fmt.Println(err)
	 }
 } else {
   color.New(color.FgRed, color.Bold).Printf(fmt.Sprint("\033[31m", "netsync: program-error[invalid-cmd]: for more commands run `netsync help` \n"))
   return
 }
}
