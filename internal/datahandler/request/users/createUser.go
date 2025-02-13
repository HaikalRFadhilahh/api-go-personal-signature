package usersrequest

type CreateUsersRequest struct {
	FirstName  string  `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName   *string `json:"lastName"`
}
