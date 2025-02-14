// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AuditLog struct {
	ID int64 `json:"id"`
	// create/update/delete
	Action string `json:"action"`
	// event/ticket/transaction
	EntityType string             `json:"entity_type"`
	EntityID   int64              `json:"entity_id"`
	UserID     pgtype.Int8        `json:"user_id"`
	Timestamp  pgtype.Timestamptz `json:"timestamp"`
}

type Event struct {
	ID          int64              `json:"id"`
	OrganizerID int64              `json:"organizer_id"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	Venue       string             `json:"venue"`
	StartTime   pgtype.Timestamptz `json:"start_time"`
	EndTime     pgtype.Timestamptz `json:"end_time"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

type Ticket struct {
	ID      int64 `json:"id"`
	EventID int64 `json:"event_id"`
	// e.g., General, VIP
	Type              string             `json:"type"`
	Price             pgtype.Numeric     `json:"price"`
	QuantityAvailable int32              `json:"quantity_available"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}

type Transaction struct {
	ID          int64          `json:"id"`
	UserID      int64          `json:"user_id"`
	TicketID    int64          `json:"ticket_id"`
	Quantity    int32          `json:"quantity"`
	TotalAmount pgtype.Numeric `json:"total_amount"`
	// pending/success/failed
	Status    pgtype.Text        `json:"status"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	ID             int64       `json:"id"`
	Email          string      `json:"email"`
	HashedPassword string      `json:"hashed_password"`
	Name           pgtype.Text `json:"name"`
	// organizer/user
	Role      string             `json:"role"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
