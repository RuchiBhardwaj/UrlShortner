package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	//"io/ioutil"
	//"net/http"
	"net/url"

	//"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


type Redirect struct {
	Id 		int
	Slug 	string 	`db:"slug" form:"slug"`
	Url  	string	`db:"url" form:"url"`
}



var db, err = sql.Open("mysql", "root:hello123@tcp(127.0.0.1:3306)/url")


func getBySlug(){
	var redirect Redirect
	url := "https://flaviocopes.com/golang-sql-database/"
	row := db.QueryRow("select id, slug, url from redirect where url = ?;", url)
	err = row.Scan(&redirect.Id, &redirect.Slug, &redirect.Url)
	print(redirect.Slug)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			urlOrig := Redirect{}
			urlOrig.short("http://example.com/this-is-a-very-long-url-bla-bla-bla")
			fmt.Println(urlOrig.Slug)
			fmt.Println(urlOrig.Url)

			Slug := urlOrig.Slug
			Url := url

			stmt, err := db.Prepare("insert into redirect (slug, url) values(?,?);")
			if err != nil {
				fmt.Print(err.Error())
			}

			_, err = stmt.Exec(Slug, Url)
			if err != nil {
				fmt.Print(err.Error())
			}

			defer stmt.Close()


		} else {
			panic(err)
		}
	}

}

func getResponseData(Url string) string {
	response, err := http.Get(Url)
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

func (u *Redirect) short(urlOrig string) *Redirect {
	shortUrl, originalUrl := tinyUrlShortener(urlOrig)
	u.Slug = shortUrl
	u.Url = originalUrl
	return u
}



func main(){
	getBySlug()
}
