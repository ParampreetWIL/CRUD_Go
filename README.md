# CRUD GoLang

This is the simple ToDo API implemented in GoLang with Fiber framework. This project uses SQLc and PGX to interact with the database.

## Steps to Run

1. Look for username and password in `docker-compose.yml` in root directory.
2. Substitute the username and password in `postgres://<username>:<password>@localhost:5431/crud_go`
3. Run the project as `DATABASE_URL=postgres://bluedog:woof@localhost:5431/crud_go go run .`