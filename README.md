# Customer List Service

## Overview
The Customer List Service is an API that allows you to manage a list of customers. You can add, update, delete, and retrieve customer information.

## Prerequisites
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

1. **Install dependencies:**
    ```sh
    npm install
    ```

2. **Start the development server:**
    ```sh
    npm start
    ```

3. **Run tests:**
    ```sh
    npm test
    ```

## License
This project is licensed under the MIT License.
