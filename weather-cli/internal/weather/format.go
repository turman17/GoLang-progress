package weather

import (
	"fmt"
	"strings"
)

func FormatCurrent(cw CurrentWeather, units string) string {
	unitTemp := "°C"
	unitSpeed := "m/s"
	if units == "imperial" {
		unitTemp = "°F"
		unitSpeed = "mph"
	}

	cond := ""
	if len(cw.Weather) > 0 {
		cond = fmt.Sprintf("%s (%s)", cw.Weather[0].Main, cw.Weather[0].Description)
	}

	location := cw.Name
	if cw.Sys.Country != "" {
		location = fmt.Sprintf("%s, %s", cw.Name, cw.Sys.Country)
	}

	lines := []string{
		fmt.Sprintf("Location:   %s", location),
		fmt.Sprintf("Condition:  %s", cond),
		fmt.Sprintf("Temp:       %.1f%s (feels %.1f%s)", cw.Main.Temp, unitTemp, cw.Main.FeelsLike, unitTemp),
		fmt.Sprintf("Humidity:   %d%%", cw.Main.Humidity),
		fmt.Sprintf("Pressure:   %d hPa", cw.Main.Pressure),
		fmt.Sprintf("Wind:       %.1f %s, %d°", cw.Wind.Speed, unitSpeed, cw.Wind.Deg),
	}
	return strings.Join(lines, "\n")
}