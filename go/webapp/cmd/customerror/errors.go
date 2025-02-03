package customerror

import (
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/maslias/webapp/cmd/utils"
)

var (
	ErrInternalFailure = errors.New("internal failure")
	ErrBadRequest      = errors.New("bad request")
	ErrNotFound        = errors.New("not found")
    ErroUnauth  = errors.New("unauthorized")
)

type HybridError struct {
	appErr    error
	sourceErr error
}

func (e HybridError) Error() string {
	return errors.Join(e.appErr, e.sourceErr).Error()
}

func (e HybridError) AppErr() error {
	return e.appErr
}

func (e HybridError) SourceErr() error {
	return e.sourceErr
}

func NewHybridError(appErr, sourceErr error) error {
	return HybridError{
		appErr:    appErr,
		sourceErr: sourceErr,
	}
}

type ErrorHttpHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorBridge(er ErrorHttpHandler, logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := er(w, r); err != nil {

			status := http.StatusInternalServerError
			outErr := err

			var hybirdError HybridError
			if errors.As(err, &hybirdError) {
				appErr := hybirdError.AppErr()
				outErr = hybirdError.SourceErr()
				switch appErr {
				case ErrBadRequest:
					status = http.StatusBadRequest
				case ErrInternalFailure:
					status = http.StatusInternalServerError
				case ErrNotFound:
					status = http.StatusNotFound
				case ErroUnauth:
					status = http.StatusNonAuthoritativeInfo
				}
			}

			_ = utils.WriteJSONError(w, status, outErr)

			start := time.Now()
			logger.Errorw(
				err.Error(),
				"status",
				status,
				"method",
				r.Method,
				"path",
				r.URL.Path,
				"duration",
				time.Since(start),
			)
		}
	}
}
