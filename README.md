## Caching In Order API
  - What is Caching.
    
    `
       Caching is the process of storing accessed data in a temporary storage location (cache) to improve performance and reduce retrieval time. Instead of repeatedly fetching data from a slow source (such as a database, API, or disk), caching allows quick access from memory or a faster storage layer.
    `
    
In this Order API I have a use the Caching in a `Getbill` function.

## Run Locally

  - run Project
  ```bash
      go run main.go
  ```

## In Postman Send This Links :

- # For Customer
  - For Create The Customer
    ```bash
      http://localhost:3030/customer
    ```

  - For List Of The Customer
      Set The Limit and skip then send the link.
    ```bash
      http://localhost:3030/customer
    ```

  - For Update & Delete The Customer
    ```bash
      http://localhost:3030/customer/:CustomerId
    ```

- # For Items
  - For Create The Items
    ```bash
      http://localhost:3030/items
    ```

  - For List of The Items
       Set The Limit and skip then send the link.
    ```bash
      http://localhost:3030/items
    ```

  - For Update & Delete The Items 
    ```bash
      http://localhost:3030/items/:ItemId
    ```

- # For Order
  - For Place The Order
    ```bash
      http://localhost:3030/Order
    ```

  - For Approve The Order
    ```bash
      http://localhost:3030/Order/:CustomerId/approve
    ```

  - For Cancel The Order
    ```bash
      http://localhost:3030/Order/:CustomerId/cancel
    ```

  - For Bill The Order
      In this `Getbill` Function i will use the caching. After send this link click a 2 time `Send` button caching is allow.
    ```bash
      http://localhost:3030/Order/:CustomerId/bill
    ```

    
