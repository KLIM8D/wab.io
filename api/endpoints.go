package api

import (
	"fmt"
	"github.com/KLIM8D/wab.io/logs"
	"github.com/KLIM8D/wab.io/utils"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	re      = regexp.MustCompile(`(^|\s)((https?:\/\/)?[\w-]+(\.[\w-]+)+\.?(:\d+)?(\/\S*)?)`)
	base    string
	factory *utils.RedisConf
)

func validateUrl(b string) bool {
	return re.MatchString(b)
}

func redirectUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	s := r.RequestURI[1:]
	e := &utils.ShortenedURL{}
	if _, err := factory.Get(s, e); err != nil {
		logs.Error.Println("Error: ", err.Error())

		data := []byte("Unable to redirect")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	} else {
		logs.Info.Printf("%q redirected to %q", s, e)
		http.Redirect(w, r, e.Url, 302)
	}
}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	var response string

	url := r.FormValue("url")

	if url == "" {
		logs.Info.Println("Error: no url form value")
	} else {
		if validateUrl(url) {
			exp := r.FormValue("expire")
			key := r.FormValue("key")

			sUrl := fmt.Sprintf("%x", utils.Shortener(url))
			go func() {
				if exists, err := factory.Exists(sUrl); err != nil || exists > 0 {
					if key != "" {
						logs.Trace.Printf("Added %q to key: %q", sUrl, key)
						factory.RPush(key, sUrl)
					}
					return
				}

				item := &utils.ShortenedURL{
					Key:     sUrl,
					Expires: 43200, //12 hours
					Url:     url,
				}

				if exp == "" && key != "" {
					user := &utils.User{}
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
			logs.Warning.Printf("Not valid URL: %q", url)
			response = base + "NotValidURL"
		}
	}

	data := []byte(response)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}
