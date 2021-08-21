package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"grpc-echo/pkg/app"
	"grpc-echo/pkg/models"
	"log"
	"net/http"
	"strconv"
)

// GetStudents Get List User
func GetStudents(c echo.Context) error {
	students, err := getRepoStudents()
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, students)
}

func getRepoStudents() (*[]models.Students, error) {
	var students []models.Students
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}
	if err := db.Find(&students).Error; err != nil {
		log.Panic(err)
		return nil, err
	}
	return &students, nil
}

// GetStudent Get UserById
func GetStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	student, err := getStudent(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, student)
}
func getStudent(id int) (*models.Students, error) {
	var student models.Students
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}
	if err := db.First(&student, id).Error; err != nil {
		log.Panic(err)
		return nil, err
	}

	return &student, nil
}

// CreateStudent Create User
func CreateStudent(c echo.Context) error {
	student := models.Students{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	rs, err := createStudent(student)

	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, rs)
}

func createStudent(student models.Students) (*models.Students, error) {
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}
	err = db.AutoMigrate(&models.Students{})
	if err != nil {
		log.Panic(err)
	}
	if err := db.Create(&student).Error; err != nil {
		log.Panic(err)
		return nil, err
	}
	return &student, nil
}

// UpdateStudent /**** Update User ****/
func UpdateStudent(c echo.Context) error {
	student := models.Students{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	fmt.Println(student)
	res, err := updateStudent(student)

	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, res)
}

func updateStudent(student models.Students) (*models.Students, error) {
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Student Update", student)
	if err := db.Save(&student).Error; err != nil {
		log.Panic(err)
		return nil, err
	}
	return &student, nil
}

//Delete
func DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	student, err := deleteStudent(id)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, student)
}
func deleteStudent(id int) (*models.Students, error) {
	var student models.Students
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}
	if err := db.Delete(&student, id).Error; err != nil {
		log.Panic(err)
		return nil, err
	}

	return &student, nil
}
