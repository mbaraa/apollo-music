package enums

// UserStatus represents the user's current subscription status
// where an active user is a user with a free subscription or an ongoing paid subscription
// an inactive user is a user with an inactive subscription or their subscription payment is overdue by 10 days
// and an unverified email is a user that didn't verify their email via OTP
type UserStatus string

const (
	ActiveStatus          UserStatus = "ACTIVE"
	InactiveStatus        UserStatus = "INACTIVE"
	UnverifiedEmailStatus UserStatus = "UNVERIFIED_EMAIL"
)
