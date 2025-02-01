# CS Ticketing API

## Overview

The Ticket Management API is a RESTful API built with Go (Golang) that allows users to create and track tickets. The API supports basic operations such as creating tickets and retrieving ticket details. It uses PostgreSQL as the database and follows best practices for structuring Go applications.

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
- [Contributing](#contributing)

## Installation

### Prerequisites
- Go (version 1.23 or later)
- Git
- Docker (for Docker setup)
- Node.js and npm (for Nodemon setup)


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

- **Endpoint:** `POST /tickets`
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
        "status": "Open"
    }
  } 
  ```

## Testing the API

You can use tools like Postman or cURL to test the API endpoints. Here’s an example using cURL to create a ticket:

```bash
curl -X POST http://localhost:8080/tickets \
-H "Content-Type: application/json" \
-d '{"ticket_title": "Software is not working", "ticket_msg": "<p>This is a valid ticket message with more than 100 characters. It should pass the validation checks.</p>", "user_id": 1}'
```

## Contributing
Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.