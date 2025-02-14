-- Create a transaction (e.g., when a user buys tickets)
-- name: CreateTransaction :one
INSERT INTO transactions (
    user_id, ticket_id, quantity, total_amount, status
) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING *;

-- Get transaction details by ID (user-specific)
-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 AND user_id = $2;  -- User can only view their own

-- List all transactions for a user
-- name: ListTransactionsByUser :many
SELECT * FROM transactions
WHERE user_id = $1
ORDER BY created_at DESC;

-- Update transaction status (e.g., "success" after payment, admin/organizer action)
-- name: UpdateTransactionStatus :one
UPDATE transactions
SET status = $2
WHERE id = $1
    RETURNING *;

-- Delete a transaction (use cautiously!)
-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;