package utils

import (
	"math"

	"gorm.io/gorm"
)

func FilterPaginate(sqlSess *gorm.DB, page, limit int64) *gorm.DB {
	return sqlSess.
		Offset(int((page - 1) * limit)).
		Limit(int(limit))
}

func CountMaxPage(totalData, limit int64) int64 {
	if totalData == 0 || limit == 0 {
		return 1
	}

	ceil := math.Ceil(float64(totalData) / float64(limit))
	return int64(ceil)
}
