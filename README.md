# CS Ticketing API

## Overview

This API provides endpoints for managing support tickets. It allows users to retrieve a list of tickets with filtering, sorting, and pagination capabilities. The API is built using Go, Gin framework, GORM ORM, and PostgreSQL database.

## Features
- Create a new ticket
- Retrieve ticket details
- Validate user input
- User-friendly error messages

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Clone the Repository](#clone-the-repository)
  - [Install Dependencies](#install-dependencies)
- [Docker Setup](#docker-setup)
  - [Using Docker Compose](#using-docker-compose)
- [Nodemon Setup](#nodemon-setup)
  - [Using Nodemon](#using-nodemon)
- [API Endpoints](#api-endpoints)
  - [Create Ticket](#create-ticket)
- [Testing the API](#testing-the-api)
- [Running Unit Tests](#running-unit-tests)
- [Load Testing with K6](#load-testing-with-k6)
- [Contributing](#contributing)

## Installation

### Prerequisites
- Go (version 1.23 or later)
- Git
- Docker (for Docker setup)
- Node.js and npm (for Nodemon setup)
- K6 (for load testing)


### Clone the Repository
```bash
git clone https://github.com/balqisgautama/cs-ticketing.git
cd cs-ticketing
```

### Install Dependencies
Run the following command to install the required Go packages:
```bash
go mod tidy
```

## Docker Setup

### Using Docker Compose
To deploy the CS Ticketing API using Docker, you can use the provided `Makefile`, `Dockerfile` and `docker-compose.yml`.

1. **Build and Run the Application**:
   In the root directory of your project, run:
   ```bash
   make run
   ```

2. **Migrate DB**:
   In the root directory of your project (different terminal), run:
   ```bash
   make db-migrate
   ```

3. **Access the API**:
   The API will be accessible at `http://localhost:8080`.

4. **Stop the Application**:
   To stop the application, run:
   ```bash
   make stop clean
   ```

5. **Restart the Application**:
   To stop the application, run:
   ```bash
   make rebuild
   ```

## Nodemon Setup

### Using Nodemon
To deploy the CS Ticketing API using Nodemon, you can use the provided `nodemon.json`, `Dockerfile` and `docker-compose.yml`.

1. **Installing Nodemon**:
    In the root directory of your project, run:
    ```bash
    npm install -g nodemon
    ```

2. **Build and Run the Application**:
    In the root directory of your project, run:
    ```bash
    nodemon
    ```

3. **Migrate DB**:
   In the root directory of your project (different terminal), run:
   ```bash
   make db-migrate
   ```

4. **Access the API**:
    The API will be accessible at `http://localhost:8080`.

## API Endpoints

### Create Ticket

* **Endpoint:** `/tickets`
* **Method:** `POST`
- **Request Body:**
  ```json
  {
    "ticket_title": "Software is not working",
    "ticket_msg": "<p>This is a valid ticket message with more than 100 characters. It should pass the validation checks.</p>",
    "user_id": 1
  }
  ```
- **Response:**
  ```json
  {
    "message": "Ticket created successfully",
    "ticket": {
        "id": 7,
        "title": "Software is not working",
        "msg": "<p>This is a valid ticket message with more than 100 characters. It should pass the validation checks.</p>",
        "user_id": 1,
        "status": "Open",
        "CreatedAt": "2025-02-03 19:21:40" // WIB time zone
    }
  } 
  ```
### Get Ticket List

* **Endpoint:** `/tickets`
* **Method:** `GET`
* **Request Body:**
    ```json
    {
        "filter": {
            "filter_name": "created_at",
            "filter_type": "before", // or "after" or "between"
            "filter_value": "2025-02-03", // Date string in "YYYY-MM-DD" format
            // Assuming "between" filter_value is in the format "start_date,end_date"
        },
        "sort": {
            "sort_name": "created_at", // or "user_id"
            "sort_dir": "asc" // or "desc"
        },
        "page_size": 20, // Must be a positive number. Will be adjusted to the nearest valid size (10, 20, 30, 40, or 50)
        "page": 1 // Must be a positive number
    }
    ```
* **Response Body:**
    ```json
    {
        "tickets": [
            {
                "id": 1,
                "title": "Ticket 1",
                "msg": "<p>This is a valid ticket message with more than 100 characters. It should pass the validation checks.</p>",
                "status": "Open",
                "user_id": 1,
                "created_at": "2025-02-03 19:21:40" // WIB time zone
            }
        ],
        "total_page": 2,
        "current_page": 1,
        "previous_page": null,
        "next_page": 2,
        "total_items": 30,
        "page_size": 20
    }
    ```

## Testing the API

You can use tools like Postman or cURL to test the API endpoints. Here’s an example using cURL to create a ticket:

```bash
curl -X POST http://localhost:8080/tickets \
-H "Content-Type: application/json" \
-d '{"ticket_title": "Software is not working", "ticket_msg": "<p>This is a valid ticket message with more than 100 characters. It should pass the validation checks.</p>", "user_id": 1}'
```

## Running Unit Tests

To ensure the functionality of the API, unit tests are provided. You can run the unit tests using the following command:

1. Navigate to the project directory:
   ```bash
   cd your-folder-name
   ```

2. Run the tests:
   ```bash
   go test ./tests
   ```
   or
   ```bash
   make test
   ```

This command will execute all the tests in the `tests` directory and report the results.

## Load Testing with K6

To perform load testing on the API, you can use the K6 script provided in the `load_test` folder.

1. Navigate to the `load_test` directory:
   ```bash
   cd load_test
   ```

2. Run the K6 load test using the following command:
   ```bash
   k6 run script.js
   ```

This command will execute the load test defined in `script.js`, simulating the specified number of virtual users and requests to the API.

## Contributing
Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.