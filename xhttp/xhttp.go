package xhttp

import (
	"encoding/json"
	"io"

	"github.com/e-nikolov/pkg/xerrors"
)

func WriteAsJSON(w io.Writer, obj interface{}) {
	bytes, err := json.Marshal(obj)
	xerrors.PanicIfError(err)

	w.Write(bytes)
}
