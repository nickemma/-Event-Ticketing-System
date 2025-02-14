-- Create event (organizer_id comes from JWT)
-- name: CreateEvent :one
INSERT INTO events (
    organizer_id, title, description, venue, start_time, end_time
) VALUES (
             $1, $2, $3, $4, $5, $6
         ) RETURNING *;

-- Get event details by ID (public)
-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1;

-- Update event details (only organizer can modify)
-- name: UpdateEvent :one
UPDATE events
SET
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    venue = COALESCE($4, venue),
    start_time = COALESCE($5, start_time),
    end_time = COALESCE($6, end_time)
WHERE id = $1 AND organizer_id = $7  -- Enforce organizer ownership
    RETURNING *;

-- Delete event (organizer-only)
-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1 AND organizer_id = $2;

-- List organizer's events (for dashboard)
-- name: ListEventsByOrganizer :many
SELECT * FROM events
WHERE organizer_id = $1
ORDER BY start_time DESC
    LIMIT $2 OFFSET $3;

-- List public events (for users)
-- name: ListPublicEvents :many
SELECT * FROM events
WHERE start_time > NOW()  -- Only future events
ORDER BY start_time DESC
    LIMIT $1 OFFSET $2;