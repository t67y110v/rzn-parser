package handlers_test

import (
	//"restApi/internal/app/apiserver"
	//"restApi/internal/app/handlers"
	//"restApi/internal/app/logging"
	//"restApi/internal/app/store"
	//"restApi/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"

	//	"restApi/internal/app/handlers"
	//	"restApi/internal/app/logging"
	//	"restApi/internal/app/store/teststore"
	"testing"
	//"github.com/gorilla/mux"
)

// func newServer(store store.UserStore, t []int, logger logging.Logger) *apiserver.Server{

// 	return nil
// }

func Teapot(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusTeapot)
}
func TestUserCreate(t *testing.T) {
	//h := handlers.NewHandlers(teststore.New(), logging.GetLogger())

	req := httptest.NewRequest(http.MethodGet, "/userCreate", nil)
	res := httptest.NewRecorder()
	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got wrong status ")
	}
}
