/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/Mahcks/nba-lottery/cmd"
)

var Version string

func main() {
	fmt.Println(Version)

	cmd.Execute()
}
