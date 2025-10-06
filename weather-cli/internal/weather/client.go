package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetWeatherByCity calls OpenWeather's /data/2.5/weather?q=<city>&units=<units>
func (c Client) GetWeatherByCity(city, units string) (CurrentWeather, error) {
	if c.BaseURL == "" {
		c.BaseURL = "https://api.openweathermap.org/data/2.5"
	}
	if units == "" {
		units = "metric"
	}

	u, err := url.Parse(c.BaseURL + "/weather")
	if err != nil {
		return CurrentWeather{}, err
	}
	q := u.Query()
	q.Set("q", city)
	q.Set("units", units)
	q.Set("appid", c.APIKey)
	u.RawQuery = q.Encode()

	httpClient := &http.Client{Timeout: c.Timeout}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return CurrentWeather{}, err
	}
	req.Header.Set("User-Agent", "weather-cli (+https://github.com/turman17)")

	resp, err := httpClient.Do(req)
	if err != nil {
		return CurrentWeather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return CurrentWeather{}, fmt.Errorf("api status %d", resp.StatusCode)
	}

	var cw CurrentWeather
	if err := json.NewDecoder(resp.Body).Decode(&cw); err != nil {
		return CurrentWeather{}, err
	}
	return cw, nil
}