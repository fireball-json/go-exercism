// Package weather 天气预测.
package weather

var (
	// CurrentCondition 当前温度.
	CurrentCondition string
	// CurrentLocation 当前地区.
	CurrentLocation string
)

// Forecast 返回天气.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
