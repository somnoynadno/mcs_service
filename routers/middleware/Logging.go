package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"mcs_service/db"
	"mcs_service/models/entities"
	"net/http"
)

var LogBody = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Error reading body: %v", err)
			http.Error(w, "Can't parse body", http.StatusBadRequest)
			return
		}

		if len(body) > 0 {
			log.Debug(string(body))
		}

		// And now set a new body, which will simulate the same data we read:
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
}

var LogPath = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodOptions {
			IP := r.Header.Get("X-Real-IP") // depends on nginx
			log.Info(fmt.Sprintf("%s: %s %s (%s)", IP, r.Method, r.RequestURI, r.Host))
			go saveRequest(r)
		}
		next.ServeHTTP(w, r)
	})
}

func saveRequest(r *http.Request) {
	h, err := json.Marshal(r.Header)
	if err != nil{
		log.Error(err)
		return
	}

	sl := entities.ServiceLogs{
		URL: r.RequestURI,
		RealIP: r.Header.Get("X-Real-IP"),
		UserAgent: r.Header.Get("User-Agent"),
		Method: r.Method,
		Headers: string(h),
	}

	DB := db.GetDB()
	err = DB.Create(&sl).Error
	if err != nil {
		log.Error(err)
	}
}