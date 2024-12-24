package handlers

import (
	"encoding/json"
	"net/http"

	"goapi/api"
	"goapi/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

// GetCoinBalance handles requests to check a user's coin balance.
// It expects a username parameter in the URL query string and returns
// the user's current coin balance in JSON format.
// On error, it returns appropriate HTTP error codes and messages.
func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
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

	// var tokenDetails *tools.CoinDetails
	tokenDetails := (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
