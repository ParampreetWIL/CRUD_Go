# CRUD GoLang

This is the simple ToDo API implemented in GoLang with Fiber framework. This project uses SQLc and PGX to interact with the database.


## Steps to Run

1. Start postgresql Docker container with command `docker compose up`
2. Look for username and password in `docker-compose.yml` in root directory.
3. Substitute the username and password in `postgres://<username>:<password>@localhost:5431/crud_go`
4. Run the project as `DATABASE_URL=postgres://bluedog:woof@localhost:5431/crud_go go run .`


## `.env` File

### `SERVER_PORT`

The port at which the GoLang CRUD Server will be hosted

### `CLIENT_ID`

Client ID can be retrieved from [Google Cloud Console](https://console.cloud.google.com/). Required for OAuth Support. 

### `CLIENT_SECRET`

Client Secret can be retrieved from [Google Cloud Console](https://console.cloud.google.com/). Required for OAuth Support. 

### `REDIRECT_URI`

URI where google will redirect after authentication. Redirect URL should be same as entered in [Google Cloud Console](https://console.cloud.google.com/) at the time of App Creation.

### `POSTGRESQL_URI`

PostgreSQL Database URL. Format `postgres://<username>:<password>@<host>:<port>/<db name>`. Look in `docker-compose.yml` file for username, password and other details. If you are using same `docker-compose.yml` file then set this as `postgres://bluedog:woof@localhost:5431/crud_go`

### `VAULT_URI`

Vault URI is a address of the Vault. If you are using `docker-compose.yml` then set it as `http://127.0.0.1:8200`. Vault is used to tokenize the data. Here it is used to tokenize email of users.

### `VAULT_TOKEN`

Vault Token is required for tokenization set it to anything. It will be used to encrypt the data. 

### `JWT_SECRET_KEY`

JWT Token is used to encrypt and sign the JWT Token so it can not be tempered by any party.
