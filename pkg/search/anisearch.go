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
		name := doc.Find(".link > a").Text()
		seed, _ := doc.Find("span[title]").Attr("title")

		h.ForEach(".links > a", func(i int, e *colly.HTMLElement) {
			if e.Text == "Magnet" {
				magnet = e.Attr("href")
				return
			}
		})

		if strings.Contains(name, "[EMBER]") && len(seed) != 0 {

			seed = strings.TrimSpace(strings.ReplaceAll(seed[0:strings.Index(seed, "/")], "Seeders:", ""))

			entry := entry{
				Name:   name,
				Size:   doc.Find(".size").Text(),
				Magnet: magnet,
				Seeds:  seed,
			}

			entrys = append(entrys, entry)
		} else {
			skipCount++
		}
	})

	c.OnHTML(".home_list_pagination > a", func(h *colly.HTMLElement) {
		if pageCount < 15 {
			next_page := h.Request.AbsoluteURL(h.Attr("href"))
			pageCount++
			c.Visit(next_page)
		}
	})

	c.Visit(link)

	message := make(map[string]interface{})
	logger.Log.Info("Completed Scraping!")

	if len(entrys) != 0 {
		message["Results"] = entrys
		logger.Log.Infof("Skipped %v entrys that did not equal '[EMEBER]'!", skipCount)
	} else {

		message["Results"] = "No Results Found!"
		logger.Log.Info("No Entrys Found!")
	}
	return message
}
