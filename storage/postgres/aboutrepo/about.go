package aboutrepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/golanguzb70/go-gin-basicauth-monolithic-template/storage/postgres/errorhandler"
)

func (r *AboutRepo) Create(ctx context.Context, req *CreateReq) (*FullResponse, error) {
	res := &FullResponse{}
	query := r.Db.Builder.Insert("about_me").Columns(
		"title, intro, resume_link, linkedin, youtube, faceebook, telegram, photo",
	).Values(req.Title, req.Intro, req.ResumeLink, req.LinkedIn, req.Youtube, req.Facebook, req.Telegram, req.Photo).Suffix(
		"RETURNING id, title, intro, resume_link, linkedin, youtube, faceebook, telegram, photo")

	err := query.RunWith(r.Db.Db).Scan(
		&res.Id, &res.Title, &res.Intro, &res.ResumeLink,
		&res.LinkedIn, &res.Youtube, &res.Facebook, &res.Telegram, &res.Photo,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *AboutRepo) Create()")
	}
	return res, nil
}

func (r *AboutRepo) FindOne(ctx context.Context, req *FindOneReq) (*FullResponse, error) {
	query := r.Db.Builder.Select("id, title, intro, resume_link, linkedin, youtube, faceebook, telegram, photo").
		From("about_me").
		Where(squirrel.Eq{"id": req.Id})

	res := &FullResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.Title, &res.Intro, &res.ResumeLink,
		&res.LinkedIn, &res.Youtube, &res.Facebook, &res.Telegram, &res.Photo,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *AboutRepo) FindOne()")
	}

	return res, nil
}

func (r *AboutRepo) Update(ctx context.Context, req *UpdateReq) (*FullResponse, error) {
	mp := make(map[string]interface{})
	mp["title"] = req.Title
	mp["intro"] = req.Intro
	mp["resume_link"] = req.ResumeLink
	mp["telegram"] = req.Telegram
	mp["faceebook"] = req.Facebook
	mp["youtube"] = req.Youtube
	mp["linkedin"] = req.LinkedIn
	mp["photo"] = req.Photo

	query := r.Db.Builder.Update("about_me").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, title, intro, resume_link, linkedin, youtube, faceebook, telegram, photo")

	res := &FullResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.Title, &res.Intro, &res.ResumeLink,
		&res.LinkedIn, &res.Youtube, &res.Facebook, &res.Telegram, &res.Photo,
	)
	if err != nil {
		return res, errorhandler.HandleDatabaseError(err, r.Log, "(r *AboutRepo) Update()")
	}

	return res, nil
}

func (r *AboutRepo) Delete(ctx context.Context, req *DeleteReq) error {
	query := r.Db.Builder.Delete("about_me").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return errorhandler.HandleDatabaseError(err, r.Log, "(r *AboutRepo) Delete()")
}
