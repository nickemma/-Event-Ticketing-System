package util

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomString(n int) string {
	var stringBuilder = strings.Builder{}
	stringBuilder.Grow(n)
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(c)
	}
	return stringBuilder.String()
}

func RandomName() pgtype.Text {
	// 20% chance to return NULL
	if rand.Float32() < 0.2 {
		return pgtype.Text{Valid: false}
	}

	// Generate realistic-looking names
	firstNames := []string{"John", "Jane", "Alex", "Sarah", "Mike", "Emma", "David", "Olivia"}
	lastNames := []string{"Doe", "Smith", "Johnson", "Brown", "Wilson", "Taylor", "Clark", "Lee"}

	name := fmt.Sprintf("%s %s",
		firstNames[rand.Intn(len(firstNames))],
		lastNames[rand.Intn(len(lastNames))],
	)

	return pgtype.Text{
		String: name,
		Valid:  true,
	}
}

func RandomEmail() string {
	// Generate base username
	username := RandomString(6 + rand.Intn(4)) // 6-9 characters

	// 30% chance to add number suffix
	if rand.Float32() < 0.3 {
		username += strconv.Itoa(rand.Intn(999))
	}

	// 25% chance to add separator
	if rand.Float32() < 0.25 {
		separators := []rune{'.', '_', '-'}
		pos := rand.Intn(len(username)-1) + 1
		username = username[:pos] + string(separators[rand.Intn(len(separators))]) + username[pos:]
	}

	domains := []string{
		"example.com",
		"gmail.com",
		"yahoo.com",
		"hotmail.com",
		"protonmail.com",
		"outlook.com",
		"mail.com",
		"test.com",
		"company.org",
		"university.edu",
	}

	return fmt.Sprintf("%s@%s", strings.ToLower(username), domains[rand.Intn(len(domains))])
}

func RandomHashPassword() string {
	return RandomString(8)
}

func RandomRole() string {
	roles := []string{"user", "organizer"}
	n := len(roles)
	return roles[rand.Intn(n)]
}
