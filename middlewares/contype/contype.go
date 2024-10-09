package contype

import (
	"fmt"
	"net/http"
)

func REST(contype string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			getContype := r.Header.Get("Content-Type")
			if getContype != contype {
				errMsg := fmt.Sprintf("wrong content type: expected %s, got %s", contype, getContype)
				http.Error(w, errMsg, http.StatusUnsupportedMediaType)
				return
			}
			w.Header().Set("Content-Type", contype)
			next.ServeHTTP(w, r)
		})
	}
}
