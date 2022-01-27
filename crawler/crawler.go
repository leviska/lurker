package crawler

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/leviska/lurker/base"
)

type Crawler struct {
	wait   sync.WaitGroup
	client base.Clienter
	saver  base.Saver
}

func NewCrawler(saver base.Saver, client base.Clienter) *Crawler {
	c := &Crawler{
		client: client,
		saver:  saver,
	}
	return c
}

func (c *Crawler) Stop() {
	c.wait.Wait()
}

func (c *Crawler) Request(req *base.Request) {
	c.wait.Add(1)
	go func() {
		defer c.wait.Done()

		if err := c.do(req); err != nil {
			log.Printf("couldn't get %v: %v\n", req.URL, err)
		}
	}()
}

func (c *Crawler) do(req *base.Request) error {
	httpReq := c.client.Get(req.URL)
	defer c.client.Put(req.URL)

	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("wrong http status code")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	return req.Parser.Parse(req, doc, c, c.saver)
}
