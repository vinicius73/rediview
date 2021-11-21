package httputil

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       interface{}
	Headers    map[string]string
}

func (r Response) Write(w http.ResponseWriter) {
	for k, v := range r.Headers {
		w.Header().Set(k, v)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(r.StatusCode)
	_, _ = w.Write(r.body())
}

func (r Response) body() []byte {
	j, _ := json.Marshal(r.Body)

	return j
}

func BuildResponse(status int, body interface{}) Response {
	return Response{
		StatusCode: status,
		Body:       body,
	}
}

func JSON(w http.ResponseWriter, statusCode int, value interface{}) error {
	if value == nil {
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	enc := json.NewEncoder(w)
	if err := enc.Encode(value); err != nil {
		return err
	}

	return nil
}
