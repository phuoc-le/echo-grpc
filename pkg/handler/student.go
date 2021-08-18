package handler

import (
	"github.com/labstack/echo/v4"
	"grpc-echo/pkg/app"
	"grpc-echo/pkg/models"
	"log"
	"net/http"
	"strconv"
)
// Get List User
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
// Get User By Id
func GetStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	student, err := getStudent(id)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, student)
}
func getStudent(id int)(*models.Students, error)  {
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

// Create User
func CreateStudent(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	student, err := createStudent(name, email, phone)

	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, student)
}

func createStudent(name, email, phone string) (*models.Students, error) {
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}
	err = db.AutoMigrate(&models.Students{})
	if err != nil {
		log.Panic(err)
	}
	student := models.Students{Name: name, Email: email, Phone: phone}
	if err := db.Create(&student).Error; err != nil {
		log.Panic(err)
		return nil, err
	}
	return &student, nil
}

//Update

func UpdateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	student, err := updateStudent(id ,name, email, phone)

	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, student)
}

func updateStudent(id int,name, email, phone string) (*models.Students, error) {
	db, err := app.GetDB()
	if err != nil {
		log.Panic(err)
	}
	var student models.Students
	db.First(&student, id)

	log.Print(student)
	//student.Name = name
	//student.Email = email
	//student.Phone =phone
	//if err := db.Save(&student).Error; err != nil {
	//	log.Panic(err)
	//	return nil, err
	//}
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
func deleteStudent(id int)(*models.Students, error)  {
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