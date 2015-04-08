package api

import (
	"github.com/KLIM8D/wab.io/logs"
	"github.com/KLIM8D/wab.io/utils"
	"github.com/satori/go.uuid"
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

/*
   Service shortens the URL posted in the form field `url`
   if key and expiration time exists, use posted duration in formfield `expire`
   if only key exists, use the users default settings
   else the systems default 12 hours
*/
func shortenUrl(w http.ResponseWriter, r *http.Request) {
	var response string

	url := r.FormValue("url")

	if url == "" {
		logs.Info.Println("Error: no url form value")
	} else {
		if validateUrl(url) {
			exp := r.FormValue("expire")
			key := r.FormValue("key")

			sUrl := <-utils.ShortUrls
			go func() {
				item := &utils.ShortenedURL{
					Key:     sUrl,
					Expires: 43200, //12 hours
					Url:     url,
				}

				if key != "" {
					if exp == "" {
						user := &utils.User{}
						if _, err := factory.Get(key, user); user != nil && err == nil {
							item.Expires = (time.Duration(user.Expires) * time.Minute).Seconds()
						}
					} else {
						if f, err := strconv.ParseFloat(exp, 64); err == nil {
							item.Expires = (time.Duration(f) * time.Minute).Seconds()
						}
					}

					if _, err := uuid.FromString(key); err == nil {
						if exists, err := factory.Exists(sUrl); err != nil && !exists {
							factory.ActivateUser(key)
						}
					}
				}

				factory.Add(item)
				if key != "" {
					logs.Trace.Printf("Added %q to key: %q", sUrl, key)
					factory.RPush(key, item.Key)
				}
			}()

			response = base + sUrl
		} else {
			logs.Warning.Printf("Invalid URL: %q", url)
			response = base + "InvalidURL"
		}
	}

	data := []byte(response)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}
