package api

import (
	"fmt"
	"github.com/KLIM8D/wab.io/lib"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	re      = regexp.MustCompile(`(^|\s)((https?:\/\/)?[\w-]+(\.[\w-]+)+\.?(:\d+)?(\/\S*)?)`)
	base    string
	factory *lib.RedisConf
)

func redirectUrl(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	s := req.RequestURI[1:]
	e := &lib.ShortenedURL{}
	if _, err := factory.Get(s, e); err != nil {
		log.Println("Error: ", err)

		data := []byte("failed")

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		res.Write(data)
	} else {
		http.Redirect(res, req, e.Url, 302)
	}
}

func shortenUrl(res http.ResponseWriter, req *http.Request) {
	var response string

	url := req.FormValue("url")

	if url == "" {
		log.Println("Error: no url form value")
	} else {
		if validateUrl(url) {
			exp := req.FormValue("expire")
			key := req.FormValue("key")

			sUrl := fmt.Sprintf("%x", lib.Shortener(url))
			go func() {
				if exists, err := factory.Exists(sUrl); err != nil || exists > 0 {
					if key != "" {
						factory.RPush(key, sUrl)
					}
					return
				}

				item := &lib.ShortenedURL{
					Key:     sUrl,
					Expires: 43200, //12 hours
					Url:     url,
				}

				if exp == "" && key != "" {
					user := &lib.User{}
					if _, err := factory.Get(key, user); user != nil && err == nil {
						item.Expires = (time.Duration(user.Expires) * time.Minute).Seconds()
					}
				} else if exp != "" {
					if f, err := strconv.ParseFloat(exp, 64); err == nil {
						item.Expires = (time.Duration(f) * time.Minute).Seconds()
					}
				}

				factory.Add(item)
			}()

			response = base + sUrl
		} else {
			log.Printf("Not valid URL: %s", url)
			response = base + "NotValidURL"
		}
	}

	data := []byte(response)

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Write(data)
}

func validateUrl(b string) bool {
	return re.MatchString(b)
}
