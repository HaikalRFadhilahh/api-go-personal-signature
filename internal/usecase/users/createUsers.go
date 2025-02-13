package usecaseusers

import (
	"fmt"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	usersrequest "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request/users"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
)

func CreateUsers(di di.DI, data *usersrequest.CreateUsersRequest) (*usersrequest.CreateUsersRequest, error) {
	// Query Insert
	createUserQuery := "INSERT INTO users(first_name,middle_name,last_name) VALUES (?,?,?)"

	// Execute Query
	_, err := di.GetDB().Exec(createUserQuery, data.FirstName, data.MiddleName, data.LastName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}

	// Return Data
	return data, nil
}
