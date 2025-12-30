package infra

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeff-upton-sp/sp-extension/internal/cmd"
	"github.com/sailpoint/atlas-go/atlas/web"
)

func (s *ExtensionService) buildRoutes() *mux.Router {
	r := web.NewRouter(web.DefaultAuthenticationConfig(s.TokenValidator))
	r.Use(web.CustomAuthorization())

	r.Handle("/invoke", s.invoke()).Methods(http.MethodPost)

	return r
}

func (s *ExtensionService) invoke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		decoder := json.NewDecoder(r.Body)

		var input cmd.InvokeInput
		if err := decoder.Decode(&input); err != nil {
			web.BadRequest(ctx, w, err)
			return
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
