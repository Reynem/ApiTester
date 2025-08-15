# API Tester

**API Tester** is a service for testing external APIs and storing a history of requests and their results.
It allows you to:

  * Send requests to specified APIs with any method (`GET`, `POST`, `PUT`, `DELETE`)
  * Pass parameters, headers, and request bodies
  * Receive and save API responses (only if in JSON format)
  * View and update saved tests
  * Delete tests

-----

## 📦 Installation and Startup

### 1\. Clone the repository

```bash
git clone https://github.com/username/apitester.git
cd apitester
```

### 2\. Move to backend directory
```bash
cd backend
```

### 3\. Install dependencies

```bash
go mod tidy
```

### 4\. Start the server

```bash
go run main.go
```

By default, the server will be available at:

```
http://localhost:8080
```

### Frontend on Vue is currently in beta state. I do not reccomend you to use it

-----

## 📡 API Endpoints

All methods are under the `/api` prefix.

### 1\. Create a test

**POST** `/api/tests`

**Request Body (`application/json`):**

```
{
  "name": "Get Users",
  "api_endpoint": "https://jsonplaceholder.typicode.com/users",
  "method": "GET",
  "parameters": {
    "id": "1"
  },
  "headers": {
    "Authorization": "Bearer token"
  },
  "body": {}
}
```

**Response (201):**

```
{
  "name": "Get Users",
  "api_endpoint": "https://jsonplaceholder.typicode.com/users?id=1",
  "response": [ { "id": 1, "name": "Leanne Graham" } ],
  "status_code": 200,
  "created_at": "2025-08-15T10:00:00Z"
}
```

-----

### 2\. Get a test by ID

**GET** `/api/tests/{id}`

**Example:**

```
GET /api/tests/1
```

**Response (200):**

```
{
  "name": "Get Users",
  "api_endpoint": "https://jsonplaceholder.typicode.com/users?id=1",
  "response": [ { "id": 1, "name": "Leanne Graham" } ],
  "status_code": 200,
  "created_at": "2025-08-15T10:00:00Z"
}
```

-----

### 3\. Get all tests

**GET** `/api/tests`

**Response (200):**

```
[
  {
    "name": "Get Users",
    "api_endpoint": "https://jsonplaceholder.typicode.com/users?id=1",
    "response": [ { "id": 1, "name": "Leanne Graham" } ],
    "status_code": 200,
    "created_at": "2025-08-15T10:00:00Z"
  }
]
```

-----

### 4\. Update a test

**PUT** `/api/tests/{id}`

**Request Body:**

```
{
  "name": "Get Posts",
  "api_endpoint": "https://jsonplaceholder.typicode.com/posts",
  "method": "GET",
  "parameters": {},
  "headers": {},
  "body": {}
}
```

**Response (200):**

```
{
  "name": "Get Posts",
  "api_endpoint": "https://jsonplaceholder.typicode.com/posts",
  "response": [ { "id": 1, "title": "Post title" } ],
  "status_code": 200,
  "created_at": "2025-08-15T10:05:00Z"
}
```

-----

### 5\. Delete a test

**DELETE** `/api/tests/{id}`

**Example:**

```
DELETE /api/tests/1
```

**Response (200):**

```
{
  "result": "Test was deleted successfully"
}
```

-----

## ⚙️ Limitations

  * The API response must be **valid JSON**; otherwise, the request will return an error.
  * The response size is limited to **1 MB**.
  * Only the following HTTP methods are supported:
      \* `GET`
      \* `POST`
      \* `PUT`
      \* `DELETE`
  * The request body is not used for `GET` requests.

-----

## 🛠 Technology Stack

  * **Go (Echo)** — HTTP server
  * **GORM + SQLite** — Database
  * **datatypes.JSON** — For storing JSON fields in the database
  * **CORS** — For supporting frontend requests
  * **Vue** - For frontend

-----

## 📜 License

MIT License
