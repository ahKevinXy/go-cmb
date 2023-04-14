package notice

import (
	"encoding/json"
	"github.com/ahKevinXy/go-cmb/models"
)

func Pay(s string) (*models.PayResultNotice, error) {
	var pay *models.PayResultNotice

	if err := json.Unmarshal([]byte(s), &pay); err != nil {
		return nil, err
	}

	return pay, nil
}
