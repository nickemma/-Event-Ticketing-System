-- Create an audit log entry (e.g., when a user creates an event, called after critical actions)
-- name: CreateAuditLog :one
INSERT INTO audit_logs (
    action, entity_type, entity_id, user_id
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- List all logs (admin-only review)
-- name: ListAuditLogs :many
SELECT * FROM audit_logs
ORDER BY timestamp DESC
    LIMIT $1 OFFSET $2;
