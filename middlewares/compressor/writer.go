package compressor

import (
	"compress/gzip"
	"net/http"
)

type cW struct {
	Resp   http.ResponseWriter
	GzResp *gzip.Writer
}

func newCW(w http.ResponseWriter) *cW {
	return &cW{
		Resp:   w,
		GzResp: gzip.NewWriter(w),
	}
}

func (c *cW) Header() http.Header {
	return c.Resp.Header()
}

func (c *cW) Write(p []byte) (int, error) {
	return c.GzResp.Write(p)
}

func (c *cW) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.Resp.Header().Set("Content-Encoding", "gzip")
	}
	c.Resp.WriteHeader(statusCode)
}

func (c *cW) Close() error {
	return c.GzResp.Close()
}
