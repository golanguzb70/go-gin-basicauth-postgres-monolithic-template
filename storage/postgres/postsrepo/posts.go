package postsrepo

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/storage/postgres/errorhandler"
)

func (r *PostsRepo) Create(ctx context.Context, req *CreateReq) (*FullResponse, error) {
	res := &FullResponse{}
	query := r.Db.Builder.Insert("posts").Columns(
		"title",
		"image",
		"content",
		"description",
	).Values(req.Title, req.Image, req.Content, req.Description).Suffix(
		"RETURNING id, title, image, content, description, created_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.Id, &res.Title,
		&res.Image, &res.Content,
		&res.Description, &CreatedAt,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *PostsRepo) Create()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *PostsRepo) FindOne(ctx context.Context, req *FindOneReq) (*FullResponse, error) {
	query := r.Db.Builder.Select("id, title, image, content, description, created_at").
		From("posts").
		Where(squirrel.Eq{"id": req.Id})

	res := &FullResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.Title,
		&res.Image, &res.Content,
		&res.Description, &CreatedAt,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *PostsRepo) FindOne()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *PostsRepo) FindList(ctx context.Context, req *FindListReq) ([]*FullResponse, error) {
	query := r.Db.Builder.Select("id, title, image, description, created_at").
		From("posts").OrderBy("id DESC").Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	res := []*FullResponse{}
	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *PostsRepo) FindList()")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &FullResponse{}
		err := rows.Scan(
			&temp.Id, &temp.Title,
			&temp.Image, &temp.Description, &CreatedAt,
		)
		if err != nil {
			return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *PostsRepo) FindList()")
		}

		temp.CreatedAt = CreatedAt.Format(time.RFC1123)
		res = append(res, temp)
	}

	return res, nil
}

func (r *PostsRepo) Update(ctx context.Context, req *UpdateReq) (*FullResponse, error) {
	mp := make(map[string]interface{})
	mp["title"] = req.Title
	mp["image"] = req.Image
	mp["content"] = req.Content
	mp["description"] = req.Description
	query := r.Db.Builder.Update("posts").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, title, image, content, description, created_at")

	res := &FullResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.Title,
		&res.Image, &res.Content,
		&res.Description, &CreatedAt,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *PostsRepo) Update()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *PostsRepo) Delete(ctx context.Context, req *DeleteReq) error {
	query := r.Db.Builder.Delete("posts").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return errorhandler.HandleDatabaseError(err, r.Log, "(r *PostsRepo) Delete()")
}
