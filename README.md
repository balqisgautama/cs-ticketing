# QR Code Generator API

## Overview
coming soon

## Features
- coming soon

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

## Usage

### Run the API Server
To start the API server, run:
```bash
go run ./cmd/main.go
```
The server will start on `http://localhost:8080`.

## Docker Setup

### Using Docker Compose
To deploy the CS Ticketing API using Docker, you can use the provided `Makefile`, `Dockerfile` and `docker-compose.yml`.

1. **Build and Run the Application**:
   In the root directory of your project, run:
   ```bash
   make run
   ```

2. **Access the API**:
   The API will be accessible at `http://localhost:8080`.

3. **Stop the Application**:
   To stop the application, run:
   ```bash
   make stop clean
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

3. **Access the API**:
    The API will be accessible at `http://localhost:8080`.


## Contributing
Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.