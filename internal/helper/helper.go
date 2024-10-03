package helper

import (
	"encoding/json"
	"log"
	"main/pkg/structs"
	"net/http"
)

func ErrorResponseHelper(resp http.ResponseWriter, code int, msg string) {
	errResp := structs.ErrorResponse{
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
	var pairSet map[structs.Pair]bool = make(map[structs.Pair]bool)

	for idx, num := range nums {
		numberMap[num] = idx
	}

	for idx, num := range nums {
		numIdx, exists := numberMap[target-num]
		if exists && idx != numIdx {

			pairArr := []int{min(idx, numIdx), max(idx, numIdx)}

			pair := structs.Pair{A: pairArr[0], B: pairArr[1]}
			_, exists := pairSet[pair]
			if !exists {
				pairs = append(pairs, pairArr)
				pairSet[pair] = true
			}
		}
	}

	return pairs
}
