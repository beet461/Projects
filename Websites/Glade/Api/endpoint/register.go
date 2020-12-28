package gladeapi

import "net/http"

func register(w http.ResponseWriter, r *http.Request) {
	var data := unmarshalData(w, r)
}
