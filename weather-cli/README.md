

# Weather CLI Tool

A simple command-line tool written in Go to fetch and display current weather information for a specified city using the OpenWeather API.

## About

This project was created as a learning exercise to explore Go's core features, including working with HTTP requests, decoding JSON, handling environment variables, and using structs. It's a practical example for anyone looking to understand how Go can be used to interact with external APIs and build CLI applications.

## Features

- Fetches current weather data for a given city from the OpenWeather API.
- Supports command-line flags for city and units (metric/imperial).
- Loads API key from environment variables.

## Project Structure

```
weather-cli/
  ├── cmd/                # Main entry point and CLI flag parsing
  ├── internal/
  │     └── weather/      # Weather API logic (fetching, parsing, etc.)
  ├── .env                # Environment variables (API key)
  ├── go.mod
  ├── go.sum
  └── README.md
```

## Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/weather-cli.git
   cd weather-cli
   ```

2. **Get an OpenWeather API key**
   - Sign up at [OpenWeather](https://openweathermap.org/api) and generate an API key.

3. **Create a `.env` file in the project root:**
   ```
   OPENWEATHER_API_KEY=your_api_key_here
   ```

4. **Install dependencies**
   - This project uses [`godotenv`](https://github.com/joho/godotenv) to load environment variables.
   ```bash
   go get github.com/joho/godotenv
   ```

5. **Build the project**
   ```bash
   go build -o weather-cli ./cmd
   ```

## Usage

Run the CLI tool with various flags:

```bash
# Get weather for London in metric units (default)
./weather-cli -city London

# Get weather for New York in imperial units
./weather-cli -city "New York" -units imperial
```

**Flags:**
- `-city` (required): Name of the city (e.g., "London", "New York").
- `-units` (optional): Units for temperature. Options: `metric` (Celsius), `imperial` (Fahrenheit). Default: `metric`.

## Learning Goals

- Making HTTP requests in Go.
- Decoding JSON responses into Go structs.
- Organizing code with packages and internal modules.
- Loading environment variables using godotenv.
- Parsing command-line flags.

## Next Steps / Improvements

- Add unit tests for the weather fetching logic.
- Improve error handling and user feedback.
- Support more features from the OpenWeather API (e.g., forecasts, weather by coordinates).
- Add support for configuration files.
- Package as a distributable binary for different platforms.

---

Happy coding!