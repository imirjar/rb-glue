package compressor

import (
	"compress/gzip"
	"io"
	"log"
)

type cR struct {
	R   io.ReadCloser
	GzR *gzip.Reader
}

func newCR(r io.ReadCloser) (*cR, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &cR{
		R:   r,
		GzR: zr,
	}, nil
}

func (c cR) Read(p []byte) (n int, err error) {
	return c.GzR.Read(p)
}

func (c *cR) Close() error {
	if err := c.R.Close(); err != nil {
		log.Print(err)
		return err
	}
	return c.GzR.Close()
}
