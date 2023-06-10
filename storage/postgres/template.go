package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golanguzb70/go-gin-basicauth-postgres-monolithic-template/models"
)

func (r *postgresRepo) TemplateCreate(ctx context.Context, req *models.TemplateCreateReq) (*models.TemplateResponse, error) {
	res := &models.TemplateResponse{}
	query := r.Db.Builder.Insert("templates").Columns(
		"template_name",
	).Values(req.TemplateName).Suffix(
		"RETURNING id, template_name, created_at, updated_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.Id, &res.TemplateName,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Create()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *postgresRepo) TemplateGet(ctx context.Context, req *models.TemplateGetReq) (*models.TemplateResponse, error) {
	query := r.Db.Builder.Select("id, template_name, created_at, updated_at").
		From("templates").
		Where(squirrel.Eq{"id": req.Id})

	res := &models.TemplateResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.TemplateName,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Get()")
	}

	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *postgresRepo) TemplateFind(ctx context.Context, req *models.TemplateFindReq) (*models.TemplateFindResponse, error) {
	var (
		res *models.TemplateFindResponse
	)

	countQuery := r.Db.Builder.Select("count(1) as count").From("templates").Where("deleted at is null")
	err := countQuery.RunWith(r.Db.Db).QueryRow().Scan(&res.Count)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")

	}

	query := r.Db.Builder.Select("id, template_name, created_at, updated_at").
		From("templates").Where("deleted  at is null").OrderBy("id").Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &models.TemplateResponse{}
		err := rows.Scan(
			&temp.Id, &temp.TemplateName,
			&CreatedAt, &UpdatedAt,
		)
		if err != nil {
			return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")
		}

		temp.CreatedAt = CreatedAt.Format(time.RFC1123)
		temp.UpdatedAt = UpdatedAt.Format(time.RFC1123)
		res.Templates = append(res.Templates, temp)
	}

	return res, nil
}

func (r *postgresRepo) TemplateUpdate(ctx context.Context, req *models.TemplateUpdateReq) (*models.TemplateResponse, error) {
	mp := make(map[string]interface{})
	mp["template_name"] = req.TemplateName
	mp["updated_at"] = time.Now()
	query := r.Db.Builder.Update("templates").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, template_name, created_at, updated_at")

	res := &models.TemplateResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.TemplateName,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) Update()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *postgresRepo) TemplateDelete(ctx context.Context, req *models.TemplateDeleteReq) error {
	query := r.Db.Builder.Delete("templates").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) Delete()")
}
