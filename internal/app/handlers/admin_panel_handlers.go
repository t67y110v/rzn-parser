package handlers

import (
	"encoding/json"
	"net/http"
	"restApi/internal/app/utils"
)

func (h *Handlers) HandleAdminAccess() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			h.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := h.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			h.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			h.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusUnauthorized, err)
			return
		}
		if err != nil {
			h.error(w, r, http.StatusBadRequest, errorIncorrectEmailOrPassword)
			h.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}

		if !utils.CheckThatUserIsAdmin(u) {
			h.error(w, r, http.StatusBadRequest, errorThisUserIsNotAdmin)
			h.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusBadRequest, err)

			return
		} else {

			type resp struct {
				Result bool `json:"result"`
			}
			res := &resp{}

			res.Result = true

			h.respond(w, r, http.StatusOK, res)
			h.logger.Infof("handle /adminAccess, status :%d", http.StatusOK)
		}

	}
}
