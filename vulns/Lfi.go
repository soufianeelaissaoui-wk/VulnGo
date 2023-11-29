package vulns

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

const ROOT string = "/tmp/%s"

func Lfi(w http.ResponseWriter, r *http.Request) {
    file_name := r.URL.Query().Get("x")

    path := fmt.Sprintf(ROOT, file_name)

    ioutil.ReadFile(path)

}