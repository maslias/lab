package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ReadJSONMin(w http.ResponseWriter, r *http.Request, payload any) error {
	maxByte := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))

	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(payload)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, payload any) error {
	// check content-type when exist
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mt := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mt != "application/json" {
			return fmt.Errorf("Content-Type is not application/json")
		}
	}

	// check size or payload
	maxByte := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&payload)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return fmt.Errorf(
				"Request body contains badly-formed JSON (at position %d)",
				syntaxErr.Offset,
			)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("Request body contains badly-formed JSON")

		case errors.As(err, &unmarshalErr):
			return fmt.Errorf(
				"Request body contains an invalid value for the %q field (at position %d)",
				unmarshalErr.Field,
				unmarshalErr.Offset,
			)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fn := strings.TrimPrefix(err.Error(), "json: unkown field ")
			return fmt.Errorf("Request body contains unknown field %s", fn)

		case errors.Is(err, io.EOF):
			return fmt.Errorf("Request body must not be empty")

		case err.Error() == "http: request body too large":
			return fmt.Errorf("Request body is to large. Max Size: %v", maxByte)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return fmt.Errorf("Request body must only contain a single JSON object")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteJSONError(w http.ResponseWriter, status int, err error) error {
	type errEnvelope struct {
		Error string `json:"error"`
	}

	return WriteJSON(w, status, errEnvelope{Error: err.Error()})
}
