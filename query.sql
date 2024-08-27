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

-- name: UpdateTask :exec
UPDATE tasks
	set name = $2,
	info = $3
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

