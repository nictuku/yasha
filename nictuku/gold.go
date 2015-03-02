package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/yasha"
	"github.com/dotabuff/yasha/dota"
)

var pp = spew.Dump

func main() {
	if len(os.Args) < 2 {
		spew.Println("Expected a .dem file as argument")
	}

	for _, path := range os.Args[1:] {
		gold := make(map[string]int)
		parser := yasha.ParserFromFile(path)

		parser.OnFileInfo = func(fileinfo *dota.CDemoFileInfo) {
			data, err := json.MarshalIndent(fileinfo, "", "  ")
			if err != nil {
				panic(err)
			}
			spew.Println(string(data))
		}

		if false {
			parser.OnCombatLog = func(combat yasha.CombatLogEntry) {
				//fmt.Println(combat.Type())
				switch t := combat.(type) {
				case *yasha.CombatLogGold:
					gold[t.Target] += t.Value
					//fmt.Println(t.Target, t.Value)
				}
			}
		}
		parser.Parse()
		for hero, total := range gold {
			fmt.Println(hero, total)
		}
	}
}
