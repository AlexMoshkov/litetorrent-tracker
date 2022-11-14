package request_validation

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

var validate = validator.New()

func DecodeAndValidate[T any](reader io.ReadCloser) (T, error) {
	var data T
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return data, err
	}
	if err := validate.Struct(data); err != nil {
		return data, err.(validator.ValidationErrors)
	}
	return data, nil
}
