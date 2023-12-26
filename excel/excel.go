package excel

import (
	"errors"
	"github.com/xuri/excelize/v2"
	"os"
)

// excelize 功能非常多，经常忘记一些常用方法，
// 对其进行二次封装，将一些常用的功能汇总

type ExcelFactory struct {
	//文件句柄
	File *excelize.File
}

func New(localFile string, autoCreate bool) (*ExcelFactory, error) {
	var f *excelize.File
	//检查文件是否存在
	_, err := os.Stat(localFile)
	if os.IsNotExist(err) {
		if autoCreate {
			//创建文件
			f = excelize.NewFile()
			f.SaveAs(localFile)
		} else {
			return nil, errors.New("Excel文件不存在")
		}
	} else {
		//打开文件
		f, err = excelize.OpenFile(localFile)
		if err != nil {
			return nil, errors.New("打开文件失败")
		}
	}

	return &ExcelFactory{
		File: f,
	}, nil
}

// NewSheet 创建工作表
func (f *ExcelFactory) NewSheet(sheetName string) (int, error) {
	index, err := f.File.NewSheet(sheetName)
	if err != nil {
		return index, errors.New("创建sheet失败")
	}
	return index, nil
}

// DeleteSheet 删除工作表
func (f *ExcelFactory) DeleteSheet(sheetName string) error {
	err := f.File.DeleteSheet(sheetName)
	if err != nil {
		return errors.New("删除sheet失败")
	}
	return nil
}

// SetActiveSheet 设置当前活动工作表
func (f *ExcelFactory) SetActiveSheet(index int) {
	f.File.SetActiveSheet(index)
}

// GetActiveSheetIndex 获取当前活动工作表索引
func (f *ExcelFactory) GetActiveSheetIndex() int {
	return f.File.GetActiveSheetIndex()
}

// GetSheetIndex 获取工作表索引
func (f *ExcelFactory) GetSheetIndex(sheetName string) (int, error) {
	return f.File.GetSheetIndex(sheetName)
}

// GetSheetName 获取工作表名
func (f *ExcelFactory) GetSheetName(index int) string {
	return f.File.GetSheetName(index)
}

// SetRow 写入一行数据
func (f *ExcelFactory) SetRow(sheetName, cell string, slice []interface{}) error {
	return f.File.SetSheetRow(sheetName, cell, &slice)
}

// SetColumn 写入一列数据
func (f *ExcelFactory) SetColumn(sheetName, cell string, slice []interface{}) error {
	return f.File.SetSheetCol(sheetName, cell, &slice)
}

// SetCell 设置单元格的值
func (f *ExcelFactory) SetCell(sheetName, cell string, value interface{}) error {
	return f.File.SetCellValue(sheetName, cell, value)
}

// GetCellValue 获取单元格的值
func (f *ExcelFactory) GetCellValue(sheetName, cell string) (string, error) {
	return f.File.GetCellValue(sheetName, cell)
}

// GetCols 按列获取全部单元格的值
func (f *ExcelFactory) GetCols(sheetName string) ([][]string, error) {
	return f.File.GetCols(sheetName)
}

// GetRows 按行获取全部单元格的值
func (f *ExcelFactory) GetRows(sheetName string) ([][]string, error) {
	return f.File.GetRows(sheetName)
}

// Save 保存Excel
func (f *ExcelFactory) Save() {
	f.File.Save()
}

// Close 关闭Excel(自动保存)
func (f *ExcelFactory) Close() {
	f.Save()
	defer f.File.Close()
}
