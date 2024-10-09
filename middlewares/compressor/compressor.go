package compressor

import (
	"log"
	"net/http"
	"strings"
)

func Compressing() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

			acceptEncoding := req.Header.Get("Accept-Encoding")
			contentEncoding := req.Header.Get("Content-Encoding")

			//client can read
			supportsGzip := strings.Contains(acceptEncoding, "gzip")
			sendsGzip := strings.Contains(contentEncoding, "gzip")

			if supportsGzip && sendsGzip {
				// log.Println("supportsGzip")
				cResp := newCW(resp)
				defer cResp.Close()
				resp = cResp
			}

			if sendsGzip {
				// log.Println("sendsGzip")
				cr, err := newCR(req.Body)
				if err != nil {
					log.Print(err)
					resp.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer cr.Close()
				req.Body = cr
			}

			next.ServeHTTP(resp, req)

		})
	}
}
