package usecasesignature

import (
	"fmt"
	"math"
	"net/http"

	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/apierror"
	requesthandler "github.com/HaikalRFadhilahh/api-go-personal-signature/internal/datahandler/request"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/di"
	"github.com/HaikalRFadhilahh/api-go-personal-signature/internal/models"
)

func GetAllSignature(di di.DI, q string, pagination *requesthandler.Pagination) ([]models.SignatureModel, error) {
	// Declare Query
	query := "SELECT s.id,s.user_id,s.token,s.purpose,s.valid_until,u.id,u.first_name,u.middle_name,u.last_name FROM signature s INNER JOIN users u ON s.user_id = u.id WHERE u.first_name LIKE CONCAT('%',?,'%') or u.middle_name LIKE CONCAT('%',?,'%') or u.last_name LIKE CONCAT('%',?,'%') LIMIT ? OFFSET ?"
	countQuery := "SELECT COUNT(*) FROM signature s INNER JOIN users u ON s.user_id = u.id WHERE u.first_name LIKE CONCAT('%',?,'%') or u.middle_name LIKE CONCAT('%',?,'%') or u.last_name LIKE CONCAT('%',?,'%')"

	// Exec Query
	res, err := di.GetDB().Query(query, q, q, q, pagination.TotalDataInPage, (pagination.ActivePage-1)*pagination.TotalDataInPage)
	if err != nil {
		fmt.Println(err.Error())
		return nil, apierror.ApiError{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		}
	}

	// Counting Row Data
	var totalDataQuery int
	if err := di.GetDB().QueryRow(countQuery, q, q, q).Scan(&totalDataQuery); err != nil {
		fmt.Println(err.Error())
		return nil, apierror.ApiError{
			StatusCode: http.StatusInternalServerError,
			Status:     "error",
			Message:    err.Error(),
		}
	}

	// Ajust Pagination Data
	pagination.TotalData = totalDataQuery
	pagination.TotalPage = int(math.Max(float64(pagination.TotalData/pagination.TotalDataInPage), 1))

	// Create Models To Save The Data
	var data []models.SignatureModel

	// Loop Reading Data
	for res.Next() {
		// Create Temp Models Data
		var tempData models.SignatureModel

		// Scanning Data
		if err := res.Scan(&tempData.ID, &tempData.UserID, &tempData.Token, &tempData.Purpose, &tempData.ValidUntil, &tempData.User.ID, &tempData.User.Firstname, &tempData.User.MiddleName, &tempData.User.LastName); err != nil {
			return nil, apierror.ApiError{
				StatusCode: http.StatusInternalServerError,
				Status:     "error",
				Message:    err.Error(),
			}
		}

		// Append Data To Main Models
		data = append(data, tempData)
	}

	// Return Data
	return data, nil
}
