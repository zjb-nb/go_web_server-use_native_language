package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserData struct {
	Name string `json:"name"`
	Age  int    `json:"ag"`
}

type RespData struct {
	Data string `json:"data"`
}

func Sign(w http.ResponseWriter, r *http.Request) {
	//用户数据会以json的形式塞在body中传送过来
	u_data := &UserData{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s\n", "ready body failed")
		return
	}

	err = json.Unmarshal(body, u_data)
	//parse json
	if err != nil {
		fmt.Fprintf(w, "%s\n", "parse jsofailed")
		return
	}
	//装数据并返回
	respdata := &RespData{Data: "123"}
	res, err := json.Marshal(respdata)
	if err != nil {
		fmt.Println("send json failed")
		return
	}
	fmt.Fprintf(w, "%s\n", string(res))

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", "hello world")
}
