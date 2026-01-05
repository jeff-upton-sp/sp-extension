package infra

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeff-upton-sp/sp-extension/internal/cmd"
	"github.com/jeff-upton-sp/sp-extension/internal/model"
	"github.com/sailpoint/atlas-go/atlas/web"
)

func (s *ExtensionService) buildRoutes() *mux.Router {
	r := web.NewRouter(web.DefaultAuthenticationConfig(s.TokenValidator))
	r.Use(web.CustomAuthorization())

	r.Handle("/functions/{id}", s.getFunction()).Methods(http.MethodGet)
	r.Handle("/functions/{id}/invoke", s.invoke()).Methods(http.MethodPost)

	return r
}

func (s *ExtensionService) getFunction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		input := cmd.GetFunctionInput{
			FunctionID: model.FunctionID(vars["id"]),
		}

		if err := input.Validate(); err != nil {
			web.BadRequest(ctx, w, err)
			return
		}

		out, err := s.app.GetFunction(ctx, input)
		if err != nil {
			web.InternalServerError(ctx, w, err)
			return
		}

		web.WriteJSON(ctx, w, out)
	}
}

func (s *ExtensionService) invoke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		decoder := json.NewDecoder(r.Body)

		var functionInput json.RawMessage
		if err := decoder.Decode(&functionInput); err != nil {
			web.BadRequest(ctx, w, err)
			return
		}

		input := cmd.InvokeInput{
			FunctionID: model.FunctionID(vars["id"]),
			Input:      functionInput,
		}

		if err := input.Validate(); err != nil {
			web.BadRequest(ctx, w, err)
			return
		}

		out, err := s.app.Invoke(ctx, input)
		if err != nil {
			web.InternalServerError(ctx, w, err)
			return
		}

		web.WriteJSON(ctx, w, out)
	}
}
