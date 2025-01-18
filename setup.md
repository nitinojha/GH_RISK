# Risk Management Application

This project is a RESTful service for managing risks, implemented in Go with in-memory storage. The application exposes APIs to create, retrieve, and list risks.

## Features
- REST API endpoints:
  - `GET /v1/risks`: Retrieve all risks.
  - `POST /v1/risks`: Create a new risk.
  - `GET /v1/risks/{id}`: Retrieve a specific risk by ID.
- In-memory storage for simplicity.
- Auto-generated UUIDs for risk identification.
- JSON-based data transfer.
- HTTP response codes for standard communication.

## Prerequisites
- Go 1.18 or higher installed on your machine.

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone https://github.com/nitinojha/GH_RISK.git
   cd GH_RISK
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the Application**
   ```bash
   go run main.go
   ```

   The server will start at `http://localhost:8080`.

## API Endpoints

### 1. `GET /v1/risks`
Retrieve all risks.
- **Response**:
  ```json
  [
    {
      "id": "uuid",
      "state": "open",
      "title": "Risk title",
      "description": "Risk description"
    }
  ]
  ```

### 2. `POST /v1/risks`
Create a new risk.
- **Request Body**:
  ```json
  {
    "state": "open",
    "title": "Risk title",
    "description": "Risk description"
  }
  ```
- **Response**:
  ```json
  {
    "id": "uuid",
    "state": "open",
    "title": "Risk title",
    "description": "Risk description"
  }
  ```

### 3. `GET /v1/risks/{id}`
Retrieve a risk by its ID.
- **Response**:
  ```json
  {
    "id": "uuid",
    "state": "open",
    "title": "Risk title",
    "description": "Risk description"
  }
  ```

## Running Tests
To run tests for the application:
```bash
go test ./... -v
```

## Test-Driven Development (TDD)
The application includes unit tests to ensure the correctness of API endpoints and business logic. The following tests are covered:

1. **Test for `GET /v1/risks`:**
   - Ensures all risks are returned.
   - Validates the structure of the response.

2. **Test for `POST /v1/risks`:**
   - Checks that valid risks are created successfully.
   - Ensures invalid requests return proper error messages.

3. **Test for `GET /v1/risks/{id}`:**
   - Verifies that existing risks are fetched correctly by ID.
   - Ensures non-existent IDs return a 404 error.

Run the tests with:
```bash
go test ./... -v
```

## Step-by-Step Guide to Verify APIs Locally

1. **Start the server:**
   ```bash
   go run main.go
   ```
   The server will be available at `http://localhost:8080`.

2. **Verify `GET /v1/risks`:**
   - Use a tool like `curl` or Postman to send a GET request:
     ```bash
     curl -X GET http://localhost:8080/v1/risks
     ```
   - The response will be an empty array initially:
     ```json
     []
     ```

3. **Verify `POST /v1/risks`:**
   - Send a POST request with a JSON body:
     ```bash
     curl -X POST http://localhost:8080/v1/risks \
     -H "Content-Type: application/json" \
     -d '{"state": "open", "title": "Risk 1", "description": "First risk"}'
     ```
   - The response will include the created risk with a UUID:
     ```json
     {
       "id": "some-uuid",
       "state": "open",
       "title": "Risk 1",
       "description": "First risk"
     }
     ```

4. **Verify `GET /v1/risks/{id}`:**
   - Use the ID from the previous step to fetch the specific risk:
     ```bash
     curl -X GET http://localhost:8080/v1/risks/some-uuid
     ```
   - The response will include the risk details:
     ```json
     {
       "id": "some-uuid",
       "state": "open",
       "title": "Risk 1",
       "description": "First risk"
     }
     ```

## Assumptions
- No authentication or authorization is implemented as per the requirements.
- Limited `state` values: `[open, closed, accepted, investigating]`.
- Data is stored in-memory; it will be lost when the server restarts.

## Notes for Reviewers
- The project uses the `gorilla/mux` router for handling HTTP routes.
- The `uuid` package is used to generate unique IDs for risks.
- All risk validations are handled in the `models.Risk` struct.

## References
- [Gorilla Mux](https://github.com/gorilla/mux)
- [UUID Package](https://pkg.go.dev/github.com/google/uuid)
