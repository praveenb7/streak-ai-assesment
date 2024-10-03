package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Solutions [][]int `json:"solutions"`
}

type Pair struct {
	a int
	b int
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func errorResponseHelper(resp http.ResponseWriter, code int, msg string) {
	errResp := ErrorResponse{
		Code: code,
		Msg:  msg,
	}

	errRespJson, err := json.Marshal(errResp)
	if err != nil {
		log.Println("Error while marshalling the errResp:", err.Error())

		resp.WriteHeader(500)
		resp.Write([]byte("Internal Server Error"))
		return
	}

	resp.WriteHeader(code)
	resp.Write(errRespJson)
}

func FindPairsHelper(nums []int, target int) [][]int {
	var pairs [][]int

	var numberMap map[int]int = make(map[int]int)
	var pairSet map[Pair]bool = make(map[Pair]bool)

	for idx, num := range nums {
		numberMap[num] = idx
	}

	for idx, num := range nums {
		numIdx, exists := numberMap[target-num]
		if exists && idx != numIdx {

			pairArr := []int{min(idx, numIdx), max(idx, numIdx)}

			pair := Pair{a: pairArr[0], b: pairArr[1]}
			_, exists := pairSet[pair]
			if !exists {
				pairs = append(pairs, pairArr)
				pairSet[pair] = true
			}
		}
	}

	return pairs
}

func FindPairs(resp http.ResponseWriter, req *http.Request) {
	var requestPayload Request

	resp.Header().Set("Content-Type", "application/json")

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in FindPairs:", r)

			errorResponseHelper(resp, 500, "Internal Server Error")
		}
	}()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Error while reading request body:", err.Error())

		errorResponseHelper(resp, 400, "Bad Request: Request payload not in expected format")
		return
	}

	err = json.Unmarshal(body, &requestPayload)
	if err != nil {
		log.Println("Error while unmarshalling request body:", err.Error())

		errorResponseHelper(resp, 400, "Bad Request: Request payload not in expected format")
		return
	}

	pairs := FindPairsHelper(requestPayload.Numbers, requestPayload.Target)

	response := Response{
		Solutions: pairs,
	}

	pairsJson, err := json.Marshal(response)
	if err != nil {
		log.Println("Error while marshalling the response:", err.Error())

		errorResponseHelper(resp, 500, "Internal Server Error")
		return
	}

	resp.WriteHeader(200)
	resp.Write(pairsJson)

}

func main() {
	port := "8090"

	r := mux.NewRouter()

	r.HandleFunc("/find-pairs", FindPairs)

	log.Println("Listening on port:", port)
	http.ListenAndServe(":"+port, r)

}
