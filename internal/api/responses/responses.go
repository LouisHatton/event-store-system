package responses

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err        error `json:"-"` // low-level runtime error
	StatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func NotFoundResponse(name string) render.Renderer {
	return &ErrResponse{
		StatusCode: http.StatusNotFound,
		StatusText: fmt.Sprintf("%s not found", name),
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Invalid request",
		ErrorText:  err.Error(),
	}
}

func ErrUnauthorised() render.Renderer {
	return &ErrResponse{
		StatusCode: 401,
		StatusText: "Unauthroised",
	}
}

func ErrForbidden() render.Renderer {
	return &ErrResponse{
		StatusCode: 403,
		StatusText: "Forbidden",
	}
}

func ErrInternalServerError() render.Renderer {
	return &ErrResponse{
		StatusText: "Internal Server Error",
		StatusCode: 500,
	}
}
