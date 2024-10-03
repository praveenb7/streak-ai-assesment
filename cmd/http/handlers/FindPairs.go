package handlers

import (
	"encoding/json"
	"io"
	"log"
	"main/internal/helper"
	"main/pkg/structs"
	"net/http"
)

func FindPairs(resp http.ResponseWriter, req *http.Request) {
	var requestPayload structs.Request

	resp.Header().Set("Content-Type", "application/json")

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in FindPairs:", r)

			helper.ErrorResponseHelper(resp, 500, "Internal Server Error")
		}
	}()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Error while reading request body:", err.Error())

		helper.ErrorResponseHelper(resp, 400, "Bad Request: Request payload not in expected format")
		return
	}

	err = json.Unmarshal(body, &requestPayload)
	if err != nil {
		log.Println("Error while unmarshalling request body:", err.Error())

		helper.ErrorResponseHelper(resp, 400, "Bad Request: Request payload not in expected format")
		return
	}

	pairs := helper.FindPairsHelper(requestPayload.Numbers, requestPayload.Target)

	response := structs.Response{
		Solutions: pairs,
	}

	pairsJson, err := json.Marshal(response)
	if err != nil {
		log.Println("Error while marshalling the response:", err.Error())

		helper.ErrorResponseHelper(resp, 500, "Internal Server Error")
		return
	}

	resp.WriteHeader(200)
	resp.Write(pairsJson)

}
