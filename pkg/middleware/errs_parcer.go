package middleware

import (
	"encoding/json"
	"github.com/Timasha/customErrs/internal"
	"github.com/Timasha/customErrs/pkg/errs"
)

type ErrsParser struct {
	ParsingType internal.ParsingType
}

func (e *ErrsParser) Parse(strErr string) (err *errs.Err, isCustom bool) {
	switch e.ParsingType {
	case errs.TypeJSON:
		{
			err = new(errs.Err)

			parsingErr := json.Unmarshal([]byte(strErr), err)
			if parsingErr != nil {
				return nil, false
			}

			return err, true
		}
	default:
		{
			return nil, false
		}
	}
}
