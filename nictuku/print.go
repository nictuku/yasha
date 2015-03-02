package main

import (
	//"encoding/json"

	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/yasha"
	//"github.com/dotabuff/yasha/dota"
)

var pp = spew.Dump

func main() {
	if len(os.Args) < 2 {
		spew.Println("Expected a .dem file as argument")
	}

	for _, path := range os.Args[1:] {
		parser := yasha.ParserFromFile(path)
		parser.Parse()

		parser.PrintDistinctCombatLogTypes()
		//parser.OnCombatLog = func(combat yasha.CombatLogEntry) {
		//		fmt.Println(combat.Type())
		//	}
	}
}
