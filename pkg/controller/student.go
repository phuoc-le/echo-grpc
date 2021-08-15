package controller
import (
	"grpc-echo/pkg/app"
	"grpc-echo/pkg/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	students, err := GetRepoStudents()
	if err != nil {
		log.Panic(err)
	}
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]models.Students, error) {
	var students []models.Students
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
