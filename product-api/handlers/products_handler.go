package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"ytb_nic_jackson-bld-mcrsrv-golang/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler interface
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	} else if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		grp := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(grp) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(grp[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := grp[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println("Found id", id)
	}

	// catch all
	//if no method is matched
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//fetch the products from the data store
	lp := data.GetProducts()

	//serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}
