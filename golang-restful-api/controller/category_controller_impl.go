package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryRequest := controller.CategoryService.Create(request.Context(), categoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRequest,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *CategoryControllerImpl) Update(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryRequest := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRequest,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *CategoryControllerImpl) Delete(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *CategoryControllerImpl) FindById(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(write, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(write, webResponse)
}
