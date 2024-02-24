package errs

import (
	"encoding/json"
	"github.com/Timasha/customErrs/internal"
)

type Err struct {
	Msg        string `json:"msg"`
	CustomCode int64  `json:"customCode"`
	HttpCode   *int64 `json:"httpCode,omitempty"`
}

func New(msg string, customCode int64, httpCode *int64) *Err {
	return &Err{
		Msg:        msg,
		CustomCode: customCode,
		HttpCode:   httpCode,
	}
}

var defaultParsingType internal.ParsingType = internal.TypeJSON

func (e *Err) Error() (strErr string) {
	switch defaultParsingType {
	case internal.TypeJSON:
		{
			byteErr, marshalErr := json.Marshal(*e)
			if marshalErr != nil {
				return ""
			}
			strErr = string(byteErr)
		}
	}

	return strErr
}

func SetDefaultParsingTypeJSON() {
	defaultParsingType = internal.TypeJSON
}
