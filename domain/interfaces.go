package domain

type ReceptorObject struct {
	UserId  string `json:"user_id"`
	Role    string `json:"role"`
	Message string `json:"message"`
}

type SigninQueryObject struct {
	UserId string `json:"user_id"`
	Password  string `json:"password"`
}

type PatchUserObject struct {
	Nickname string `json:"nickname", omitempty`
	Comment string `json:"comment", omitempty`
}

type UserPublic struct {
	UserId string `json:"user_id", omitempty`
	Nickname string `json:"nickname", omitempty`
	Comment string `json:"comment", omitempty`
}

type UserPrivate struct {
	UserId string `json:"user_id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type SigninQueryResponse struct {
	Message string `json:"message"`
	User UserPublic `json:"user"`
}

type GetUserQueryResponse struct {
	Message string `json:"message"`
	User UserPublic `json:"user"`
}

type PatchUserResponse struct {
	Message string `json:"message"`
	Recipe []UserPublic `json:"recipe"`
}

type CloseResponse struct {
	Message string `json:"message"`
}

type Err struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Cause string `json:"cause"`
}

func (e Err) Error() string {
	return e.Message
}