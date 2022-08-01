package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	status  int
}

//sends a BadRequest response, with a struct ResponseError
func (resp *ResponseError) SendBadRequest(rw http.ResponseWriter, message string) {
	resp.status = http.StatusBadRequest
	resp.Error = "Bad Request"
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(resp.status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(rw, string(output))
}

type ResponseCoupons struct {
	Item_ids []string `json:"item_ids"`
	Total    int      `json:"total"`
	status   int
}

//sends a Ok response, with a struct ResponseCoupons
func (resp *ResponseCoupons) SendOk(rw http.ResponseWriter, item_ids []string, total int) {
	resp.status = http.StatusOK
	resp.Item_ids = item_ids
	resp.Total = total
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(resp.status)
	output, _ := json.Marshal(&resp)
	fmt.Fprintln(rw, string(output))
}

type ResponseFavorites struct {
	Favorites []map[string]int `json:"favorites"`
	status    int
}

//sends a Ok response, with a struct ResponseFavorites
func (resp *ResponseFavorites) SendOk(rw http.ResponseWriter, favorites []map[string]int) {
	resp.status = http.StatusOK
	resp.Favorites = favorites
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(resp.status)
	output, _ := json.Marshal(&resp)
	fmt.Fprintln(rw, string(output))
}
