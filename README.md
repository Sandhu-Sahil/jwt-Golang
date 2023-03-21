# JWT Token Authentication in Golang and MongoDB Backend

This repository contains a sample project for implementing JWT token authentication in a Golang and MongoDB backend. The project uses middlewares for verification and proper validation of login and register data.

## Prerequisites
- Golang
- MongoDB
- Postman or any similar tool for API testing

## Installation

1. Clone the repository
2. Navigate to the project directory
3. Run `go mod download` to download the required dependencies
4. Create a `.env` file and add the following variables:
    ```
    PORT=<port_number>
    MONGO_URI=<mongodb_uri>
    JWT_SECRET=<jwt_secret_key>
    ```

5. Run the project using `go run main.go`

## API Endpoints

### /user/register

* POST Request
* Endpoint - http://localhost:<port_number>/user/register
* Request Body - JSON

    ```
    {
    "name": "<user_name>",
    "email": "<user_email>",
    "password": "<user_password>"
    }
    ```

* Response - JSON
    ```
    {
    "message": "User registered successfully!"
    }
    ```


### /user/login

* POST Request
* Endpoint - http://localhost:<port_number>/user/login
* Request Body - JSON
    ```
    {
    "email": "<user_email>",
    "password": "<user_password>"
    }
    ```

* Response - JSON
    ```
    {
    "token": "<jwt_token>"
    }
    ```


### /check/get/:id

* GET Request
* Endpoint - http://localhost:<port_number>/check/get/:id
* Headers - 
    * Authorization: Bearer <jwt_token>
* Response - JSON
    ```
    {
    "message": "User authorized successfully!",
    "id": "<user_id>"
    }
    ```


## Middlewares

The project uses the following middlewares for JWT token verification and proper validation of login and register data:

1. `validateRegisterData` - Validates the user registration data before creating a new user.
2. `validateLoginData` - Validates the user login data before generating a new JWT token.
3. `verifyToken` - Verifies the JWT token sent in the request headers before granting access to the protected routes.

## Conclusion

This project provides a basic setup for implementing JWT token authentication in a Golang and MongoDB backend. You can customize and extend the project based on your specific requirements.
