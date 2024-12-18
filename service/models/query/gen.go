package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                 = new(Query)
	MovieMessageModel *movieMessageModel
	MovieTagsModel    *movieTagsModel
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	MovieMessageModel = &Q.MovieMessageModel
	MovieTagsModel = &Q.MovieTagsModel
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                db,
		MovieMessageModel: newMovieMessageModel(db, opts...),
		MovieTagsModel:    newMovieTagsModel(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	MovieMessageModel movieMessageModel
	MovieTagsModel    movieTagsModel
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		MovieMessageModel: q.MovieMessageModel.clone(db),
		MovieTagsModel:    q.MovieTagsModel.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		MovieMessageModel: q.MovieMessageModel.replaceDB(db),
		MovieTagsModel:    q.MovieTagsModel.replaceDB(db),
	}
}

type queryCtx struct {
	MovieMessageModel IMovieMessageModelDo
	MovieTagsModel    IMovieTagsModelDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		MovieMessageModel: q.MovieMessageModel.WithContext(ctx),
		MovieTagsModel:    q.MovieTagsModel.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
