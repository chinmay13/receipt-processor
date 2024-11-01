# Receipt Processing Service

This repository contains a backend API service built in Go with the Gin framework. It processes and manages receipt data, allowing points to be calculated based on receipt content and points to be retrieved for specific receipts.

## Overview

The Receipt Processing Service provides two primary API endpoints:

POST /receipts/process: Accepts receipt details, validates them, and processes them to award points.\
GET /receipts/{id}/points: Retrieves the awarded points for a specific receipt.

## Technologies Used and Reason to Choose Them

Go: Primary language for backend development \
Gin: Lightweight, fast HTTP web framework for handling requests \
Go Validator: For validation of request data

# Running the Project

### Using Go

1. Clone the git repository
```
    git clone https://github.com/chinmay13/receipt-processor.git
    cd receipt-processor\receipt-processing
```
2. Install dependencies

```
    go mod tidy
```

3. Run the application

```
    go run main.go
```

### Using Docker

1. Build the Docker image

```
    docker build -t receipt-processing .
```

2. Run the Docker container

```
    docker run -p 8080:8080 receipt-processing
```

The API will be available at http://localhost:8080

## Design Decisions

1. Gin Framework: Selected for its high performance, minimal memory footprint, and ease of handling JSON APIs.
2. Custom Validation: Validates that fields like Total and Price are float values in string format. Custom error messages provide specific feedback for incorrect values.
3. Separation of Business Logic and HTTP Handlers: Ensures that business logic can be reused and tested independently of the HTTP framework.
4. Modular Test Structure: Consolidates all test files within a single package to keep tests organized and avoid clutter in the main application structure.

## Obervations

1. Total is given in the input JSON for receipts/process API which can be derived from individual item prices.
2. Floating point values are given as strings that need to be converted to float64 before being used in points calculation.

## Challenges

Data Validation: Ensuring correct formats for dates, times, and float values represented as strings required extensive validation logic.

## Future Scope

1. Database Integration: Add MongoDB for durable storage of receipts and points data.
2. Enhanced Error Handling: Implement structured error handling and logging for better observability and debugging.
3. Asynchronous Processing: Use Kafka to handle receipt processing asynchronously, allowing for non-blocking receipt submission.
4. Caching: Add Redis to cache frequently requested data, such as points total, for faster access.

## Making it Production-Ready

1. Security: Add JWT-based authentication for secure access to endpoints.
2. Environment Management: Use environment variables for configuration (e.g., ports, database URIs, etc.) rather than hardcoding them.
3. Monitoring: Set up monitoring and alerting (e.g., Prometheus and Grafana) to observe API health and performance.

# About Me:

My name is Chinmay Bhate, I am currently pursuing a Master’s in Computer Science at the Rochester Institute of Technology.
With three years of hands-on experience, I am an adaptable and driven professional skilled in developing innovative solutions that deliver measurable impact.

- Email: [cb4490@rit.edu](mailto:cb4490@rit.edu)
- Website: [My Portfolio](https://chinmay13.github.io/)
