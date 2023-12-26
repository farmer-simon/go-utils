package utils

import "github.com/farmer-simon/go-utils/excel"

func NewExcel(localFile string, autoCreate bool) (*excel.ExcelFactory, error) {
	return excel.New(localFile, autoCreate)
}
