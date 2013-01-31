package main

import "github.com/PuerkitoBio/goquery"

func getBusListfrmMTC(bussite string) ([]string, error) {
	var b *goquery.Document
	var bb *goquery.Selection
	var fn func(int, *goquery.Selection) string
	b, e := goquery.NewDocument(bussite)
	if e != nil {
		return nil, e
	}
	bb = b.Find("Option")
	bb = bb.Not(":empty")
	fn = func(i int, s *goquery.Selection) string {
		return s.Text()
	}
	return bb.Map(fn), nil
}
