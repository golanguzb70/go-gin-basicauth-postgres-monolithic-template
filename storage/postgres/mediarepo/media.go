package mediarepo

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/storage/postgres/errorhandler"
)

func (r *MediaRepo) Create(ctx context.Context, req *CreateReq) (*FullResponse, error) {
	res := &FullResponse{}
	query := r.Db.Builder.Insert("medias").Columns(
		"media_name",
	).Values(req.Name).Suffix(
		"RETURNING id, media_name, created_at, updated_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.Id, &res.Name,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *MediaRepo) Create()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *MediaRepo) FindOne(ctx context.Context, req *FindOneReq) (*FullResponse, error) {
	query := r.Db.Builder.Select("id, media_name, created_at, updated_at").
		From("medias").
		Where(squirrel.Eq{"id": req.Id})

	res := &FullResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.Name,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *MediaRepo) FindOne()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *MediaRepo) FindList(ctx context.Context, req *FindListReq) ([]*FullResponse, error) {
	query := r.Db.Builder.Select("id, media_name, created_at, updated_at").
		From("medias").OrderBy("id").Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	res := []*FullResponse{}
	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *MediaRepo) FindList()")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &FullResponse{}
		err := rows.Scan(
			&temp.Id, &temp.Name,
			&CreatedAt, &UpdatedAt,
		)
		if err != nil {
			return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *MediaRepo) FindList()")
		}

		temp.CreatedAt = CreatedAt.Format(time.RFC1123)
		temp.UpdatedAt = UpdatedAt.Format(time.RFC1123)
		res = append(res, temp)
	}

	return res, nil
}

func (r *MediaRepo) Update(ctx context.Context, req *UpdateReq) (*FullResponse, error) {
	mp := make(map[string]interface{})
	mp["media_name"] = req.Name
	mp["updated_at"] = time.Now()
	query := r.Db.Builder.Update("medias").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, media_name, created_at, updated_at")

	res := &FullResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.Name,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *MediaRepo) Update()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *MediaRepo) Delete(ctx context.Context, req *DeleteReq) error {
	query := r.Db.Builder.Delete("medias").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return errorhandler.HandleDatabaseError(err, r.Log, "(r *MediaRepo) Delete()")
}
