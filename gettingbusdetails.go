package main

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
)

type busdet struct {
	rawdet   string
	routenum string
	servtype string
	origin   string
	dest     string
	jmin     int
	stname   []string
}

//Contacts remote host with query and creates a struct with details of
//the bus
//INPUT : bus number 
//OUTPUT: struct with the bus details 
func newBus(bef, bus, aft string) (*busdet, error) {
	var bd *busdet
	url := bef + bus + aft
	d, e := goquery.NewDocument(url)
	if e != nil {
		return nil, e
	}
	r := d.Find("[BGColor='#EAEAEA'],[BGColor='white']").
		Not(":contains('Route')").Text()
	bd = &busdet{rawdet: r, routenum: bus}
	setStageNames(bd)
	setOrigin(bd)
	setDest(bd)
	setServiceType(bd)
	setJourneyTime(bd)
	return bd, nil
}

func setServiceType(b *busdet) {
	// regpat := regexp.MustCompile(b.routenum+"([A-Za-z ]+?)[A-Z]")
	regpat := regexp.MustCompile(b.routenum + `([\w ]+?)` + b.origin)
	b.servtype = regpat.FindStringSubmatch(b.rawdet)[1]
}

func setJourneyTime(b *busdet) {
	// regpat := regexp.MustCompile(b.routenum+`[A-Za-z ]+?[A-Z. ]+?(\d+?)\d\.`)
	regpat := regexp.MustCompile(b.routenum + b.servtype + b.origin + b.dest + `(\d+?)\d\.`)
	b.jmin, _ = strconv.Atoi(regpat.FindStringSubmatch(b.rawdet)[1])
}

//Set the intermediate points between destination and origin
func setStageNames(b *busdet) {
	actreg := regexp.MustCompile(b.routenum + `.+?(\d\..*)`)
	temp := actreg.FindStringSubmatch(b.rawdet)[1]
	regpat := regexp.MustCompile(`(.*?\.?)(?:\d\.|$)`)
	for _, v := range regpat.FindAllStringSubmatch(temp, -1)[1:] {
		b.stname = append(b.stname, v[1])
	}
}

func setOrigin(b *busdet) {
	b.origin = b.stname[0]
}

func setDest(b *busdet) {
	b.dest = b.stname[len(b.stname)-1]
}
