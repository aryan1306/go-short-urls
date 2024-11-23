# Go URL Shortener

A simple and efficient URL shortener service built with Go, Fiber, and Redis. This service allows you to create shortened URLs that redirect to their original destinations.

## Features

- Fast URL shortening using random string generation
- Permanent redirects to original URLs
- Redis-based storage for quick access and persistence
- RESTful API endpoints
- Environment-based configuration
- Error handling and validation

## Prerequisites

Before running this application, make sure you have the following installed:

- Go 1.16 or higher
- Redis server
- Git (for cloning the repository)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/aryan1306/go-short-urls.git
cd go-short-urls
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory:
```env
REDIS_URL=redis://localhost:6379
```

## Running the Application

1. Start the Redis server:
```bash
redis-server
```

2. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8000`

## API Documentation

### Create Short URL

Creates a shortened URL from a provided original URL.

**Endpoint:** `POST /shorten`

**Request Body:**
```json
{
    "url": "https://example.com/very-long-url"
}
```

**Success Response:**
```json
{
    "status": 201,
    "message": "short url created",
    "data": {
        "url": "http://localhost:8000/abc123def4"
    }
}
```

**Error Responses:**
```json
{
    "status": 400,
    "message": "Bad Request"
}
```
```json
{
    "status": 500,
    "message": "Error generating random string"
}
```

### Access Shortened URL

Redirects to the original URL.

**Endpoint:** `GET /:url`

- Replace `:url` with the shortened URL key
- Returns a 301 permanent redirect to the original URL
- Returns 404 if the URL is not found

## Project Structure

```
go-short-urls/
├── main.go
├── .env
├── internal/
│   ├── randomString/
│   │   └── randomString.go
│   └── redisClient/
│       └── redis.go
└── README.md
```

## Dependencies

- [Fiber](https://github.com/gofiber/fiber/v2) - Web framework
- [Godotenv](https://github.com/joho/godotenv) - Environment variable loader
- [Redis](https://github.com/go-redis/redis) - Redis client for Go

## Error Handling

The application includes comprehensive error handling for:
- Invalid request bodies
- Empty URLs
- Redis connection issues
- Non-existent shortened URLs
- Random string generation failures

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Thanks to the Go community for the amazing tools and libraries
- Fiber framework for the efficient web server
- Redis for the powerful key-value store

## Contact

Aryan - [@aryandotexe](https://twitter.com/aryandotexe)
Project Link: [https://github.com/aryan1306/go-short-urls](https://github.com/aryan1306/go-short-urls)
