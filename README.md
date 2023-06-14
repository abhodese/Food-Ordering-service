
## Getting Started

# E-Commerce System

This is a simple e-commerce system implemented in Golang. It consists of two services: Seller Service and Buyer Service. The Seller Service allows sellers to add and manage their products, while the Buyer Service enables buyers to search for products and place orders.

## Features

- Seller Service:
  - Add products with attributes like price, quantity, etc.
  - Update product details.
  - Retrieve product information.

- Buyer Service:
  - Search for products based on name or other criteria.
  - Place orders for available products.
  - Retrieve order information.

## Prerequisites

- Golang 1.17 or higher
- Docker
- Docker Compose

## Installation and Setup

1. Clone the repository:

   ```shell
   git clone https://github.com/your/repository.git


2. Navigate to the project directory:

    ```shell
    cd e-commerce-system
    ```

3. Build the Docker containers:

    ```shell
    docker-compose build
    ```
4. Run the Docker containers:

    ```shell
    docker-compose up
    ```
5. The Seller Service will be available at: http://localhost:8080
   The Buyer Service will be available at: http://localhost:8081

## Usage

### Seller Service

-To add a product:
    Send a POST request to http://localhost:8080/products with the product details in the request body.
-To update a product:
    Send a PUT request to http://localhost:8080/products/{product_id} with the updated product details in the request body.
-To retrieve a product:
    Send a GET request to http://localhost:8080/product?id={product_id}.

### Buyer Service
-To search for products:
    Send a GET request to http://localhost:8081/products?name={product_name} to search for products by name.
    Send a GET request to http://localhost:8081/products?category={product_category} to search for products by category.
    Send a GET request to http://localhost:8081/products?price={product_price} to search for products by price.



