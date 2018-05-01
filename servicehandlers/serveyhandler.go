package servicehandlers

import (
	"fmt"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/structures"
)

type ServeyHandler struct {
}

func (s ServeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(s, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (s ServeyHandler) Get(r *http.Request) SrvcRes {
	param := r.URL.Query().Get("title")
	if param != "" {
		fmt.Println("in if")
		var servey structures.ServeyDetails
		servey = dao.GetServeysByTitle(param)
		return Response200OK(servey)
	}
	param = r.URL.Query()["status"][0]
	fmt.Println(param)
	var servey []structures.ServeyDetails
	servey = dao.GetServeysByStatus(param)
	return Response200OK(servey)
}

func (s ServeyHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (s ServeyHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()

}
