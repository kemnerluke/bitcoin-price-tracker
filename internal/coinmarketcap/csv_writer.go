package coinmarketcap

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

func WriteCoinsToExcel(coinArray []CryptoCurrencyCoin) (csvFile *excelize.File, err error) {

	csvFile = excelize.NewFile()

	if err = csvFile.SetCellValue("Sheet1", "A1", "Coin"); err != nil {
		return nil, err
	}
	if err = csvFile.SetCellValue("Sheet1", "B1", "Price"); err != nil {
		return nil, err
	}
	if err = csvFile.SetCellValue("Sheet1", "C1", "Symbol"); err != nil {
		return nil, err
	}
	if err = csvFile.SetCellValue("Sheet1", "D1", "ID"); err != nil {
		return nil, err
	}

	if err = csvFile.SetCellValue("Sheet1", "E1", "Date Created"); err != nil {
		return nil, err
	}

	for i, coin := range coinArray {

		if err = csvFile.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), coin.Name); err != nil {
			return nil, err
		}
		if err = csvFile.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), coin.Quote); err != nil {
			return nil, err
		}
		if err = csvFile.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), coin.Symbol); err != nil {
			return nil, err
		}
		if err = csvFile.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), coin.ID); err != nil {
			return nil, err
		}
		if err = csvFile.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), coin.DateAdded); err != nil {
			return nil, err
		}

	}

	err = csvFile.SaveAs("coins.xlsx")
	if err != nil {
		return nil, err
	}

	return csvFile, err
}
