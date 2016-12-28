package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	"github.com/PuerkitoBio/goquery"
	"github.com/garyburd/redigo/redis"
)

const (
	httpFolder = "http://bitly.com/nuvi-plz"
	redisSrv   = "127.0.0.1:6379"
	index      = "s" //"ProcessedRecords"
	data       = "l" //"NEWS_XML"
)

type xml struct {
	xmlfilename string
	xmldata     string
}

func main() {
	conn, err := redis.Dial("tcp", redisSrv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Hello Redis")

	/*
		 //// testing redirect function
		//url := "https://www.google.com"
		url := "http://bitly.com/nuvi-plz"

		r, err := getRedirectUrl(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(r)
	*/

	/*
		////testing scrape function
		//url := "http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/"
		//url := `http://localhost:8080/test.html`
		rec, _ := scrape(httpFolder, "table tr td")
		for i, file := range rec {
			fmt.Println(i+1, file)
		}
	*/

	/*
		////testing download function
		fileName := "1482964046455.zip"
		url := "http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/"
		if err := download(fileName, url+fileName); err != nil {
			fmt.Println(err)
		}
	*/

	/*
		 //// testing data.save function
		data := xml{}
		data.xmlfilename = "00a5e539322693f3"
		data.xmldata = "......................... 00a5e539322693f39d9923e1967 ....................."

		if err := data.save(conn); err != nil {
			fmt.Println(err.Error())
		}*/

	/*
		cleanup(conn)
		return
	*/
}

func (x xml) save(conn redis.Conn) error {
	resp, err := conn.Do("SISMEMBER", index, x.xmlfilename)
	if err != nil {
		return err
	} else if reflect.DeepEqual(resp, int64(1)) {
		return errors.New("Record already exists, so skipped " + x.xmlfilename)
	}

	//fmt.Printf("ret type: %T, ret: %v\n", resp, resp) // for debugging

	resp, err = conn.Do("LPUSH", data, x.xmldata)
	if err != nil {
		return err
	}

	resp, err = conn.Do("SADD", index, x.xmlfilename)
	if err != nil {
		return err
	}

	return nil
}

func cleanup(conn redis.Conn) {
	for {
		resp, err := conn.Do("SPOP", index, 1)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if reflect.DeepEqual(resp, []interface{}{}) {
			break
		}

	}

	for {
		resp, err := conn.Do("LPOP", data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if resp == nil {
			break
		}
	}
}

func scrape(url string, path string) (children []string, err error) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	// Find the table items
	//path := ".sidebar-reviews article .content-block"
	//path := "table tr td"

	doc.Find(path).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//a, _ := s.Find("a").Attr("href")
		text := s.Find("a").Text()

		if attr, exists := s.Find("a").Attr("href"); exists && text != "Parent Directory" {

			children = append(children, attr)
		}

	})
	//fmt.Println(count, "records")
	return
}

func download(localpath string, url string) (err error) {
	fmt.Println("downloading", url)

	out, err := os.Create(localpath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}

func getRedirectUrl(url string) (redirectUrl string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	if response, err := client.Do(req); err != nil {
		if response.StatusCode/100 == 3 { // if there is a redirect
			if url, err := response.Location(); err != nil {
				return "", err
			} else {
				return url.String(), nil
			}
		} else {
			return "", err
		}
	}

	redirectUrl = url
	return
}
