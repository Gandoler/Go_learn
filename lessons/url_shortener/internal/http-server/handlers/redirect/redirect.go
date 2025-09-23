package redirect

import (
	"errors"
	"log/slog"
	"net/http"
	resp "url_shortener/internal/lib/api/responce"
	"url_shortener/internal/lib/logger/sl"
	"url_shortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=URLGetter
type UrlGetter interface {
	GetUrl(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		resUrl, err := urlGetter.GetUrl(alias)

		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)

			render.JSON(w, r, resp.Error("not found"))

			return
		}
		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			render.JSON(w, r, resp.Error("internal error"))

			return

		}

		log.Info("got url", slog.String("url", resUrl))

		// redirect to found url
		http.Redirect(w, r, resUrl, http.StatusFound)
	}

}
