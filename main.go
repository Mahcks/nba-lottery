/*
Copyright © 2022 Max max@mahcks.com

*/
package main

import (
	"github.com/Mahcks/nba-lottery/cmd"
	"github.com/Mahcks/nba-lottery/internal/structures"
)

var Version string

func main() {
	if len(Version) > 0 {
		structures.SetVersion(Version)
	}

	cmd.Execute()
}
