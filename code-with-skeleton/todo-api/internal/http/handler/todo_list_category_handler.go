package handler

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/http/middleware"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/parser"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/presenter/json"
	todo_list_usecase "github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/usecase/todo_list"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/usecase/todo_list/entity"
)

type TodoListCategoryHandler struct {
	parser          parser.Parser
	presenter       json.JsonPresenter
	categoryUsecase todo_list_usecase.ICategoryUsecase
}

func NewTodoListCategoryHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	categoryUsecase todo_list_usecase.ICategoryUsecase,
) *TodoListCategoryHandler {
	return &TodoListCategoryHandler{parser, presenter, categoryUsecase}
}

func (w *TodoListCategoryHandler) Register(app fiber.Router) {
	app.Get("/categories", middleware.VerifyJWTToken, w.GetAll)
	app.Get("/categories/:id", middleware.VerifyJWTToken, w.GetByID)
	app.Post("/categories", middleware.VerifyJWTToken, w.Create)
	app.Put("/categories/:id", middleware.VerifyJWTToken, w.Update)
	app.Delete("/categories/:id", middleware.VerifyJWTToken, w.Delete)
}

// @Summary         Get All Todo List Categories
// @Description     Get all Todo List Categories
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Success			200 {object} entity.GeneralResponse{data=[]entity.CategoryResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/categories [get]
func (w *TodoListCategoryHandler) GetAll(c *fiber.Ctx) error {
	data, err := w.categoryUsecase.GetAllCategories(c.Context())
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Get Todo List Category by ID
// @Description     Get a Todo List Category by its ID
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List Category"
// @Success			200 {object} entity.GeneralResponse{data=entity.CategoryResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/categories/{id} [get]
func (w *TodoListCategoryHandler) GetByID(c *fiber.Ctx) error {
	categoryID, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.categoryUsecase.GetCategoryByID(c.Context(), categoryID)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Create Todo List Category
// @Description     Create a new Todo List Category
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           req body entity.CategoryReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse{data=entity.CategoryResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/categories [post]
func (w *TodoListCategoryHandler) Create(c *fiber.Ctx) error {
	var req entity.CategoryReq

	err := w.parser.ParserBodyRequestWithUserID(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.categoryUsecase.CreateCategory(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Update Todo List Category by ID
// @Description     Update an existing Todo List Category by its ID
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List Category"
// @Param			req body entity.CategoryReq true "Payload Request Body"
// @Success			200 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/categories/{id} [put]
func (w *TodoListCategoryHandler) Update(c *fiber.Ctx) error {
	categoryID, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	var req entity.CategoryReq
	err = w.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	req.ID = categoryID

	err = w.categoryUsecase.UpdateCategoryByID(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

// @Summary         Delete Todo List Category by ID
// @Description     Delete a Todo List Category by its ID
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List Category"
// @Success			200 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/categories/{id} [delete]
func (w *TodoListCategoryHandler) Delete(c *fiber.Ctx) error {
	categoryID, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.categoryUsecase.DeleteCategoryByID(c.Context(), categoryID)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
