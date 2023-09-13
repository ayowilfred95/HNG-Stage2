# API Documentation

This document provides details on how to set up, run, and use the Person API. The API allows you to perform CRUD operations on person records stored in a MongoDB database.


## API Endpoints

### Create Person

**Endpoint:** `/api`

- **Method:** POST
- **Request Format:** JSON
  - Example Request Body:
    ```json
    {
      "name": "John Doe",
      "hobby": "Reading"
    }
    ```
- **Response Format:** JSON
  - Example Response:
    ```json
    {
      "name": "John Doe",
      "hobby": "Reading",
      "ID": "65015b47ba98afa16ae01399"
    }
    ```

### Get Persons

**Endpoint:** `/api`

- **Method:** GET
- **Response Format:** JSON
  - Example Response:
    ```json
    {
      "data": [
        {
          "ID": "65015b47ba98afa16ae01399",
          "name": "John Doe",
          "hobby": "Reading"
        },
        {
          "ID": "65015b47ba98afa16ae0139a",
          "name": "Jane Smith",
          "hobby": "Gardening"
        }
      ]
    }
    ```

### Get Person

**Endpoint:** `/api/:userId`

- **Method:** GET
- **Response Format:** JSON
  - Example Response:
    ```json
    {
      "data": {
        "ID": "65015b47ba98afa16ae01399",
        "name": "John Doe",
        "hobby": "Reading"
      }
    }
    ```

### Update Person

**Endpoint:** `/api/person/:userId`

- **Method:** PUT
- **Request Format:** JSON
  - Example Request Body:
    ```json
    {
      "name": "Updated Name",
      "hobby": "Updated Hobby"
    }
    ```
- **Response Format:** JSON
  - Example Response:
    ```json
    {
      "name": "Updated Name",
      "hobby": "Updated Hobby",
      "result": {
        "MatchedCount": 1,
        "ModifiedCount": 1
      }
    }
    ```

### Delete Person

**Endpoint:** `/api/person/:userId`

- **Method:** DELETE
- **Response Format:** JSON
  - Example Response:
    ```json
    {
      "result": {
        "DeletedCount": 1
      },
      "deleted_person": {
        "name": "John Doe"
      }
    }
    ```

## Request and Response Formats

- The API accepts and returns data in JSON format. 


## Sample Usage

### Create a New Person

```shell
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "hobby": "Reading"
}' http://localhost:8000/api/person


## Known Limitations
- The API does not perform input validation on name and hobby fields.
- Numeric IDs for persons are generated from MongoDB InsertedID which is a set of strings and numbers.

## Instructions for setting up and deploying the API locally or on a server.
### Locally
- Install postman extension on Vscode 
- Start the server by running either <go run main.go> or <CompileDaemon -command="go run main.go"> in your terminal
- Add a new request , starting from POST  and enter the url <localhost:8000/api> 
 
 ## Link to Postman Collection
