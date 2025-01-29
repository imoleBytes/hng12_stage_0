# hng12_stage_0

# Public API for Basic Information

## Description

This project is a simple public API built using Go and the Gin framework. It provides basic information such as a registered email address, the current datetime in ISO 8601 format, and the GitHub URL of the project's codebase. The API is deployed to a publicly accessible endpoint and follows best practices for CORS handling.

## Features

- Returns basic information in JSON format.
- Provides the current datetime as an ISO 8601 formatted timestamp.
- Supports CORS for cross-origin requests.
- Fully open-source and hosted on GitHub.

---

## Technology Stack

- **Programming Language:** Go
- **Web Framework:** Gin
- **Version Control:** GitHub
- **Deployment:** Publicly accessible endpoint
- **CORS Handling:** Enabled for cross-origin requests

---

## Setup Instructions

### Prerequisites

Ensure you have the following installed on your machine:

- Go (1.18 or later)
- Git

### Clone the Repository

```sh
git clone https://github.com/imoleBytes/hng12_stage_0.git
cd hng12_stage_0
```

### Install Dependencies

```sh
go mod tidy
```

### Run the Application Locally

```sh
go run .
```

The API will start running at `http://localhost:8080` (unless configured otherwise).

---

## API Documentation

### Endpoint

#### GET /

Returns basic information in JSON format.

#### Request

- **Method:** GET
- **URL:** `https://hng12-stage-0-basic-api.onrender.com`

#### Response (200 OK)

```json
{
  "email": "your-email@example.com",
  "current_datetime": "2025-01-30T09:30:00Z",
  "github_url": "https://github.com/imoleBytes/hng12_stage_0"
}
```

#### Example Usage

Using `curl`:

```sh
curl -X GET https://hng12-stage-0-basic-api.onrender.com/
```

Using JavaScript (fetch API):

```javascript
fetch("https://hng12-stage-0-basic-api.onrender.com/")
  .then((response) => response.json())
  .then((data) => console.log(data));
```

---

## Deployment

The API is deployed to a publicly accessible endpoint. Ensure that it remains available and properly configured for production use.

---

## Contributors

- **Imole M. Kolawole** – [GitHub Profile](https://github.com/imoleBytes)

For more Go developers, visit: [Hire Go Developers](https://hng.tech/hire/golang-developers)
