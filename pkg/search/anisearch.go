package search

import (
	"animescrapper/pkg/logger"
	"strings"

	"github.com/gocolly/colly"
)

type entry struct {
	Name   string `json:"name"`
	Size   string `json:"size"`
	Seeds  string `json:"seed"`
	Magnet string `json:"magnet"`
}

func AniSearch(title string) map[string]interface{} {

	var entrys []entry

	c := colly.NewCollector()
	skipCount := 0
	pageCount := 0
	title = strings.ReplaceAll(title, " ", "+")

	link := "https://animetosho.org/search?q=" + title

	c.OnHTML(".home_list_entry", func(h *colly.HTMLElement) {

		doc := h.DOM

		var magnet string
		name := doc.Find(".link > a").Text()             //get name of torrent
		seed, _ := doc.Find("span[title]").Attr("title") // get seed string

		h.ForEach(".links > a", func(i int, e *colly.HTMLElement) { // loop through links and find Magnet link
			if e.Text == "Magnet" {
				magnet = e.Attr("href") //add Magnet link to magnet var
				return
			}
		})

		if strings.Contains(name, "[EMBER]") && len(seed) != 0 { // Filter out seed string that are empty, and only add [EMBER] files | Note: Remove filter for [Ember]?

			seed = strings.TrimSpace(strings.ReplaceAll(seed[0:strings.Index(seed, "/")], "Seeders:", "")) //Clean up Seed String

			entry := entry{ //Create entry for torrent
				Name:   name,
				Size:   doc.Find(".size").Text(),
				Magnet: magnet,
				Seeds:  seed,
			}

			entrys = append(entrys, entry) //add entry to slice
		} else {
			skipCount++ // add to skip count
		}
	})

	c.OnHTML(".home_list_pagination > a", func(h *colly.HTMLElement) { //loop through all pages, max is 15
		if pageCount < 15 {
			next_page := h.Request.AbsoluteURL(h.Attr("href"))
			pageCount++
			c.Visit(next_page)
		}
	})

	c.Visit(link) // get site html

	message := make(map[string]interface{})
	logger.Log.Info("Completed Scraping!")

	if len(entrys) != 0 { //send entrys if found
		message["Results"] = entrys
		logger.Log.Infof("Skipped %v entrys that did not equal '[EMEBER]'!", skipCount)
	} else { //if entrys are empty then search returned 0 entrys

		message["Results"] = "No Results Found!"
		logger.Log.Info("No Entrys Found!")
	}
	return message
}
