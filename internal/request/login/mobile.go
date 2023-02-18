package login

type MobileLoginRequest struct {
	Mobile string `json:"mobile" valid:"required,stringlength(11|11)"`
}
