package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	"github.com/kishor82/goapi/api"
	"github.com/kishor82/goapi/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	params := api.CoinBalanceParamas{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface

	database, err = tools.NewDatabase()

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	response := api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
