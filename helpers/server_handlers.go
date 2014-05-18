package lifecycle_helpers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/abhishekkr/gol/golhashmap"
)

func Index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, _ := template.ParseFiles("helpers/public/index.html")
	t.Execute(w, nil)
}

func Data(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	node := req.FormValue("node")
	if node == "" {
		node = "node"
	}
	for _, child := range ChildNodes(node) {
		response := fmt.Sprintf("<h3>%s</h3><br/>", child)
		w.Write([]byte(response))
		w.Write([]byte("<ul>"))
		response = ZmqRead("tsds", fmt.Sprintf("%s:%s", node, child))
		hmap := golhashmap.Csv_to_hashmap(response)
		for key, val := range hmap {
			stats := fmt.Sprintf("<li>%s: <b>%s</b></li>", key, val)
			w.Write([]byte(stats))
		}
		w.Write([]byte("</ul>"))
	}
}

func F1(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, _ := template.ParseFiles("helpers/public/help.html")
	t.Execute(w, nil)
}

func Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, _ := template.ParseFiles("helpers/public/status.html")
	t.Execute(w, nil)
}
