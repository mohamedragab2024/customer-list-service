# Customer List Service

## Overview
The Customer List Service is an API that allows you to manage a list of customers. You can add, update, delete, and retrieve customer information.

## Prerequisites
- Go 1.21
- Docker

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
1. **Initialize/Reset Database:**
    ```sh
    make db
    ```

2. **Run the application:**
    ```sh
    make run
    ```

2. **Build the application:**
    ```sh
    make build
    ```

4. **Run tests:**
    ```sh
    make test
    ```


## Build and Deploy

1. **Build Docker image:**
    ```sh
    make docker-build
    ```

## License
This project is licensed under the MIT License.
