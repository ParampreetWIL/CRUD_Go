-- name: GetDoneTasks :many
SELECT * FROM tasks
WHERE isDone = true;

-- name: GetPendingTasks :many
SELECT * FROM tasks
WHERE isDone = false;

-- name: GetAllTasks :many
SELECT * FROM tasks;

-- name: AddTask :one
INSERT INTO tasks (
	name,
	info,
	isDone
) VALUES (
	$1, $2, false
)
RETURNING *;

-- name: AddUser :one
INSERT INTO users (
	email_token,
	jwt_token,
	name
) VALUES (
	$1, $2, $3
)
RETURNING *;

-- name: GetUserByJWT :many
SELECT * FROM users 
WHERE jwt_token = $1;

-- name: UpdateTask :exec
UPDATE tasks
	set name = $2,
	info = $3,
	isDone = $4
WHERE id = $1;

-- name: UpdateTaskAsDone :exec
UPDATE tasks
	SET isDone = true
WHERE id = $1;

-- name: UpdateTaskAsNotDone :exec
UPDATE tasks
	SET isDone = false
WHERE id = $1;

-- name: DeleteTasks :exec
DELETE FROM tasks
WHERE id = $1;

