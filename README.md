# UrlShortner
A URL shortner written in Golang.

**Files**
--
sql.go
--
URL shortner using sql and go.

You can create your table using this :
--

CREATE TABLE redirect ( id int NOT NULL AUTO_INCREMENT, slug varchar(200) collate utf8mb4_unicode_ci NOT NULL, url varchar(1000) collate utf8mb4_unicode_ci NOT NULL, PRIMARY KEY (id) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='URL shortener Table';

Edit your database connection on this code:
--

sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db-redirects")

sql.Open("mysql", "username:password@tcp(yourIpAddress:port)/yourDataabse")

--
Run using this command -> go run sql.go

Below command will Install all the dependencies recursively.

go get -d ./...





test.go
--

go run test.go 

Here I have used these to shorten Url

https://tinyurl.com/

https://is.gd/







