package vulns
import (
	"io"
	"net/http"
)


func Xss(w http.ResponseWriter, r *http.Request) {
 io.WriteString(w, r.URL.Query().Get("x"))
}