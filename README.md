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

Create your local environment file from the provided example:

```bash
cp .env.example .env
```

On PowerShell:

```powershell
Copy-Item .env.example .env
```

Then set your OpenWeather API key in `.env`:

```env
OPEN_WEATHER_API_KEY=your_openweather_api_key
```

| Variable | Required | Description |
| --- | --- | --- |
| `VIA_CEP_BASE_URL` | Yes | Base URL used to retrieve address information from ViaCEP |
| `NOMINATIM_BASE_URL` | Yes | Base URL used to geocode addresses with Nominatim |
| `OPEN_WEATHER_BASE_URL` | Yes | Base URL used to retrieve current weather data |
| `OPEN_WEATHER_API_KEY` | Yes | API key used to authenticate with OpenWeather |

The `.env.example` file contains safe example values and is versioned with the project. The `.env` file contains local configuration and secrets and must not be committed.

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

## Running with Docker

Build the image:

```bash
docker build -t address-weather-api .
```

Run the container using your local environment file:

```bash
docker run --env-file .env -p 8080:8080 address-weather-api
```

The API will be available at:

```text
http://localhost:8080
```

The `--env-file` option reads `.env` from the host and injects its values as environment variables. The file is not copied into the container image.

Test the weather endpoint:

```bash
curl http://localhost:8080/weather/85035000
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
- Add basic rate limiting
- Add API documentation
