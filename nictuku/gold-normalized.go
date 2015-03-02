package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"strconv"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
	"github.com/davecgh/go-spew/spew"
	"github.com/dotabuff/yasha"
	"github.com/dotabuff/yasha/dota"
)

var pp = spew.Dump

var whitelist = map[uint64]bool{76561197985166697: true}

var steamHero = map[uint64]string{}
var suffixSteam = map[string]uint64{} // 0005

type playerGold struct {
	suffix      string
	minute      int
	gold        uint
	totalEarned plotter.XYs
	totalGained plotter.XYs
}

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
		players := []*playerGold{}
		for _, p := range playersSuffix() {
			players = append(players, &playerGold{suffix: p})
		}
		p.Title.Text = "Gold for" + path
		p.X.Label.Text = "Minute"
		p.Y.Label.Text = "Gold"

		parser.OnFileInfo = func(fileinfo *dota.CDemoFileInfo) {
			for _, pls := range fileinfo.GetGameInfo().GetDota().GetPlayerInfo() {
				steamHero[pls.GetSteamid()] = pls.GetHeroName()
				//				jsonify(pls)
			}
		}

		if true {

			parser.OnEntityPreserved = func(e *yasha.PacketEntity) {

				if e != nil {
					if e.Name == "DT_DOTA_PlayerResource" {
						//for _, p := range playersSuffix() {
						//	x := "m_iTotalEarnedGold." + p
						//	fmt.Println(p, ":", e.Values[x])
						//	fmt.Println(x)
						//}
						//jsonify(e)
						// 1800 ticks per minute, 30 per second.

						//fmt.Println(e.Tick)
						//fmt.Println("gold", e.Values["m_iTotalEarnedGold."+playersSuffix()[3]])
						for _, p := range players {

							g, ok := e.Values["m_iTotalEarnedGold."+p.suffix].(uint)
							if !ok {
								fmt.Println("not uint")
								return
							}
							suffixSteam[p.suffix] = e.Values["m_iPlayerSteamIDs."+p.suffix].(uint64)

							if e.Tick/1800 != p.minute {
								//t := float64(e.Tick - tick)
								//fmt.Println("t", t)
								delta := g - p.gold // removes the background ones
								if delta == 0 {
									return
								}

								//	fmt.Println("delta", delta)
								//	fmt.Println("earned", delta-100)
								p.totalEarned = append(p.totalEarned, struct{ X, Y float64 }{float64(e.Tick / 1800), float64(delta) - 100})
								p.totalGained = append(p.totalEarned, struct{ X, Y float64 }{float64(e.Tick / 1800), float64(delta)})

								if err != nil {
									fmt.Printf("BOOM", err)
								}
								p.gold = g
								p.minute = e.Tick / 1800
							}
						}
					}

					if e.Name == "DT_DOTA_DataRadiant" { // || e.Name == "DT_DOTA_DataDire"
						//fmt.Println(e.Tick)

						//fmt.Println("relgold", e.Values["m_iReliableGold.0003"])
						//fmt.Println("unrelgold", e.Values["m_iUnreliableGold.0003"])
					}
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
		for i, pl := range players {
			if _, ok := whitelist[suffixSteam[pl.suffix]]; !ok {
				continue
			}
			l, s, err := plotter.NewLinePoints(pl.totalEarned)
			if err != nil {
				panic(err)
			}
			l.Color = plotutil.Color(i)
			p.Add(l, s)
			p.Legend.Add(steamHero[suffixSteam[pl.suffix]], l)
			// err = plotutil.AddLinePoints(p, pl.suffix, pl.totalEarned)
		}
		// Save the plot to a PNG file.
		if err := p.Save(10, 5, fmt.Sprintf("gold-%v.png", path)); err != nil {
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
