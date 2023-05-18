package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherData struct {
	Day   WeatherDetail     `json:"day"`
	Hours []HourlyCondition `json:"hours"`
	Month []WeatherDay      `json:"month"`
}

type WeatherDetail struct {
	Icon            string `json:"icon"`
	FeelsLike       string `json:"feelsLike"`
	Temperature     string `json:"temperature"`
	TempMaxSince7am string `json:"temperatureMaxSince7am"`
	Phrase          string `json:"phrase"`
	PhraseImg       string `json:"phrase_img"`
	Altimeter       string `json:"altimeter"`
	BarometerTrend  string `json:"barometerTrend"`
	Humidity        string `json:"humidity"`
	DewPoint        string `json:"dewPoint"`
	Visibility      string `json:"visibility"`
	WindSpeed       string `json:"windSpeed"`
	WindDirCompass  string `json:"windDirCompass"`
	WindDirDegrees  string `json:"windDirDegrees"`
	UVIndex         string `json:"uvIndex"`
	UVDescription   string `json:"uvDescription"`
	// Alarm           []Alarm `json:"alarm"`
	Air      string `json:"air"`
	AirLevel string `json:"air_level"`
}

type HourlyCondition struct {
	Time        string  `json:"time"`
	TimeUtc     int     `json:"timeUtc"`
	Weather     string  `json:"wea"`
	WeatherImg  string  `json:"wea_img"`
	Icon        int     `json:"icon"`
	Temperature int     `json:"tem"`
	FeelsLike   int     `json:"temfeels"`
	UVIndex     int     `json:"uvIndex"`
	Pressure    float64 `json:"pressure"`
	Humidity    int     `json:"humidity"`
	DayOrNight  string  `json:"dayOrNight"`
	CloudCover  int     `json:"cloudCover"`
	DayOfWeek   string  `json:"dayOfWeek"`
	Visibility  int     `json:"visibility"`
	Wind        string  `json:"wind"`
	WindSpeed   int     `json:"windSpeed"`
	PrecipPct   int     `json:"precipPct"`
	Air         string  `json:"air"`
}

type WeatherDay struct {
	Date       string        `json:"date"`
	DateOfWeek string        `json:"dateOfWeek"`
	Sunrise    string        `json:"sunrise"`
	Sunset     string        `json:"sunset"`
	Moonrise   string        `json:"moonrise"`
	Moonset    string        `json:"moonset"`
	MoonIcon   string        `json:"moonIcon"`
	MoonPhrase string        `json:"moonPhrase"`
	Day        WeatherDetail `json:"day"`
	Night      WeatherDetail `json:"night"`
	Air        string        `json:"air"`
	RainPcpn   string        `json:"rain_pcpn"`
}

func main() {
	url := "https://v0.yiketianqi.com/api/worldchina?appid=39877347&appsecret=rBf7o5bU&city=眉山"
	
	// 发送请求并获取天气信息
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
	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("解析JSON出错:", err)
		return
	}
	
	// 输出小时实况信息
	fmt.Println("小时实况信息:")
	for _, hour := range weatherData.Hours {
		fmt.Printf("时间: %s\n", hour.Time)
		fmt.Printf("天气状况: %s\n", hour.Weather)
		fmt.Printf("温度: %d°C\n", hour.Temperature)
		fmt.Printf("体感温度: %d°C\n", hour.FeelsLike)
		fmt.Printf("湿度: %d%%\n", hour.Humidity)
		
		fmt.Println("--------------------")
	}
	
	// 输出天气预报信息
	fmt.Println("天气预报信息:")
	for _, day := range weatherData.Month {
		fmt.Printf("日期: %s\n", day.Date)
		fmt.Printf("日出时间: %s\n", day.Sunrise)
		fmt.Printf("日落时间: %s\n", day.Sunset)
		fmt.Printf("降水量(mm): %s\n", day.RainPcpn)
		fmt.Printf("白天天气状况: %s\n", day.Day.Phrase)
		fmt.Printf("夜间天气状况: %s\n", day.Night.Phrase)
		fmt.Printf("白天温度: %s°C\n", day.Day.Temperature)
		fmt.Printf("夜间温度: %s°C\n", day.Night.Temperature)
		
		// fmt.Printf("降雨几率: %s%%\n", day.Day.PrecipPct)
		fmt.Println("--------------------")
	}
	
}
