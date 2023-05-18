package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Alarm struct {
	Type    string `json:"alarm_type"`
	Level   string `json:"alarm_level"`
	Content string `json:"alarm_content"`
}

type AQI struct {
	Air       string `json:"air"`
	Level     string `json:"air_level"`
	Tips      string `json:"air_tips"`
	PM25      string `json:"pm25"`
	PM25Desc  string `json:"pm25_desc"`
	PM10      string `json:"pm10"`
	PM10Desc  string `json:"pm10_desc"`
	O3        string `json:"o3"`
	O3Desc    string `json:"o3_desc"`
	NO2       string `json:"no2"`
	NO2Desc   string `json:"no2_desc"`
	SO2       string `json:"so2"`
	SO2Desc   string `json:"so2_desc"`
	CO        string `json:"co"`
	CODesc    string `json:"co_desc"`
	Kouzhao   string `json:"kouzhao"`
	Waichu    string `json:"waichu"`
	Yundong   string `json:"yundong"`
	Kaichuang string `json:"kaichuang"`
	Jinghuaqi string `json:"jinghuaqi"`
	CityID    string `json:"cityid"`
	City      string `json:"city"`
	CityEn    string `json:"cityEn"`
	Country   string `json:"country"`
	CountryEn string `json:"countryEn"`
}

type WeatherInfo struct {
	City         string `json:"city"`
	CityID       string `json:"cityid"`
	Date         string `json:"date"`
	Week         string `json:"week"`
	UpdateTime   string `json:"update_time"`
	CityEn       string `json:"cityEn"`
	Country      string `json:"country"`
	CountryEn    string `json:"countryEn"`
	Weather      string `json:"wea"`
	WeatherImg   string `json:"wea_img"`
	Temperature  string `json:"tem"`
	Temperature1 string `json:"tem1"`
	Temperature2 string `json:"tem2"`
	Wind         string `json:"win"`
	WindSpeed    string `json:"win_speed"`
	WindMeter    string `json:"win_meter"`
	Humidity     string `json:"humidity"`
	Visibility   string `json:"visibility"`
	Pressure     string `json:"pressure"`
	Air          string `json:"air"`
	AirPM25      string `json:"air_pm25"`
	AirLevel     string `json:"air_level"`
	AirTips      string `json:"air_tips"`
	Rain         string `json:"rain_pcpn"`
	Alarm        Alarm  `json:"alarm"`
	AQI          AQI    `json:"aqi"`
}

func main() {
	
	url := "https://v0.yiketianqi.com/api?unescape=1&version=v61&appid=39877347&appsecret=rBf7o5bU&cityid=101271501"
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
	var data WeatherInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解析JSON出错:", err)
		return
	}
	
	// 提取天气信息
	fmt.Println("城市编号:", data.CityID)
	fmt.Println("日期:", data.Date)
	fmt.Println("星期:", data.Week)
	fmt.Println("更新时间:", data.UpdateTime)
	fmt.Println("城市名称:", data.City)
	fmt.Println("国家:", data.Country)
	fmt.Println("实时天气情况:", data.Weather)
	fmt.Println("实时温度:", data.Temperature)
	fmt.Println("最高温度:", data.Temperature1)
	fmt.Println("最低温度:", data.Temperature2)
	fmt.Println("风向:", data.Wind)
	fmt.Println("风力等级:", data.WindSpeed)
	fmt.Println("风速:", data.WindMeter)
	fmt.Println("湿度:", data.Humidity)
	fmt.Println("能见度:", data.Visibility)
	fmt.Println("气压:", data.Pressure)
	fmt.Println("空气质量:", data.Air)
	fmt.Println("PM2.5:", data.AirPM25)
	fmt.Println("空气质量等级:", data.AirLevel)
	fmt.Println("空气质量提示:", data.AirTips)
	fmt.Println("降雨量:", data.Rain)
	fmt.Println("预警类型:", data.Alarm.Type)
	fmt.Println("预警等级:", data.Alarm.Level)
	fmt.Println("预警内容:", data.Alarm.Content)
	fmt.Println("AQI空气质量:", data.AQI.Air)
	fmt.Println("AQI空气质量等级:", data.AQI.Level)
	fmt.Println("AQI空气质量提示:", data.AQI.Tips)
	fmt.Println("PM2.5指数:", data.AQI.PM25)
	fmt.Println("PM2.5指数等级:", data.AQI.PM25Desc)
	fmt.Println("PM10指数:", data.AQI.PM10)
	fmt.Println("PM10指数等级:", data.AQI.PM10Desc)
	fmt.Println("O3指数:", data.AQI.O3)
	fmt.Println("O3指数等级:", data.AQI.O3Desc)
	fmt.Println("NO2指数:", data.AQI.NO2)
	fmt.Println("NO2指数等级:", data.AQI.NO2Desc)
	fmt.Println("SO2指数:", data.AQI.SO2)
	fmt.Println("SO2指数等级:", data.AQI.SO2Desc)
	fmt.Println("CO指数:", data.AQI.CO)
	fmt.Println("CO指数等级:", data.AQI.CODesc)
	fmt.Println("口罩建议:", data.AQI.Kouzhao)
	fmt.Println("外出建议:", data.AQI.Waichu)
	fmt.Println("运动建议:", data.AQI.Yundong)
	fmt.Println("开窗建议:", data.AQI.Kaichuang)
	fmt.Println("净化器建议:", data.AQI.Jinghuaqi)
}
