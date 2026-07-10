# Address Weather API

Small Go API that receives a Brazilian postal code (CEP), resolves its address, finds geographic coordinates, and returns current weather data for that location.

This project is mainly a learning project focused on backend fundamentals with Go: HTTP clients, external API integration, layered architecture, structured logging, and tests.

## Flow

```text
Postal code
  -> ViaCEP: address data
  -> Nominatim/OpenStreetMap: latitude and longitude
  -> OpenWeather: current weather
  -> API response
```

## Stack

- Go
- Gin
- slog
- godotenv
- httptest

## Project structure

```text
cmd/api
  application entrypoint

routes
  HTTP route registration

internal/handlers
  HTTP handlers and handler-facing interfaces

internal/services
  use case orchestration

internal/clients
  external API clients

internal/domain
  application response/domain models

internal/validation
  input validation helpers
```

## Endpoint

```http
GET /weather/:postalCode
```

Example:

```http
GET http://localhost:8080/weather/85035000
```

## Environment variables

Create a `.env` file in the project root:

```env
VIA_CEP_BASE_URL=https://viacep.com.br/ws/
NOMINATIM_BASE_URL=https://nominatim.openstreetmap.org/search
OPEN_WEATHER_BASE_URL=https://api.openweathermap.org/data/2.5/weather
OPEN_WEATHER_API_KEY=your_openweather_api_key
```

`OPEN_WEATHER_API_KEY` is required. ViaCEP and Nominatim do not require API keys.

## Running locally

```bash
go run ./cmd/api
```

Server starts on:

```text
http://localhost:8080
```

With Air hot reload:

```bash
air
```

## Tests

Run all tests:

```bash
go test ./...
```

Generate coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Current test coverage focus

- Handler route behavior
- Address client with fake HTTP server
- Geocoding client with fake HTTP server
- Weather client with fake HTTP server

## External APIs

- [ViaCEP](https://viacep.com.br/)
- [Nominatim](https://nominatim.org/)
- [OpenWeather](https://openweathermap.org/api)

## Notes

Nominatim public usage requires a custom `User-Agent` and should not be used heavily. For production usage, consider caching, rate limiting, or a dedicated geocoding provider.

## Roadmap

- Improve error responses
- Add service-level tests with fake clients
- Add in-memory cache with TTL
- Add Dockerfile
- Add basic rate limiting
- Add API documentation
