package storage

import (
	"context"
	"errors"
	"github.com/MukizuL/hezzl-test/internal/errs"
	"github.com/MukizuL/hezzl-test/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

func (s *Storage) CreateGoods(ctx context.Context, projectId int, name string) (int, error) {
	tx, err := s.pg.Begin(ctx)
	if err != nil {
		s.logger.Error("Failed to begin transaction",
			zap.String("method", "CreateGoods"),
			zap.Int("projectId", projectId),
			zap.String("name", name),
			zap.Error(err))

		return 0, errs.ErrInternalServerError
	}
	defer tx.Rollback(ctx)

	var id int
	err = tx.QueryRow(ctx, `INSERT INTO goods (project_id, name) VALUES ($1, $2) RETURNING id`, projectId, name).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.ForeignKeyViolation:
				return 0, errs.ErrProjectNotFound
			default:
				s.logger.Error("Failed to create goods", zap.Error(err))
				return 0, errs.ErrInternalServerError
			}
		}

		s.logger.Error("Failed to create goods",
			zap.String("method", "CreateGoods"),
			zap.Int("projectId", projectId),
			zap.String("name", name),
			zap.Error(err))

		return 0, errs.ErrInternalServerError
	}

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Failed to commit transaction",
			zap.String("method", "CreateGoods"),
			zap.Int("projectId", projectId),
			zap.String("name", name),
			zap.Error(err))

		return 0, errs.ErrInternalServerError
	}

	return id, nil
}

func (s *Storage) GetGood(ctx context.Context, id int) (*models.Goods, error) {
	var goods models.Goods
	err := s.pg.QueryRow(ctx, `SELECT id, project_id, name, COALESCE(description, ''), priority, removed, created_at FROM goods WHERE id = $1`, id).
		Scan(&goods.ID, &goods.ProjectID, &goods.Name, &goods.Description, &goods.Priority, &goods.Removed, &goods.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			default:
				s.logger.Error("Failed to get a good", zap.Error(err))
				return nil, errs.ErrInternalServerError
			}
		}

		s.logger.Error("Failed to get a good",
			zap.String("method", "GetGood"),
			zap.Int("id", id),
			zap.Error(err))

		return nil, errs.ErrInternalServerError
	}

	return &goods, nil
}

func (s *Storage) GetGoods(ctx context.Context) ([]models.Goods, error) {
	var result []models.Goods
	rows, err := s.pg.Query(ctx, `SELECT id, project_id, name, COALESCE(description, ''), priority, removed, created_at FROM goods ORDER BY priority`)
	if err != nil {
		s.logger.Error("Failed to get goods",
			zap.String("method", "GetGoods"),
			zap.Error(err))

		return nil, errs.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var goods models.Goods
		err = rows.Scan(&goods.ID, &goods.ProjectID, &goods.Name, &goods.Description, &goods.Priority, &goods.Removed, &goods.CreatedAt)
		if err != nil {
			s.logger.Error("Error in row",
				zap.String("method", "GetGoods"),
				zap.Error(err))
			continue
		}

		result = append(result, goods)
	}

	if rows.Err() != nil {
		s.logger.Error("Error in rows",
			zap.String("method", "GetGoods"),
			zap.Error(err))

		return nil, errs.ErrInternalServerError
	}

	if len(result) == 0 {
		return nil, errs.ErrGoodsNotFound
	}

	return result, nil
}

func (s *Storage) UpdateGood(ctx context.Context, id, projectId int, name, description string) error {
	tx, err := s.pg.Begin(ctx)
	if err != nil {
		s.logger.Error("Failed to begin transaction",
			zap.String("method", "UpdateGood"),
			zap.Int("id", id),
			zap.Int("projectId", projectId),
			zap.Error(err))

		return errs.ErrInternalServerError
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, `SELECT id FROM goods WHERE id = $1 AND project_id = $2 AND removed = false FOR UPDATE`, id, projectId).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errs.ErrGoodsNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			default:
				s.logger.Error("Failed to get a good", zap.Error(err))
				return errs.ErrInternalServerError
			}
		}

		s.logger.Error("Failed to get a good",
			zap.String("method", "GetGood"),
			zap.Int("id", id),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	_, err = tx.Exec(ctx, `UPDATE goods SET name = $1, description = $2 WHERE id = $3 AND project_id = $4`, name, description, id, projectId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			default:
				s.logger.Error("Failed to update a good", zap.Error(err))
				return errs.ErrInternalServerError
			}
		}

		s.logger.Error("Failed to update a good",
			zap.String("method", "GetGood"),
			zap.Int("id", id),
			zap.Int("projectId", projectId),
			zap.String("name", name),
			zap.String("description", description),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Failed to commit transaction",
			zap.String("method", "CreateGoods"),
			zap.Int("projectId", projectId),
			zap.String("name", name),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	return nil
}

func (s *Storage) RemoveGoods(ctx context.Context, id, projectId int) error {
	tx, err := s.pg.Begin(ctx)
	if err != nil {
		s.logger.Error("Failed to begin transaction",
			zap.String("method", "RemoveGoods"),
			zap.Int("id", id),
			zap.Int("projectId", projectId),
			zap.Error(err))

		return errs.ErrInternalServerError
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, `SELECT id FROM goods WHERE id = $1 AND project_id = $2 AND removed = false FOR UPDATE`, id, projectId).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errs.ErrGoodsNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			default:
				s.logger.Error("Failed to get a good", zap.Error(err))
				return errs.ErrInternalServerError
			}
		}

		s.logger.Error("Failed to get a good",
			zap.String("method", "GetGood"),
			zap.Int("id", id),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	_, err = tx.Exec(ctx, `UPDATE goods SET removed = true WHERE id = $1 AND project_id = $2`, id, projectId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			default:
				s.logger.Error("Failed to update a good", zap.Error(err))
				return errs.ErrInternalServerError
			}
		}

		s.logger.Error("Failed to remove a good",
			zap.String("method", "RemoveGoods"),
			zap.Int("id", id),
			zap.Int("projectId", projectId),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	err = tx.Commit(ctx)
	if err != nil {
		s.logger.Error("Failed to commit transaction",
			zap.String("method", "RemoveGoods"),
			zap.Int("projectId", projectId),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	return nil
}
