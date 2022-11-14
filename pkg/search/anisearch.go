package search

import (
	"animescrapper/pkg/logger"
	"os"
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
}

func AniSearch(query Query) any {

	var entrys []entry

	c := colly.NewCollector()
	skipCount := 0
	pageCount := 0
	title := strings.ReplaceAll(query.Title, " ", "+")

	logger.Log.Info(title)

	link := "https://animetosho.org/search?q=" + title

	c.OnHTML(".home_list_entry", func(h *colly.HTMLElement) {

		doc := h.DOM

		var magnet string
		//get name of torrent
		name := doc.Find(".link > a").Text()
		// get seed string
		seed, _ := doc.Find("span[title]").Attr("title")

		// loop through links and find Magnet link
		h.ForEach(".links > a", func(i int, e *colly.HTMLElement) {
			if e.Text == "Magnet" {
				//add Magnet link to magnet var
				magnet = e.Attr("href")
				return
			}
		})

		if len(seed) == 0 {
			skipCount++
			return
		} else {
			//Clean up Seed String
			seed = strings.TrimSpace(
				strings.ReplaceAll(seed[0:strings.Index(seed, "/")],
					"Seeders:", ""))
		}

		if len(query.Filter) > 0 {
			for _, filter := range query.Filter {
				if !strings.Contains(name, filter) {
					skipCount++
					return
				}
			}
		}
		//Check if seedMin was placed
		if query.MinSeed != 0 {

			seedValue, _ := strconv.Atoi(seed)
			//Check if seedValue is less then the seedMin
			if seedValue < query.MinSeed {
				skipCount++
				return
			}
		}

		//Create entry for torrent
		entry := entry{
			Name:   name,
			Size:   doc.Find(".size").Text(),
			Magnet: magnet,
			Seeds:  seed,
		}
		//add entry to slice
		entrys = append(entrys, entry)

	})
	//loop through all pages, max is 15
	c.OnHTML(".home_list_pagination > a", func(h *colly.HTMLElement) {
		// set a default for MaxPage
		maxPage, err := strconv.Atoi(os.Getenv("MAX_PAGE"))

		if err != nil {
			logger.Log.Fatal(err)
		}

		if pageCount < maxPage {
			next_page := h.Request.AbsoluteURL(h.Attr("href"))
			pageCount++
			c.Visit(next_page)
		}
	})
	// get site html
	c.Visit(link)

	// set a default for MaxEntry
	if query.MaxEntry == 0 {
		query.MaxEntry = len(entrys)
	}

	// slice up the query.MaxEntry
	if len(entrys) > query.MaxEntry {
		entrys = entrys[:query.MaxEntry]
	}

	logger.Log.Info("Completed Scraping!")

	//send entrys if found, else return 0
	if len(entrys) != 0 {
		logger.Log.Infof(
			"Skipped %v entrys that had no seed or was filtered out!",
			skipCount)

		return entrys
	} else {
		logger.Log.Info("No Entrys Found!")
		return entrys
	}

}
