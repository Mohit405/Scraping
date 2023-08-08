package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type ScrappedData struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	ImageURL string `json:"imageurl"`
}

func main() {
	var AllScrappedData []ScrappedData
	c := colly.NewCollector()

	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		scrapIT := ScrappedData{
			h.ChildText("h2"),
			h.ChildAttr("a", "href"),
			h.ChildAttr("img", "src"),
		}

		AllScrappedData = append(AllScrappedData, scrapIT)
	})

	c.OnHTML("[class=page-numbers]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://scrapeme.live/shop/")

	data, err := json.Marshal(AllScrappedData)
	if err != nil {
		log.Print("Failed to marshal the Scrapeddata")
		return
	}

	os.WriteFile("product.json", data, 0644)
}
