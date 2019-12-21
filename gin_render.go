package extjson

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type GinJSON struct {
	Callback string
	Data     interface{}
}

func GinRender(data interface{}, callback string) GinJSON {
	return GinJSON{Data: data, Callback: callback}
}

var jsonContentType = []string{"application/json; charset=utf-8"}
var jsonpContentType = []string{"application/javascript; charset=utf-8"}
var jsonAsciiContentType = []string{"application/json"}

// Render (JSON) writes data with custom ContentType.
func (r GinJSON) Render(w http.ResponseWriter) (err error) {
	if r.Callback != "" {
		return r.jsonp(w)
	}
	if err = writeJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

// WriteContentType (JSON) writes JSON ContentType.
func (r GinJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// WriteJSON marshals the given interface object and writes it with custom ContentType.
func writeJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func (r GinJSON) jsonp(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	ret, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	if r.Callback == "" {
		_, err = w.Write(ret)
		return err
	}

	callback := template.JSEscapeString(r.Callback)
	_, err = w.Write([]byte(callback))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("("))
	if err != nil {
		return err
	}
	_, err = w.Write(ret)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(")"))
	if err != nil {
		return err
	}

	return nil
}
