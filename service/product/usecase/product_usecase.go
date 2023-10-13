package usecase

import (
	"context"
	"time"

	"project-version3/superindo-task/service/domain"
)

type productUsecase struct {
	productRepo domain.ProductRepository
	// authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

// NewProductUsecase will create new an articleUsecase object representation of domain.ProductUsecase interface
func NewProductUsecase(a domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepo: a,
		// authorRepo:     ar,
		contextTimeout: timeout,
	}
}

func (a *productUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.Product, nextCursor string, err error) {
	// if num == 0 {
	// 	num = 10
	// }

	// ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	// defer cancel()

	// res, nextCursor, err = a.productRepo.Fetch(ctx, cursor, num)
	// if err != nil {
	// 	return nil, "", err
	// }
	return
}

func (a *productUsecase) GetByID(c context.Context, id int64) (res domain.Product, err error) {
	// ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	// defer cancel()

	// res, err = a.productRepo.GetByID(ctx, id)
	// if err != nil {
	// 	return
	// }
	return
}

func (a *productUsecase) Store(c context.Context, m *domain.Product) (err error) {
	// ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	// defer cancel()
	// existedProduct, _ := a.GetByTitle(ctx, m.Title)
	// if existedProduct != (domain.Product{}) {
	// 	return domain.ErrConflict
	// }

	// err = a.articleRepo.Store(ctx, m)
	return
}
