package main

import "github.com/PuerkitoBio/goquery"

func getBusList() []string {
  const bussite = `http://www.mtcbus.org/Routes.asp`
  var b *goquery.Document
  var bb *goquery.Selection
  var fn func(int, *goquery.Selection) string

  b, e := goquery.NewDocument(bussite)
  if e != nil {
    panic(e)
  }
  bb = b.Find("Option")
  bb = bb.Not(":empty")
  fn = func(i int, s *goquery.Selection) string {
    return s.Text()
  }
  return bb.Map(fn)
}
