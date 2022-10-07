package services

import (
	"crypto/rand"
	"dtskominfo-gin/apimodels"
	"dtskominfo-gin/database"
	"dtskominfo-gin/models"
	"encoding/json"
	"fmt"
	"time"
)

func SaveOrder(req apimodels.Request) (apimodels.Response, error) {
	var res apimodels.Response
	PrettyPrint(req)
	db := database.GetDb()
	var items []models.Item
	var total int64
	for _, vitem := range req.Items {
		var item models.Item
		item.Price = vitem.Price
		item.Quantity = int(vitem.Quantity)
		item.ItemCode = vitem.ItemCode
		item.Description = vitem.Description
		items = append(items, item)
		total += (vitem.Quantity * vitem.Price)
	}

	order := models.Order{
		OrderID:      generateRandomID(),
		CustomerName: req.CustomerName,
		OrderAt:      currentTime(),
		DetaiItem:    items,
	}

	errdb := db.Create(&order).Error
	if errdb != nil {
		return res, errdb
	}

	return apimodels.Response{
		Data:         req,
		DateTrans:    fmt.Sprintf("%v", dateTimeEpoch(currentTime())),
		OrderID:      order.OrderID,
		ResponseCode: "00",
		Status:       "Success",
		Total:        total,
	}, nil

}

func generateRandomID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("%v", currentTime())
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// epoch time
func currentTime() int64 {
	return time.Now().Unix()
}

func dateTimeEpoch(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
