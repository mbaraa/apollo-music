package env

import (
	"net"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// if len(MachineAddress()) == 0 {
	// os.Setenv("MACHINE_IP", getMachineIP())
	// }
}

func PortNumber() string         { return os.Getenv("PORT") }
func DBUser() string             { return os.Getenv("DB_USER") }
func DBPassword() string         { return os.Getenv("DB_PASSWORD") }
func DBHost() string             { return os.Getenv("DB_HOST") }
func AllowedClients() string     { return os.Getenv("ALLOWED_CLIENTS") }
func MachineAddress() string     { return os.Getenv("MACHINE_IP") }
func JWTSecret() []byte          { return []byte(os.Getenv("JWT_SECRET")) }
func Development() bool          { return os.Getenv("DEVELOPMENT") == "true" }
func MusicDirectory() string     { return os.Getenv("MUSIC_DIRECTORY") }
func GoogleClientId() string     { return os.Getenv("GOOGLE_CLIENT_ID") }
func GoogleClientSecret() string { return os.Getenv("GOOGLE_CLIENT_SECRET") }
func MailerHost() string         { return os.Getenv("MAILER_HOST") }
func MailerPort() string         { return os.Getenv("MAILER_PORT") }
func MailerFrom() string         { return os.Getenv("MAILER_FROM") }
func MailerPassword() string     { return os.Getenv("MAILER_PASSWORD") }
func FrontendAddress() string    { return os.Getenv("FRONTEND_ADDRESS") }
func StripeSecretKey() string    { return os.Getenv("STRIPE_SK") }

func getMachineIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:"+PortNumber())
	if err != nil {
		panic(err)
	}

	err = conn.Close()
	if err != nil {
		panic(err)
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
