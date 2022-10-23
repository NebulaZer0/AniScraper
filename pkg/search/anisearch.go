package search

import (
	"animescrapper/pkg/logger"
	"strconv"

	"strings"

	"github.com/gocolly/colly"
)

type entry struct {
	Name   string `json:"name"`
	Size   string `json:"size"`
	Seeds  string `json:"seed"`
	Magnet string `json:"magnet"`
}

type Query struct {
	Title    string   `json:"title"`
	Filter   []string `json:"filter"`
	MinSeed  int      `json:"minSeed"`
	MaxEntry int      `json:"maxEntry"`
	MaxPage  int      `json:"maxPage"`
}

func AniSearch(query Query) map[string]interface{} {

	var entrys []entry

	c := colly.NewCollector()
	skipCount := 0
	pageCount := 0
	title := strings.ReplaceAll(query.Title, " ", "+")

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

		if len(seed) == 0 {
			skipCount++
			return
		} else {
			seed = strings.TrimSpace(strings.ReplaceAll(seed[0:strings.Index(seed, "/")], "Seeders:", "")) //Clean up Seed String
		}

		if len(query.Filter) > 0 { //Check if filter were placed
			for _, filter := range query.Filter {
				if !strings.Contains(name, filter) {
					skipCount++
					return
				}
			}
		}

		if query.MinSeed != 0 { //Check if seedMin was placed

			seedValue, _ := strconv.Atoi(seed)

			if seedValue < query.MinSeed { //Check if seedValue is less then the seedMin
				skipCount++
				return
			}
		}

		entry := entry{ //Create entry for torrent
			Name:   name,
			Size:   doc.Find(".size").Text(),
			Magnet: magnet,
			Seeds:  seed,
		}

		entrys = append(entrys, entry) //add entry to slice

	})

	// set a default for MaxPage
	if query.MaxPage == 0 {
		query.MaxPage = 15
	}

	c.OnHTML(".home_list_pagination > a", func(h *colly.HTMLElement) { //loop through all pages, max is 15
		if pageCount < query.MaxPage {
			next_page := h.Request.AbsoluteURL(h.Attr("href"))
			pageCount++
			c.Visit(next_page)
		}
	})

	c.Visit(link) // get site html

	// set a default for MaxEntry
	if query.MaxEntry == 0 {
		query.MaxEntry = len(entrys)
	}

	// slice up the query.MaxEntry
	if len(entrys) > query.MaxEntry {
		entrys = entrys[:query.MaxEntry]
	}
	message := make(map[string]interface{})
	logger.Log.Info("Completed Scraping!")

	if len(entrys) != 0 { //send entrys if found
		message["Results"] = entrys
		logger.Log.Infof("Skipped %v entrys that had no seed or was filtered out!", skipCount)
	} else { //if entrys are empty then search returned 0 entrys

		message["Results"] = "No Results Found!"
		logger.Log.Info("No Entrys Found!")
	}
	return message
}
