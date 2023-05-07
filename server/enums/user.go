package enums

// UserStatus represents the user's current subscription status
// where an active user is a user with a free subscription or an ongoing paid subscription
// and an inactive user is a user with an overdue subscription by 10 days
type UserStatus string

const (
	ActiveStatus          UserStatus = "ACTIVE"
	InactiveStatus        UserStatus = "INACTIVE"
	UnverifiedEmailStatus UserStatus = "UNVERIFIED_EMAIL"
)
