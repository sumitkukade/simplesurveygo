package servicehandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/structures"
)

type TakeServeyHandler struct {
}

func (s TakeServeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(s, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (s TakeServeyHandler) Get(r *http.Request) SrvcRes {
	tokens, _ := r.Header["Token"]

	fmt.Println("TOKENS : ", tokens)
	fmt.Println("HEADERS : ", r.Header)

	token := tokens[0]

	user := dao.GetSessionDetails(token)
	userserveys := dao.GetAllUserServeys(user.Username)
	return Response200OK(userserveys)
}

func (s TakeServeyHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (s TakeServeyHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	tokens, _ := r.Header["Token"]

	fmt.Println("TOKENS : ", tokens)
	fmt.Println("HEADERS : ", r.Header)

	token := tokens[0]
	user := dao.GetSessionDetails(token)
	var ans structures.PostServeyInput
	err = json.Unmarshal(body, &ans)
	res := dao.CreateAnswerForUser(ans, user.Username)
	return Response200OK(res)
}
