package types

import (
	"bytes"
	"compress/gzip"
	"github.com/json-iterator/go"
)

type Metadata struct {
	Body    []byte   `json:"body,omitempty"`
	Created int64    `json:"created,omitempty"`
	Type    []byte   `json:"type,omitempty"`
	Tags    [][]byte `json:"tags,omitempty"`
}

func (t *Metadata) Encode() ([]byte, error) {
	dt, err := jsoniter.Marshal(t)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	w := gzip.NewWriter(buf)
	if _, err := w.Write(dt); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *Metadata) Decode(dt []byte) error {
	gr, err := gzip.NewReader(bytes.NewReader(dt))
	if err != nil {
		return err
	}
	return jsoniter.NewDecoder(gr).Decode(t)
}
