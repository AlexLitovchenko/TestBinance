package app

import (
	"AlexLitovchenko/TestBinance/server/internal/hendlers"
	"fmt"
	"net/http"
)

func Run() {
	http.HandleFunc("/api/v1/rates", hendlers.GetRatesHandler)
	fmt.Println("Listening at port 3001...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}
