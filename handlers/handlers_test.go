package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jorgesiachoque08/melicoupons/requests"
	"github.com/jorgesiachoque08/melicoupons/resources"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Initialize() {
	s.Router = mux.NewRouter()
	s.initializeRoutes()

}

func (s *Server) Run(serverPort string) {
	fmt.Println("Run server: http://localhost:" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, s.Router))
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/coupon", Coupon).Methods("POST")
	s.Router.HandleFunc("/topFavorites", TopFavorites).Methods("GET")
}

var s Server

func TestMain(m *testing.M) {
	s = Server{}
	s.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestTopFavorites(t *testing.T) {
	favorites := resources.ResponseFavorites{}
	ts := httptest.NewServer(http.HandlerFunc(TopFavorites))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	checkResponseCode(t, 200, res.StatusCode)
	json.Unmarshal(body, &favorites)
	if len(favorites.Favorites) == 0 {
		t.Errorf("Favorite is empty")
	}
}

func TestCoupon(t *testing.T) {
	ResponseCoupons := resources.ResponseCoupons{}
	couponRequest := requests.CouponRequest{[]string{"MCO451563457", "MCO507358090", "MCO559835283", "MCO657747635", "MCO801347755", "MCO587955729"}, 211000}
	ts := httptest.NewServer(http.HandlerFunc(Coupon))
	defer ts.Close()
	json_data, err := json.Marshal(couponRequest)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	checkResponseCode(t, 200, res.StatusCode)
	json.Unmarshal(body, &ResponseCoupons)
	if ResponseCoupons.Total != 201518 {
		t.Errorf("the total is incorrect")
	}

	if len(ResponseCoupons.Item_ids) != 3 {
		t.Errorf("the item number does not correspond to the correct ones.")
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
