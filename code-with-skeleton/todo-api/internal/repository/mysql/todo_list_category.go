package mysql

import (
	"context"

	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/config"
	apperr "github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/error"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/helper"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/code-with-skeleton/todo-api/internal/repository/mysql/entity"
	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type ITodoListCategoryRepository interface {
	TrxSupportRepo
	LockByID(ctx context.Context, dbTrx TrxObj, categoryID int64) (*entity.TodoListCategory, error)
	GetAll(ctx context.Context) ([]*entity.TodoListCategory, error)
	GetByID(ctx context.Context, categoryID int64) (*entity.TodoListCategory, error)
	Create(ctx context.Context, dbTrx TrxObj, category *entity.TodoListCategory) error
	Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, category *entity.TodoListCategory) error
	DeleteByID(ctx context.Context, dbTrx TrxObj, categoryID int64) error
}

type TodoListCategoryRepository struct {
	GormTrxSupport
}

func NewTodoListCategoryRepository(mysql *config.Mysql) *TodoListCategoryRepository {
	return &TodoListCategoryRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TodoListCategoryRepository) LockByID(ctx context.Context, dbTrx TrxObj, categoryID int64) (result *entity.TodoListCategory, err error) {
	funcName := "TodoListCategoryRepository.LockByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.Trx(dbTrx).
		Raw("SELECT * FROM todo_list_categories WHERE id = ? FOR UPDATE", categoryID).
		Scan(&result).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoListCategoryRepository) GetAll(ctx context.Context) ([]*entity.TodoListCategory, error) {
	var categories []*entity.TodoListCategory
	err := r.db.Raw("SELECT * FROM todo_list_categories").Scan(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *TodoListCategoryRepository) GetByID(ctx context.Context, categoryID int64) (*entity.TodoListCategory, error) {
	var category entity.TodoListCategory
	err := r.db.Raw("SELECT * FROM todo_list_categories WHERE id = ?", categoryID).Scan(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *TodoListCategoryRepository) Create(ctx context.Context, dbTrx TrxObj, category *entity.TodoListCategory) error {
	return r.db.Create(category).Error
}

func (r *TodoListCategoryRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, category *entity.TodoListCategory) error {
	return r.db.Save(category).Error
}

func (r *TodoListCategoryRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, categoryID int64) error {
	return r.db.Delete(&entity.TodoListCategory{}, categoryID).Error
}
