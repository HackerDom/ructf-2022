package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/usernamedt/doctor-service/models"
	"github.com/usernamedt/doctor-service/pkg/app"
	"github.com/usernamedt/doctor-service/pkg/e"
	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/service/jobservice"
	"net/http"
)

// @Router /api/v1/jobs/{id} [get]
func GetJob(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).String()
	valid := validation.Validation{}
	valid.AlphaNumeric(id, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	jobService := jobservice.JobService{}

	job, err := jobService.Get(id, appG.C)
	switch err.(type) {
	case nil:
		appG.Response(http.StatusOK, e.SUCCESS, job)
	case models.NonExistJobError:
		appG.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_JOB, nil)
	default:
		logging.Error(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_JOB_FAIL, nil)
	}
}

// @Router /api/v1/jobs/{id} [put]
func SubmitJob(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.PostForm("id")).String()
	valid := validation.Validation{}
	valid.AlphaNumeric(id, "id")

	logging.Infof("SubmitJob: received: %s", id)
	jobService := jobservice.JobService{}

	job, err := jobService.Add(id, appG.C)
	switch err.(type) {
	case nil:
		appG.Response(http.StatusOK, e.SUCCESS, job)
	default:
		logging.Error(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_JOB_FAIL, nil)
	}
}
