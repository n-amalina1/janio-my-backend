package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"janio-my-backend/models"
	"net/http"
)

func GetOrdersMYProvider(db *sql.DB) (models.MYProviderOrdersParams, error) {
	var (
		MYProviderOrdersParams models.MYProviderOrdersParams
	)

	res, err := http.Get("http://localhost:8080/my/orders")

	if err != nil {
		panic(err.Error())
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &MYProviderOrdersParams)

	return MYProviderOrdersParams, nil
}

func PostStatusMYProvider(db *sql.DB, status models.MYOrderStatus) (models.MYOrderStatus, error) {
	json, err := json.Marshal(
		models.MYOrderStatus{
			OrderID:     status.OrderID,
			OrderStatus: status.OrderStatus,
		})
	if err != nil {
		return models.MYOrderStatus{}, fmt.Errorf("update Status DB: %v", err)
	}

	_, errS := http.Post("http://localhost:8080/id/order/update", "application/json", bytes.NewBuffer(json))
	if errS != nil {
		return models.MYOrderStatus{}, fmt.Errorf("update Status DB: %v", errS)
	}

	return status, nil
}
