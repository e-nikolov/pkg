package xhttp

import (
	"io"
	"encoding/json"
	"github.com/e-nikolov/pkg/xerror"
)

func WriteAsJSON(w io.Writer, obj interface{}) {
	bytes, err := json.Marshal(obj)
	xerror.PanicIfError(err)

	w.Write(bytes)
}
