## Albums API

The **Albums API** is a simple RESTful API built with the [Gin Web Framework](https://github.com/gin-gonic/gin) in Go. It allows users to view, add, and search for albums by ID.

---

## Features
- Retrieve all albums (`GET /albums`)
- Add a new album (`POST /albums`)
- Get a specific album by ID (`GET /albums/:id`)
- Welcome message on the root endpoint (`GET /`)

---

## Getting Started

### Prerequisites
1. Install [Go](https://golang.org/dl/) (at least version 1.17).
2. Install the Gin framework:
   ```sh
   go get -u github.com/gin-gonic/gin

### Running the Application

    Clone this repository or copy the code into a new Go file, e.g., main.go.
    Open a terminal and navigate to the directory containing the main.go file.
    Run the application:

    go run .

    The server will start on localhost:8080.

### API Endpoints
#### 1. Home

    URL: /
    Method: GET
    Description: Displays a welcome message.
    Response:

    Welcome to the Albums API! Use /albums to view the list of albums.

#### 2. Get All Albums

    URL: /albums
    Method: GET
    Description: Retrieves the list of all albums.
    Response Example:

    [
      {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
      },
      {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
      },
      {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
      }
    ]

#### 3. Add an Album

    URL: /albums
    Method: POST
    Description: Adds a new album to the list.
    Request Body (JSON):

{
  "id": "4",
  "title": "Kind of Blue",
  "artist": "Miles Davis",
  "price": 29.99
}

##### Response Example:

    {
      "id": "4",
      "title": "Kind of Blue",
      "artist": "Miles Davis",
      "price": 29.99
    }

#### 4. Get Album by ID

    URL: /albums/:id
    Method: GET
    Description: Retrieves a specific album by its id.
    Response Example (Success):

{
  "id": "1",
  "title": "Blue Train",
  "artist": "John Coltrane",
  "price": 56.99
}

Response Example (Not Found):

    {
      "message": "album not found"
    }

#### Testing the API

You can test the API using tools like Postman, curl, or your web browser.
Example curl Commands:

    Get all albums:

curl http://localhost:8080/albums

Add a new album:

curl -X POST -H "Content-Type: application/json" -d '{"id":"4","title":"Kind of Blue","artist":"Miles Davis","price":29.99}' http://localhost:8080/albums

Get album by ID:

    curl http://localhost:8080/albums/1

Project Structure

This project consists of:

    main.go: The main file containing the application logic.
    Endpoints:
        /: Home route.
        /albums: List and add albums.
        /albums/:id: Retrieve albums by ID.

Notes

    This API is for demonstration purposes and does not include persistent storage (albums are stored in memory).
    In production, use gin.SetMode(gin.ReleaseMode) to disable debug logs.

Happy coding! 🎵