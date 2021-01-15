package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/auth"
	"github.com/yeldisbayev/education/architecture/clean/golang/cinema-example/entity"
)

// Handler struct
type Handler struct {
	useCase auth.UseCase
}

// NewHandler constructor
func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Username string `json:"Username" validate:"required,min=5,max=25"`
	Password string `json:"Password" validate:"required,min=8,max=50"`
}

// SignUp handler
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	inp := new(signInput)

	if err := json.NewDecoder(r.Body).Decode(inp); err != nil {
		w.WriteHeader(400)
		return
	}

	validate := validator.New()

	if err := validate.Struct(inp); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(
			entity.NewQueryFormError(err),
		)

		return

	}

	if err := h.useCase.SignUp(r.Context(), inp.Username, inp.Password); err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)

}
