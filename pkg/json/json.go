package json

import (
	"encoding/json"
	"fmt"
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

func Write(w http.ResponseWriter, status int, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println("[DEBUG] marshal error:", err)
		return err
	}
	fmt.Println("[DEBUG] json size:", len(data), "bytes")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	n, err := w.Write(data)
	fmt.Println("[DEBUG] written bytes:", n, "err:", err)
	return err
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]string{"error": msg}
	data, _ := json.Marshal(resp)
	w.Write(data)
}
