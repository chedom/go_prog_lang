package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

var itemList = template.Must(template.New("itemList").Parse(`
<h1>Item list, count: {{len .}}</h1>
<table>
	<tr>
		<th>Name</th>
		<th>Price</th>
	</tr>
{{ range $k, $v := .}}
	<tr>
		<th>{{$k}}</th>
		<th>{{$v}}</th>
	</tr>
{{ end }}
</table>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)

	http.HandleFunc("/add_item", db.create)
	http.HandleFunc("/price", db.read)
	http.HandleFunc("/update_item", db.update)
	http.HandleFunc("/delete_item", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
	}

}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	item := queryValues.Get("item")
	rawPrice := queryValues.Get("price")
	price, err := strconv.ParseFloat(rawPrice, 32)

	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price is not correct: %v", err)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item %s already exists", item)
		return
	}

	db[item] = dollars(price)

	fmt.Fprintf(w, "%s: %f\n", item, dollars(price))

}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	price, ok := db[item]
	mu.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	rawPrice := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(rawPrice, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price is not correct: %v", err)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s", item)
	}

	db[item] = dollars(price)

	fmt.Fprintf(w, "%s: %f\n", item, dollars(price))
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	mu.Lock()
	delete(db, item)
	mu.Unlock()

	fmt.Fprintf(w, "ok")

}