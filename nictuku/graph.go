package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/yasha"
	//"github.com/dotabuff/yasha/dota"
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
)

var pp = spew.Dump

func main() {
	if len(os.Args) < 2 {
		spew.Println("Expected a .dem file as argument")
	}

	for _, path := range os.Args[1:] {
		//gold := make(map[string]int)
		parser := yasha.ParserFromFile(path)

		p, err := plot.New()
		if err != nil {
			panic(err)
		}

		p.Title.Text = "Gold for" + path
		p.X.Label.Text = "Tick"
		p.Y.Label.Text = "Gold"
		a := plotter.XYs{}
		tick := 0
		parser.OnEntityPreserved = func(e *yasha.PacketEntity) {

			if e != nil {
				if e.Name == "DT_DOTA_PlayerResource" {
					//for _, p := range playersSuffix() {
					//	x := "m_iTotalEarnedGold." + p
					//	fmt.Println(p, ":", e.Values[x])
					//	fmt.Println(x)
					//}
					//jsonify(e)
					fmt.Println(e.Tick)
					fmt.Println("gold", e.Values["m_iTotalEarnedGold."+playersSuffix()[3]])
					g, ok := e.Values["m_iTotalEarnedGold."+playersSuffix()[3]].(uint)
					if !ok {
						fmt.Println("not uint")
						return
					}

					a = append(a, struct{ X, Y float64 }{float64(e.Tick), float64(g)})

					if err != nil {
						fmt.Printf("BOOM", err)
					}
					return
				}

				if e.Name == "DT_DOTA_DataRadiant" { // || e.Name == "DT_DOTA_DataDire"
					fmt.Println(e.Tick, "delta", e.Tick-tick)
					tick = e.Tick

					fmt.Println("relgold", e.Values["m_iReliableGold.0003"])
					fmt.Println("unrelgold", e.Values["m_iUnreliableGold.0003"])
				}
			}
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
		err = plotutil.AddLinePoints(p, playersSuffix()[3], a)
		// Save the plot to a PNG file.
		if err := p.Save(4, 5, fmt.Sprintf("gold-%v.png", path)); err != nil {
			panic(err)
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
