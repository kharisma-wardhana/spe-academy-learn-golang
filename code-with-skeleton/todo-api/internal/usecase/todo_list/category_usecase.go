package todo_list_usecase

import (
	"context"
	"fmt"

	generalEntity "github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/entity"
	apperr "github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/error"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/helper"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/repository/mysql"
	mentity "github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/repository/mysql/entity"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/usecase"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/usecase/todo_list/entity"
	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type ICategoryUsecase interface {
	GetAllCategories(ctx context.Context) ([]*entity.CategoryResponse, error)
	GetCategoryByID(ctx context.Context, categoryID int64) (*entity.CategoryResponse, error)
	CreateCategory(ctx context.Context, categoryReq entity.CategoryReq) (*entity.CategoryResponse, error)
	UpdateCategoryByID(ctx context.Context, categoryReq entity.CategoryReq) error
	DeleteCategoryByID(ctx context.Context, categoryID int64) error
}

type CategoryUsecase struct {
	categoryRepo mysql.ITodoListCategoryRepository
}

func NewCategoryUsecase(categoryRepo mysql.ITodoListCategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{categoryRepo}
}

func (c *CategoryUsecase) GetAllCategories(ctx context.Context) ([]*entity.CategoryResponse, error) {
	funcName := "CategoryUsecase.GetAllCategories"
	captureFieldError := generalEntity.CaptureFields{}

	categories, err := c.categoryRepo.GetAll(ctx)
	if err != nil {
		helper.LogError("categoryRepo.GetAll", funcName, err, captureFieldError, "")

		return nil, err
	}

	var response []*entity.CategoryResponse
	for _, category := range categories {
		response = append(response, &entity.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   helper.ConvertToJakartaTime(category.CreatedAt),
		})
	}

	return response, nil
}

func (c *CategoryUsecase) GetCategoryByID(ctx context.Context, categoryID int64) (*entity.CategoryResponse, error) {
	funcName := "CategoryUsecase.GetCategoryByID"
	captureFieldError := generalEntity.CaptureFields{
		"category_id": helper.ToString(categoryID),
	}

	data, err := c.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		helper.LogError("categoryRepo.GetByID", funcName, err, captureFieldError, "")
		if errwrap.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperr.ErrUserNotFound()
		}
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.CategoryResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
	}, nil
}

func (c *CategoryUsecase) CreateCategory(ctx context.Context, categoryReq entity.CategoryReq) (*entity.CategoryResponse, error) {
	funcName := "CategoryUsecase.CreateCategory"
	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(categoryReq),
	}

	if errMsg := usecase.ValidateStruct(categoryReq); errMsg != "" {
		return nil, errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	category := &mentity.TodoListCategory{
		Name:        categoryReq.Name,
		Description: categoryReq.Description,
		CreatedAt:   helper.DatetimeNowJakarta(),
		CreatedBy:   categoryReq.CreatedBy,
	}

	if err := c.categoryRepo.Create(ctx, nil, category, false); err != nil {
		helper.LogError("categoryRepo.Create", funcName, err, captureFieldError, "")
		if errwrap.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errwrap.Wrap(apperr.ErrInvalidRequest(), "category already exists")
		}

		return nil, err
	}

	return &entity.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   helper.ConvertToJakartaTime(category.CreatedAt),
		CreatedBy:   category.CreatedBy,
	}, nil
}

func (c *CategoryUsecase) UpdateCategoryByID(ctx context.Context, categoryReq entity.CategoryReq) error {
	funcName := "CategoryUsecase.UpdateCategoryByID"
	captureFieldError := generalEntity.CaptureFields{
		"todo_list_category_id": helper.ToString(categoryReq.ID),
		"payload":               helper.ToString(categoryReq),
	}

	if errMsg := usecase.ValidateStruct(categoryReq); errMsg != "" {
		return errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	category := &mentity.TodoListCategory{
		ID:          categoryReq.ID,
		Name:        categoryReq.Name,
		Description: categoryReq.Description,
	}

	if err := mysql.DBTransaction(c.categoryRepo, func(trx mysql.TrxObj) error {
		lockedData, err := c.categoryRepo.LockByID(ctx, trx, categoryReq.ID)
		if err != nil {
			helper.LogError("cateoryRepo.LockByID", funcName, err, captureFieldError, "")

			return err
		}
		if lockedData == nil {
			helper.LogError("cateoryRepo.LockByID", funcName, fmt.Errorf("category not found"), captureFieldError, "")

			return fmt.Errorf("DATA IS NOT EXIST")
		}

		if err := c.categoryRepo.Update(ctx, trx, lockedData, &mentity.TodoListCategory{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   lockedData.CreatedAt, // Keep the original created_at
		}); err != nil {
			helper.LogError("cateoryRepo.Update", funcName, err, captureFieldError, "")

			return err
		}

		return nil
	}); err != nil {
		helper.LogError("cateoryRepo.DBTransaction", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}

func (c *CategoryUsecase) DeleteCategoryByID(ctx context.Context, categoryID int64) error {
	funcName := "CategoryUsecase.DeleteCategoryByID"
	captureFieldError := generalEntity.CaptureFields{
		"todo_list_categort_id": helper.ToString(categoryID),
	}

	err := c.categoryRepo.DeleteByID(ctx, nil, categoryID)
	if err != nil {
		helper.LogError("categoryRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
