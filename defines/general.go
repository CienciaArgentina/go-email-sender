package defines

const (
	StringEmpty = ""

	// Email Auth Keys
	Identity                  = ""
	CienciaArgentinaEmail     = "EMAIL_USERNAME"
	CienciaArgentinaPassword  = "EMAIL_PASSWORD"
	CienciaArgentinaEmailSmtp = "smtp.gmail.com"
	CienciaArgentinaEmailSmtpPort = "smtp.gmail.com:587"

	// Email Metadata
	Mime = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"

	// API Endpoints
	Ping = "/ping"
	PostEmail = "/email"

	Port = ":8080"
)
