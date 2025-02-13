package usecaseusers

import (
	"fmt"
	"log"
	"math"
	"net/http"

	apierror "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	requesthandler "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/models"
)

func GetAllUsers(db di.DI, pagination *requesthandler.Pagination, p *requesthandler.Pagination, q string) ([]models.UsersModel, error) {

	// Declare Query
	// query := "select id,first_name,middle_name,last_name,created_at,updated_at from users where first_name like CONCAT('%',?,'%') or middle_name like CONCAT('%',?,'%') or last_name like CONCAT('%',?,'%') LIMIT ? OFFSET ?"
	query := "select id,first_name,middle_name,last_name from users where first_name like CONCAT('%',?,'%') or middle_name like CONCAT('%',?,'%') or last_name like CONCAT('%',?,'%') LIMIT ? OFFSET ? "

	// Execute Query
	res, err := db.GetDB().Query(query, q, q, q, pagination.TotalDataInPage, (pagination.ActivePage-1)*pagination.TotalDataInPage)
	if err != nil {
		return nil, apierror.ApiError{StatusCode: http.StatusInternalServerError, Status: "error", Message: err.Error()}
	}

	// Create Slice Data Users
	var d []models.UsersModel

	// Read Data
	for res.Next() {
		// Temp Data
		var data models.UsersModel

		// Scan Read Data
		err := res.Scan(&data.ID, &data.Firstname, &data.MiddleName, &data.LastName)
		if err != nil {
			log.Println(err.Error())
			return nil, apierror.InternalServerError
		}

		// Append Data
		d = append(d, data)
	}

	// Count Data
	// Count Data
	queryCountData := "select COUNT(*) as jumlah from users where first_name like CONCAT('%',?,'%') or middle_name like CONCAT('%',?,'%') or last_name like CONCAT('%',?,'%')"
	if err := db.GetDB().QueryRow(queryCountData, q, q, q).Scan(&pagination.TotalData); err != nil {
		fmt.Println(err.Error())
		return nil, apierror.InternalServerError
	}

	// Mapping
	pagination.TotalPage = int(math.Max(float64(pagination.TotalData/pagination.TotalDataInPage), 1))

	return d, nil
}
