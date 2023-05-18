package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	cityCode := 110000
	apiKey := "your_key" // 在高德开放平台申请的API密钥
	count := 2
	// 创建Excel文件
	file := excelize.NewFile()
	
	// 在Sheet1中写入天气信息
	file.SetCellValue("Sheet1", "A1", "省份")
	file.SetCellValue("Sheet1", "B1", "城市")
	file.SetCellValue("Sheet1", "C1", "天气")
	file.SetCellValue("Sheet1", "D1", "温度")
	file.SetCellValue("Sheet1", "E1", "风向")
	file.SetCellValue("Sheet1", "F1", "风力")
	file.SetCellValue("Sheet1", "G1", "湿度")
	file.SetCellValue("Sheet1", "H1", "报告时间")
	flag := 1
	for cityCode <= 820000 && cityCode >= 100000 && count < 200 {
		for flag == 1 {
			url := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?city=%s&key=%s", strconv.Itoa(cityCode), apiKey)
			response, err := http.Get(url)
			if err != nil {
				fmt.Println("请求出错:", err)
				return
			}
			defer response.Body.Close()
			
			// 读取响应内容
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("读取响应出错:", err)
				return
			}
			
			// 解析JSON数据
			var data map[string]interface{}
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println("解析JSON出错:", err)
				return
			}
			
			// 提取天气信息
			weatherInfo, ok := data["lives"].([]interface{})
			if !ok || len(weatherInfo) == 0 {
				fmt.Println("未找到天气信息")
			}
			if len(weatherInfo) == 0 {
				fmt.Println("错误的城市编码")
				flag = 0
				continue
			}
			weatherData, ok := weatherInfo[0].(map[string]interface{})
			if !ok {
				fmt.Println("天气信息格式错误")
				flag = 0
				continue
			}
			province := weatherData["province"].(string)
			city := weatherData["city"].(string)
			weather := weatherData["weather"].(string)
			temperature := weatherData["temperature"].(string)
			winddirection := weatherData["winddirection"].(string)
			windpower := weatherData["windpower"].(string)
			humidity := weatherData["humidity"].(string)
			reportTime := weatherData["reporttime"].(string)
			
			file.SetCellValue("Sheet1", fmt.Sprintf("A%d", count), province)
			file.SetCellValue("Sheet1", fmt.Sprintf("B%d", count), city)
			file.SetCellValue("Sheet1", fmt.Sprintf("C%d", count), weather)
			file.SetCellValue("Sheet1", fmt.Sprintf("D%d", count), temperature)
			file.SetCellValue("Sheet1", fmt.Sprintf("E%d", count), winddirection)
			file.SetCellValue("Sheet1", fmt.Sprintf("F%d", count), windpower)
			file.SetCellValue("Sheet1", fmt.Sprintf("G%d", count), humidity)
			file.SetCellValue("Sheet1", fmt.Sprintf("H%d", count), reportTime)
			
			flag = 1
			count++
			cityCode += 100
		}
		cityCode = (cityCode/10000 + 1) * 10000
		flag = 1
	}
	// 保存文件
	err := file.SaveAs("weather.xlsx")
	if err != nil {
		fmt.Println("保存文件出错:", err)
		return
	}
	
	fmt.Println("天气信息已导入到weather.xlsx文件")
}
