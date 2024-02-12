package main

import (
	"launchingprogramlist/config"
	"launchingprogramlist/runer"
	"log"
	"sort"
)

func main() {
	configs, err := config.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(configs, func(i, j int) bool {
		a := configs[i]
		b := configs[j]
		return a.SortIndex < b.SortIndex
	})
	for _, configApp := range configs {
		err := runer.RunApp(configApp)
		if err != nil {
			log.Fatal(err)
		}

	}
}
