package middleware

import "github.com/Timasha/customErrs/pkg/errs"

type BaseBodyWithErr struct {
	Err *errs.Err `json:"err" validate:"required"`
}

func (e *ErrsParser) BodyWithErrInterceptor(
	body string,
) (errBody BaseBodyWithErr) {
	customErr, isCustom := e.Parse(body)
	if isCustom {
		return BaseBodyWithErr{customErr}
	}
	return
}
