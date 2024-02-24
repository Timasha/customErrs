package errs

import (
	"encoding/json"
	"github.com/Timasha/customErrs/internal"
)

func Parse(strErr string) (err *Err, isInternal bool) {
	switch defaultParsingType {
	case internal.TypeJSON:
		{
			err = new(Err)

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
