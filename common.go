package main

import (
	"errors"
	"log"
	"time"

	"github.com/beevik/etree"
)

func fixDate(resp []byte) ([]byte, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(resp); err != nil {
		return nil, err
	}

	root := doc.SelectElement("rss")
	els := root.FindElements("./channel/item")
	if len(els) == 0 {
		return nil, errors.New("token is error or no valid subscribe")
	}

	for i := range els {
		el := els[i]
		elPubDate := el.FindElement("./torrent/pubDate")
		pubTime, err := time.ParseInLocation("2006-01-02T15:04:05", elPubDate.Text(), time.FixedZone("CST", 60*60*8))
		if err != nil {
			log.Print(err)
		}
		el.CreateElement("pubDate").SetText(pubTime.Format(time.RFC1123Z))
	}

	b, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}

	return b, nil
}
