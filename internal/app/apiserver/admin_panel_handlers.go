package apiserver

import (
	"encoding/json"
	"net/http"
	"restApi/internal/app/utils"
)

func (s *Server) handleAdminAccess() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusUnauthorized, err)
			return
		}
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}

		if !utils.CheckThatUserIsAdmin(u) {
			s.error(w, r, http.StatusBadRequest, errorThisUserIsNotAdmin)
			s.logger.Warningf("handle /adminAccess, status :%d, error :%e", http.StatusBadRequest, err)

			return
		} else {

			type resp struct {
				Result bool `json:"result"`
			}
			res := &resp{}

			res.Result = true

			s.respond(w, r, http.StatusOK, res)
			s.logger.Infof("handle /adminAccess, status :%d", http.StatusOK)
		}

	}
}
