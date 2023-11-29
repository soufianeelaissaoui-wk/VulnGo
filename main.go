package main
import (
	. "vulngo/vulns"
	"net/http"
)

func main() {
 http.HandleFunc("/xss", Xss)
 http.HandleFunc("/xss2", Xss2)
 http.HandleFunc("/cmdi", Cmdi)
 http.HandleFunc("/sqli", Sqli)
 http.HandleFunc("/lfi", Lfi)
 http.HandleFunc("/ssrf", Ssrf)

 http.ListenAndServe(":5000", nil)
}
