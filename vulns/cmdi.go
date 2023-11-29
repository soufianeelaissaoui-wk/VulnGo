package vulns

import "os/exec"
import "net/http"

func Cmdi(w http.ResponseWriter, r *http.Request) {
  // Vunerable
  exec.Command("sh", "-c", r.URL.Query().Get("x2")).Output()

  // Vulnerable
  userInput1 := "cat" // value supplied by user input
  userInput2 := r.URL.Query().Get("x1")
  exec.Command(userInput1, userInput2)
}