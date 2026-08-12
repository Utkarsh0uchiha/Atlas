package handlers

import (
	"html/template"
	"net/http"
)

func NewDashboardHandler(temp *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := temp.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
