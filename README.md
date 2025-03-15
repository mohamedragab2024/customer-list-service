# Customer List Service

## Overview
The Customer List Service is an API that allows you to manage a list of customers. You can add, update, delete, and retrieve customer information.

## Prerequisites
- Go 1.x
- Docker
- Docker Compose

## Running the API

To run the API, you can use the provided `Makefile`. This will build and start the service using Docker.

1. **Build and start the service:**
    ```sh
    make up
    ```

2. **Stop the service:**
    ```sh
    make down
    ```

3. **Rebuild the service:**
    ```sh
    make rebuild
    ```

## API Endpoints

### Get All Customers
- **URL:** `/customers`
- **Method:** `GET`
- **Description:** Retrieves a list of all customers.

### Get Customer by ID
- **URL:** `/customers/{id}`
- **Method:** `GET`
- **Description:** Retrieves a customer by their ID.

### Add a New Customer
- **URL:** `/customers`
- **Method:** `POST`
- **Description:** Adds a new customer.
- **Body:**
    ```json
    {
        "name": "Customer Name",
        "email": "customer@example.com"
    }
    ```

### Update a Customer
- **URL:** `/customers/{id}`
- **Method:** `PUT`
- **Description:** Updates an existing customer.
- **Body:**
    ```json
    {
        "name": "Updated Name",
        "email": "updated@example.com"
    }
    ```

### Delete a Customer
- **URL:** `/customers/{id}`
- **Method:** `DELETE`
- **Description:** Deletes a customer by their ID.

## Development

To run the service locally for development purposes, you can use the following commands:

1. **Run the application:**
    ```sh
    make run
    ```

2. **Build the application:**
    ```sh
    make build
    ```

3. **Run tests:**
    ```sh
    make test
    ```

4. **Initialize/Reset Database:**
    ```sh
    make db
    ```

## Build and Deploy

1. **Build Docker image:**
    ```sh
    make docker-build
    ```

2. **Run with Docker Compose:**
    ```sh
    make up
    ```

## License
This project is licensed under the MIT License.
