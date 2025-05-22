package handler

import (
	"employee_management_api/controller"
	"employee_management_api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateEmployees(ctx *gin.Context) {
	var employee model.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := controller.CreateEmployee(employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"Message": "Employee Created Successfully"})
}

func GetEmployees(ctx *gin.Context) {
	employee, err := controller.GetEmployees()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}
	ctx.JSON(http.StatusOK, employee)
}

func GetEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Employee Id"})
		return
	}

	employee, err := controller.GetEmployeeById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}

/*func UpdateEmployee(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Employee Id"})
		return
	}

	employee, err := controller.UpdateEmployee(id)
	if err != nil {

	}
}*/
