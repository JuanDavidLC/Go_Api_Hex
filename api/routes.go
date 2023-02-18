package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/JuanDavidLC/Go_Api_Hex/infrastructure/controllers"
)

type user struct {
	handler controllers.UserHandler
}

func initRoutes(handler controllers.UserHandler) {

	http.Handle("/user/", &user{handler: handler})

}

func (u *user) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	userMethods(w, r, u.handler)

}

func userMethods(w http.ResponseWriter, r *http.Request, handler controllers.UserHandler) {

	switch r.Method {

	case http.MethodPost:

		url, err := r.URL.Parse(r.URL.Path)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		base := strings.Split(url.String(), "/")
		if base[2] != "" {

			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "The POST method is not supported for this route")

			return

		}

		handler.Store(w, r)

	case http.MethodGet:

		url, err := r.URL.Parse(r.URL.Path)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		base := strings.Split(url.String(), "/")
		if base[2] != "" {

			id, err := strconv.ParseInt(base[2], 10, 64)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			handler.GetById(w, r, id)
			return

		}

		handler.Get(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

}
