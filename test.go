package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
)

const (
	TINY_URL = 1
	IS_GD    = 2
)

type UrlShortener struct {
	ShortUrl    string
	OriginalUrl string
}

func getResponseData(urlOrig string) string {
	response, err := http.Get(urlOrig)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	return string(contents)
}

func tinyUrlShortener(urlOrig string) (string, string) {
	escapedUrl := url.QueryEscape(urlOrig)
	tinyUrl := fmt.Sprintf("http://tinyurl.com/api-create.php?url=%s", escapedUrl)
	return getResponseData(tinyUrl), urlOrig
}


func (u *UrlShortener) short(urlOrig string) *UrlShortener {
		shortUrl, originalUrl := tinyUrlShortener(urlOrig)
		u.ShortUrl = shortUrl
		u.OriginalUrl = originalUrl
		return u
}

func main() {
	urlOrig := UrlShortener{}
	urlOrig.short("http://example.com/this-is-a-very-long-url-bla-bla-bla")
	fmt.Println(urlOrig.ShortUrl)
	fmt.Println(urlOrig.OriginalUrl)



}