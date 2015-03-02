package main

import (
	"encoding/json"
	"fmt"
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
	v := 0
	for _, path := range os.Args[1:] {
		gold := make(map[string]int)
		parser := yasha.ParserFromFile(path)

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
		parser.OnEntityCreated = func(e *yasha.PacketEntity) {
			if e != nil {
				jsonify(e)
			}
			v++
			//if v > 1000000000 {
			//	os.Exit(0)
			//}
		}
		parser.Parse()
		//for hero, total := range gold {
		//	fmt.Println(hero, total)
		//}

		/*
		   	{
		       "Tick": 69914,
		       "Index": 37,
		       "SerialNum": 857,
		       "ClassId": 392,
		       "EntityHandle": 1755173,
		       "Name": "DT_DOTA_PlayerResource",
		       Values : {
		   	     "m_iTotalEarnedGold.0000": 14235,
		         "m_iTotalEarnedGold.0001": 14725,
		         "m_iTotalEarnedGold.0002": 11609,
		         "m_iTotalEarnedGold.0003": 22709,
		         "m_iTotalEarnedGold.0004": 14638,
		         "m_iTotalEarnedGold.0005": 12340,
		         "m_iTotalEarnedGold.0006": 8321,
		         "m_iTotalEarnedGold.0007": 11093,
		         "m_iTotalEarnedGold.0008": 10400,
		         "m_iTotalEarnedGold.0009": 6238,
		    }
		*/
		if false {
			for _, e := range parser.Entities {
				if e != nil && e.Name == "DT_DOTA_PlayerResource" {
					for _, p := range playersSuffix() {
						x := "m_iTotalEarnedGold." + p
						fmt.Println(p, ":", e.Values[x])
						fmt.Println(x)
					}
					//	for k, v := range e.Values {
					//		fmt.Println(k, v)
					//	}
					// fmt.Println(e.Name)
				}
			}
		}
	}
}

func jsonify(in interface{}) {
	data, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		panic(err)
	}
	spew.Println(string(data))
}

func playersSuffix() []string {
	players := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		players = append(players, fmt.Sprintf("%04d", i))
	}
	return players
}
