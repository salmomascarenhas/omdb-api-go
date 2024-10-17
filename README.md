# OMDB API Wrapper in Go

This is a simple API wrapper built in Go that interacts with the OMDB API to search for movies. It uses the [OMDB API](https://www.omdbapi.com/) to retrieve movie data by title. The API uses `chi` for routing and includes middleware for logging and error handling.

## Features

- Search for movies by title.
- Returns JSON-formatted responses.
- Includes error handling and logging.

## Prerequisites

- Go 1.18+
- OMDB API Key (Sign up at [OMDB](https://www.omdbapi.com/apikey.aspx) to get your API key)

## Installation

1. Clone the repository:

   ```bash
   git clone git@github.com:salmomascarenhas/omdb-api-go.git
   cd omdb-api-go
   ```

2. Install the dependencies:

   ```bash
   go mod tidy
   ```

3. Set the OMDB API Key as an environment variable:

   ```bash
   export OMDB_API_KEY=your_omdb_api_key
   ```

## Usage

To start the server, run:

```bash
go run main.go
```

The server will start at `http://localhost:3000`.

### API Endpoints

#### `GET /`

Search for a movie by its title.

- **Query Parameters:**
  - `s`: The search term for the movie title.

- **Example Request:**

  ```bash
  curl "http://localhost:3000/?s=Inception"
  ```

- **Example Response:**

  ```json
  {
    "data": {
      "Search": [
        {
          "Title": "Inception",
          "Year": "2010",
          "ImdbID": "tt1375666",
          "Type": "movie",
          "Poster": "https://..."
        },
        ...
      ],
      "TotalResults": "123",
      "Response": "True"
    }
  }
  ```

- **Error Response:**

  ```json
  {
    "error": "something wrong with omdb api"
  }
  ```

## Project Structure

- `omdb/`: Contains the logic for interacting with the OMDB API.
- `api/`: Contains the HTTP handler and routing.
- `main.go`: Entry point for the server.

## Logging

This project uses `slog` for logging errors and information. Logs will include request IDs and any encountered errors.

## Middleware

- `Recoverer`: Recovers from panics and logs them.
- `RequestID`: Generates a unique request ID for each incoming request.
- `Logger`: Logs details about each incoming request.

## Contributing

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Make your changes.
4. Commit your changes: `git commit -m 'Add some feature'`.
5. Push to the branch: `git push origin feature-name`.
6. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.