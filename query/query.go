package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

const CSDN_URL string  = "https://blog.csdn.net/cjqh_hao"

func main() {
	records := getArticles()
	f, err := os.Create("./query/title.txt")
	if err != nil {
		fmt.Println(f, err)
		return
	}

	num := len(records)
	for i := 0; i < num; i++ {
		l, err := f.WriteString(records[i])
		if err != nil {
			fmt.Println(l, err)
			return
		}
    }
	f.Close()
	return
}

// 抓取文章名称和链接
func getArticles() []string {
	doc, err := goquery.NewDocument(CSDN_URL)
    if err  != nil {
		fmt.Println(err)
        os.Exit(1)
	}

	var records []string
    doc.Find(".article-item-box").Each(func(i int, s *goquery.Selection) {
		text:= s.Find("a").First().Text()
		url, _:= s.Find("a").First().Attr("href")
		record := text + " " + url
		records = append(records, record)
	})
	return records
}