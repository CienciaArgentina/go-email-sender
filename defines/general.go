package defines

const (
	StringEmpty = ""

	// Email Auth Keys
	Identity                      = ""
	CienciaArgentinaEmail         = "EMAIL_USERNAME"
	CienciaArgentinaPassword      = "EMAIL_PASSWORD"
	CienciaArgentinaEmailSmtp     = "EMAIL_SMTP"
	CienciaArgentinaEmailSmtpPort = "EMAL_SMTP_PORT"

	// Email Metadata
	Mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// API Endpoints
	Ping      = "/email/ping"
	PostEmail = "/email"

	DefaultPort = ":8081"
	Port = "EMAILSENDER_PORT"
)
