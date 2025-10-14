package link

import (
	"net/http"
	"strconv"
	"url-shortener/config"
	"url-shortener/pkg/req"
	"url-shortener/pkg/res"
)

type handler struct {
	*config.Config
	Repo *Repository
}

type HandlerDeps struct {
	*config.Config
	Repo *Repository
}

func NewHandler(mux *http.ServeMux, deps HandlerDeps) {
	h := handler(deps)
	mux.HandleFunc("POST /link", h.Create())
	mux.HandleFunc("GET /{hash}", h.GoTo())
	mux.HandleFunc("PATCH /link/{id}", h.Update())
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
			existedLink, _ := h.Repo.FindByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := h.Repo.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLink, http.StatusCreated)
	}
}

func (h *handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		foundLink, err := h.Repo.FindByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, foundLink.Url, http.StatusTemporaryRedirect)
	}
}

func (h *handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "Empty body", http.StatusBadRequest)
			return
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
		updatedLink, err := h.Repo.Update(link)
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
		_, err = h.Repo.FindById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = h.Repo.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, http.StatusOK)
	}
}
