package mysql

import (
	"context"
	"database/sql"
	"project-version3/superindo-task/domain"

	"github.com/sirupsen/logrus"
)

type mysqlProductRepository struct {
	Conn *sql.DB
}

func NewMysqlProductRepository(conn *sql.DB) domain.ProductRepository {
	return &mysqlProductRepository{conn}
}

func (m *mysqlProductRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Product, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Product, 0)
	for rows.Next() {
		t := domain.Product{}
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.Author = authorID
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlProductRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Product, nextCursor string, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM product ORDER BY created_at LIMIT ? `

	res, err = m.fetch(ctx, query, num)
	if err != nil {
		return nil, "", err
	}

	return
}
func (m *mysqlProductRepository) GetByID(ctx context.Context, id int64) (res domain.Product, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM product WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Product{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlProductRepository) Store(ctx context.Context, a *domain.Product) (err error) {
	query := `INSERT  product SET title=? , content=? , author_id=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}
