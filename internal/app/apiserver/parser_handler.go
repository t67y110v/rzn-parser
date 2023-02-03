package apiserver

import (
	"encoding/json"

	"net/http"
	parser "restApi/internal/app/parser"
)

func (s *Server) handleParser() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			s.logger.Warningf("handle /parser, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		count, err := parser.Parser(req.Login, req.Password)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errorIncorrectEmailOrPassword)
			s.logger.Warningf("handle /parser, status :%d, error :%e", http.StatusBadRequest, err)
			return
		}
		type resp struct {
			Result string `json:"result"`
		}
		res := &resp{}
		res.Result = count
		s.respond(w, r, http.StatusOK, res)
		s.logger.Infof("handle /parser, status :%d", http.StatusOK)
	}
}
