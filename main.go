package main
import (
	. "vulngo/vulns"
	"net/http"
)

func main() {
 http.HandleFunc("/xss", Xss)
 http.ListenAndServe(":5000", nil)
}
