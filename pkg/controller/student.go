package controller
import (
	"grpc-echo/pkg/models"
	"grpc-echo/pkg/storage"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetStudents
func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]models.Students, error) {
	students := []models.Students{}
	db, err := storage.Get()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.DB.Client.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}
