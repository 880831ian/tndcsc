package services

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"tndcsc/models"
)

func FetchData() (models.Data, error) {
	resp, err := http.Get("http://tndcsc.com.tw/swimPool.aspx")
	if err != nil {
		return models.Data{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return models.Data{}, err
	}

	// 提取游泳池和健身中心人數
	swimCount := doc.Find(".w3_agile_logo font").First().Text()
	fitnessCount := doc.Find(".w3_agile_logo font").Eq(1).Text()

	// 設定人數上限
	swimPoolLimit := 480
	fitnessCenterLimit := 180

	data := models.Data{
		SwimCount:          swimCount,
		FitnessCount:       fitnessCount,
		SwimPoolLimit:      swimPoolLimit,
		FitnessCenterLimit: fitnessCenterLimit,
	}

	// 計算游泳池和健身中心的百分比
	swimCountInt, _ := strconv.Atoi(swimCount)
	fitnessCountInt, _ := strconv.Atoi(fitnessCount)

	if swimPoolLimit > 0 {
		data.SwimPercentage = fmt.Sprintf("%.1f%%", float64(swimCountInt)/float64(swimPoolLimit)*100)
	}

	if fitnessCenterLimit > 0 {
		data.FitnessPercentage = fmt.Sprintf("%.1f%%", float64(fitnessCountInt)/float64(fitnessCenterLimit)*100)
	}

	return data, nil
}
