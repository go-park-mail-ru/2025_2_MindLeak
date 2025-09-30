package json

import (
	"encoding/json"
	"io"
	"net/http"
)

func Read(r *http.Request, int interface{}) error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &int)
}

func Write(w http.ResponseWriter, status int, int interface{}) error {
	resp, err := json.Marshal(int)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	//w.WriteHeader(status)
	return nil
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]string{"error": msg}
	data, _ := json.Marshal(resp)
	w.Write(data)
}
