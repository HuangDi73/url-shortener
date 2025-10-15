package stat

import (
	"net/http"
	"time"
	"url-shortener/config"
	"url-shortener/pkg/middleware"
	"url-shortener/pkg/res"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type handler struct {
	StatRepo *Repository
	*config.Config
}

type HandlerDeps struct {
	StatRepo *Repository
	*config.Config
}

func NewHandler(mux *http.ServeMux, deps HandlerDeps) {
	h := handler(deps)
	mux.Handle("GET /stat", middleware.IsAuthed(h.GetStat(), h.Config))
}

func (h *handler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		stats := h.StatRepo.GetStats(by, from, to)
		res.Json(w, stats, http.StatusOK)
	}
}
