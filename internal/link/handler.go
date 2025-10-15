package link

import (
	"fmt"
	"net/http"
	"strconv"
	"url-shortener/config"
	"url-shortener/internal/stat"
	"url-shortener/pkg/middleware"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"
)

type handler struct {
	*config.Config
	LinkRepo *Repository
	StatRepo *stat.Repository
}

type HandlerDeps struct {
	*config.Config
	LinkRepo *Repository
	StatRepo *stat.Repository
}

func NewHandler(mux *http.ServeMux, deps HandlerDeps) {
	h := handler(deps)
	mux.HandleFunc("POST /link", h.Create())
	mux.Handle("GET /links", middleware.IsAuthed(h.GetAll(), h.Config))
	mux.HandleFunc("GET /{hash}", h.GoTo())
	mux.Handle("PATCH /link/{id}", middleware.IsAuthed(h.Update(), h.Config))
	mux.HandleFunc("DELETE /link/{id}", h.Delete())
}

func (h *handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "Empty body", http.StatusBadRequest)
			return
		}
		body, err := req.HandleBody[CreateLink](w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			existedLink, _ := h.LinkRepo.FindByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := h.LinkRepo.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLink, http.StatusCreated)
	}
}

func (h *handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}
		links := h.LinkRepo.GetAll(limit, offset)
		count := h.LinkRepo.Count()
		gotLinks := GetAllLinksResponse{
			Links: *links,
			Count: count,
		}
		res.Json(w, gotLinks, http.StatusOK)
	}
}

func (h *handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		foundLink, err := h.LinkRepo.FindByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		h.StatRepo.AddClick(foundLink.ID)
		http.Redirect(w, r, foundLink.Url, http.StatusTemporaryRedirect)
	}
}

func (h *handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "Empty body", http.StatusBadRequest)
			return
		}
		email, ok := r.Context().Value("CtxEmailKey").(string)
		if ok {
			fmt.Println(email)
		}
		body, err := req.HandleBody[UpdateLink](w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link := &Link{
			ID:   uint(id),
			Url:  body.Url,
			Hash: body.Hash,
		}
		updatedLink, err := h.LinkRepo.Update(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, updatedLink, http.StatusCreated)
	}
}

func (h *handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = h.LinkRepo.FindById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = h.LinkRepo.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, http.StatusOK)
	}
}
