-- Create ticket (organizer_id inferred from event)
-- name: CreateTicket :one
INSERT INTO tickets (
    event_id, type, price, quantity_available
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- Get ticket details by ID (public)
-- name: GetTicket :one
SELECT * FROM tickets
WHERE id = $1;

-- List all tickets for an event (public)
-- name: ListTicketsByEvent :many
SELECT * FROM tickets
WHERE event_id = $1
ORDER BY type;

-- Update ticket (organizer-only via event ownership)
-- name: UpdateTicket :one
UPDATE tickets
SET
    quantity_available = $2,
    price = $3
WHERE id = $1
    RETURNING *;

-- Delete ticket (organizer-only)
-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE id = $1;