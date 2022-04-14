package main

import (
	"fmt"
	"bufio"
	"os"
	c2d "github.com/acheong08/GoShell/data"
	mod "github.com/acheong08/GoShell/modules"
	"github.com/vova616/screenshot"
)


func shell(){
	var cmd c2d.Cmd
	var input string
	reader := bufio.NewReader(os.Stdin)
	cmd.Id = "1"
	cmd.Os = "linux"
	cmd.Target = "all"
	cmd.Kind = "shell"
	for cmd.Command != "exit\n"{
		fmt.Print("Enter command: ")
		cmd.Command, _ = reader.ReadString('\n')
		out, err := mod.ExecuteCMD(cmd)
		if err == nil{
			fmt.Printf("\nExec output: %s", out)
			fmt.Println(input)
		} else{
			fmt.Print("Error occured: ")
			fmt.Println(err)
		}
	}
}
func screenshot(){

}