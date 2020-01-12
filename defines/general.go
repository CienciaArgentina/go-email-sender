package defines

const (
	StringEmpty = ""

	// Email Auth Keys
	Identity                  = ""
	CienciaArgentinaEmail     = "EMAIL_USERNAME"
	CienciaArgentinaPassword  = "EMAIL_PASSWORD"
	CienciaArgentinaEmailSmtp = "EMAIL_SMTP"
	CienciaArgentinaEmailSmtpPort = "smtp.gmail.com:587"

	// Email Metadata
	Mime = "MIME-version: 1.0;\nContent-Type: text/HTML;"

	// API Endpoints
	Ping = "/ping"
	PostEmail = "/email"

	DefaultPort = ":8080"
)
