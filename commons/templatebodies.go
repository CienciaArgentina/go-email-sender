package commons

type ConfirmationMailBody struct {
	TokenizedUrl string
}

type ForgotUsernameBody struct {
	Username string
}

type SendPasswordResetBody struct {
	URL string
}
