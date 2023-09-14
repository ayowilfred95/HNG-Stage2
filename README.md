# HNG-STAGE2

## Brief description of the API.
- This is a simple REST API capable of CRUD operations on a "person" resource which include <"name"> and <"hobby"> ,interfacing with mongodb database using a go package called <mongodb driver>. 
- The API dynamically handle parameters, such as adding or retrieving a person by name, id or hobby.  
- The entire project is hosted on GitHub

## Getting Started

- First install Go on your machine (version go1.20.6)
- Install Go extension in Vscode

### Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/ayowilfred95/HNG-Stage2.git

2. cd <project-name> 

3.  run <go mod download> to download all the dependencies for this project

4. Configure your environment variable by adding .env file in the root of the project


4. run <CompileDaemon -command="go run main.go"> in your terminal to get the project running
    - compileDaemon is just like nodemon in express (helps to make changes automatically to our project without the need to restart the server again)

###  Dependencies
- Fiber (Go framework that enahance high performance and support https protocol).
    if you are coming from Nodejs world, then fiber is the framework that looks like Express framework in Go.
- Mongodb Driver (Package that allow Go project to interact with mongodb database)


## Folder structure
- api folder contain the function that load the environment variables
- controller folder contain the handlers functions that handle the CRUD logic
- database folder contain the logic that initailized mongodb database connection
- model folder contain our struct that defined the fields, datatype and schema needed for our  
  databse
- router folder contain the routes for handling the CRUD operation on a person
- main file is the entry file which runs all the other files imported into it.

## Endpoints
These are lists of the available endpoints

- POST /api - Create a new person and store it in database.
- GET /api/{userId} - Retrieve person details by ID, name, or hobby from the database.
- GET /api - Get all person from the database.
- PUT /api/{userId} - Update a person by ID, name, or hobby to the database.
- DELETE /api/{userId} - Delete a person by ID, name, or hobby from the database.

## UML diagrams link
For visual insights into the project's structure, you can access UML diagrams via this link.
https://drive.google.com/file/d/1VTAiWQEudXSsaDMKDqBu-V4nGbxzPlvR/view?usp=sharing
