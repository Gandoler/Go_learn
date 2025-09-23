package save

import (
	"errors"
	"log/slog"
	"net/http"
	resp "url_shortener/internal/lib/api/responce"
	"url_shortener/internal/lib/logger/sl"
	"url_shortener/internal/lib/random"
	"url_shortener/internal/storage"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveUrl(urlToSave string, alias string) (int64, error)
}

const aliasLength = 6

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("%s: %s", op, sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			var validateErr validator.ValidationErrors
			errors.As(err, &validateErr)
			log.Error("%s: %s", op, sl.Err(err))

			render.JSON(w, r, resp.Error("invalid request"))
			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias, err = random.RandomStringUrl(aliasLength)
			if err != nil {
				log.Error("%s: %s", op, sl.Err(err))
			}
		}

		id, err := urlSaver.SaveUrl(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info("url already exists %s: %s", op, sl.Err(err), slog.String("url", req.URL))

			render.JSON(w, r, resp.Error("url already exists"))
			return
		}
		if err != nil {
			log.Error("%s: %s", op, sl.Err(err))
			render.JSON(w, r, resp.Error("failed to save url"))
			return
		}

		log.Info("saved url", slog.Int64("id", id), slog.String("url", req.URL))

		ResponseOk(w, r, alias)

	}
}

func ResponseOk(w http.ResponseWriter, r *http.Request, alias string) {
	response := Response{Response: resp.OK(), Alias: alias}
	render.JSON(w, r, response)
}
