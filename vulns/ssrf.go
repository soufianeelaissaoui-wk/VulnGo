package vulns

import (
	"log"
	"net/http"
	"io/ioutil"
)

func Ssrf(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("x")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(body)
}