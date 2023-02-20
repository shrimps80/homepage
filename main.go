package main
import (
"io/ioutil"
"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  content, _ := ioutil.ReadFile("./index.html")
  w.Write(content)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8082", nil)
}
