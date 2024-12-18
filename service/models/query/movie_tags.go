package query

import (
	"context"
	"easy-video-net/models/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newMovieTagsModel(db *gorm.DB, opts ...gen.DOOption) movieTagsModel {
	_movieTagsModel := movieTagsModel{}

	_movieTagsModel.movieTagsModelDo.UseDB(db, opts...)
	_movieTagsModel.movieTagsModelDo.UseModel(&model.MovieTagsModel{})

	tableName := _movieTagsModel.movieTagsModelDo.TableName()
	_movieTagsModel.ALL = field.NewAsterisk(tableName)
	_movieTagsModel.ID = field.NewInt64(tableName, "id")
	_movieTagsModel.MovieID = field.NewString(tableName, "movie_id")
	_movieTagsModel.Tag = field.NewString(tableName, "tag")
	_movieTagsModel.Source = field.NewString(tableName, "source")
	_movieTagsModel.CreateTime = field.NewTime(tableName, "create_time")
	_movieTagsModel.UpdateTime = field.NewTime(tableName, "update_time")

	_movieTagsModel.fillFieldMap()

	return _movieTagsModel
}

type movieTagsModel struct {
	movieTagsModelDo

	ALL        field.Asterisk
	ID         field.Int64
	MovieID    field.String
	Tag        field.String
	Source     field.String // 标签来源
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (m movieTagsModel) Table(newTableName string) *movieTagsModel {
	m.movieTagsModelDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m movieTagsModel) As(alias string) *movieTagsModel {
	m.movieTagsModelDo.DO = *(m.movieTagsModelDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *movieTagsModel) updateTableName(table string) *movieTagsModel {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.MovieID = field.NewString(table, "movie_id")
	m.Tag = field.NewString(table, "tag")
	m.Source = field.NewString(table, "source")
	m.CreateTime = field.NewTime(table, "create_time")
	m.UpdateTime = field.NewTime(table, "update_time")

	m.fillFieldMap()

	return m
}

func (m *movieTagsModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *movieTagsModel) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["id"] = m.ID
	m.fieldMap["movie_id"] = m.MovieID
	m.fieldMap["tag"] = m.Tag
	m.fieldMap["source"] = m.Source
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
}

func (m movieTagsModel) clone(db *gorm.DB) movieTagsModel {
	m.movieTagsModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m movieTagsModel) replaceDB(db *gorm.DB) movieTagsModel {
	m.movieTagsModelDo.ReplaceDB(db)
	return m
}

type movieTagsModelDo struct{ gen.DO }

type IMovieTagsModelDo interface {
	gen.SubQuery
	Debug() IMovieTagsModelDo
	WithContext(ctx context.Context) IMovieTagsModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMovieTagsModelDo
	WriteDB() IMovieTagsModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMovieTagsModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMovieTagsModelDo
	Not(conds ...gen.Condition) IMovieTagsModelDo
	Or(conds ...gen.Condition) IMovieTagsModelDo
	Select(conds ...field.Expr) IMovieTagsModelDo
	Where(conds ...gen.Condition) IMovieTagsModelDo
	Order(conds ...field.Expr) IMovieTagsModelDo
	Distinct(cols ...field.Expr) IMovieTagsModelDo
	Omit(cols ...field.Expr) IMovieTagsModelDo
	Join(table schema.Tabler, on ...field.Expr) IMovieTagsModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMovieTagsModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMovieTagsModelDo
	Group(cols ...field.Expr) IMovieTagsModelDo
	Having(conds ...gen.Condition) IMovieTagsModelDo
	Limit(limit int) IMovieTagsModelDo
	Offset(offset int) IMovieTagsModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMovieTagsModelDo
	Unscoped() IMovieTagsModelDo
	Create(values ...*model.MovieTagsModel) error
	CreateInBatches(values []*model.MovieTagsModel, batchSize int) error
	Save(values ...*model.MovieTagsModel) error
	First() (*model.MovieTagsModel, error)
	Take() (*model.MovieTagsModel, error)
	Last() (*model.MovieTagsModel, error)
	Find() ([]*model.MovieTagsModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MovieTagsModel, err error)
	FindInBatches(result *[]*model.MovieTagsModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MovieTagsModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMovieTagsModelDo
	Assign(attrs ...field.AssignExpr) IMovieTagsModelDo
	Joins(fields ...field.RelationField) IMovieTagsModelDo
	Preload(fields ...field.RelationField) IMovieTagsModelDo
	FirstOrInit() (*model.MovieTagsModel, error)
	FirstOrCreate() (*model.MovieTagsModel, error)
	FindByPage(offset int, limit int) (result []*model.MovieTagsModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMovieTagsModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m movieTagsModelDo) Debug() IMovieTagsModelDo {
	return m.withDO(m.DO.Debug())
}

func (m movieTagsModelDo) WithContext(ctx context.Context) IMovieTagsModelDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m movieTagsModelDo) ReadDB() IMovieTagsModelDo {
	return m.Clauses(dbresolver.Read)
}

func (m movieTagsModelDo) WriteDB() IMovieTagsModelDo {
	return m.Clauses(dbresolver.Write)
}

func (m movieTagsModelDo) Session(config *gorm.Session) IMovieTagsModelDo {
	return m.withDO(m.DO.Session(config))
}

func (m movieTagsModelDo) Clauses(conds ...clause.Expression) IMovieTagsModelDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m movieTagsModelDo) Returning(value interface{}, columns ...string) IMovieTagsModelDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m movieTagsModelDo) Not(conds ...gen.Condition) IMovieTagsModelDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m movieTagsModelDo) Or(conds ...gen.Condition) IMovieTagsModelDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m movieTagsModelDo) Select(conds ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m movieTagsModelDo) Where(conds ...gen.Condition) IMovieTagsModelDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m movieTagsModelDo) Order(conds ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m movieTagsModelDo) Distinct(cols ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m movieTagsModelDo) Omit(cols ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m movieTagsModelDo) Join(table schema.Tabler, on ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m movieTagsModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m movieTagsModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m movieTagsModelDo) Group(cols ...field.Expr) IMovieTagsModelDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m movieTagsModelDo) Having(conds ...gen.Condition) IMovieTagsModelDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m movieTagsModelDo) Limit(limit int) IMovieTagsModelDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m movieTagsModelDo) Offset(offset int) IMovieTagsModelDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m movieTagsModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMovieTagsModelDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m movieTagsModelDo) Unscoped() IMovieTagsModelDo {
	return m.withDO(m.DO.Unscoped())
}

func (m movieTagsModelDo) Create(values ...*model.MovieTagsModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m movieTagsModelDo) CreateInBatches(values []*model.MovieTagsModel, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m movieTagsModelDo) Save(values ...*model.MovieTagsModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m movieTagsModelDo) First() (*model.MovieTagsModel, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieTagsModel), nil
	}
}

func (m movieTagsModelDo) Take() (*model.MovieTagsModel, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieTagsModel), nil
	}
}

func (m movieTagsModelDo) Last() (*model.MovieTagsModel, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieTagsModel), nil
	}
}

func (m movieTagsModelDo) Find() ([]*model.MovieTagsModel, error) {
	result, err := m.DO.Find()
	return result.([]*model.MovieTagsModel), err
}

func (m movieTagsModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MovieTagsModel, err error) {
	buf := make([]*model.MovieTagsModel, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m movieTagsModelDo) FindInBatches(result *[]*model.MovieTagsModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m movieTagsModelDo) Attrs(attrs ...field.AssignExpr) IMovieTagsModelDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m movieTagsModelDo) Assign(attrs ...field.AssignExpr) IMovieTagsModelDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m movieTagsModelDo) Joins(fields ...field.RelationField) IMovieTagsModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m movieTagsModelDo) Preload(fields ...field.RelationField) IMovieTagsModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m movieTagsModelDo) FirstOrInit() (*model.MovieTagsModel, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieTagsModel), nil
	}
}

func (m movieTagsModelDo) FirstOrCreate() (*model.MovieTagsModel, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieTagsModel), nil
	}
}

func (m movieTagsModelDo) FindByPage(offset int, limit int) (result []*model.MovieTagsModel, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m movieTagsModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m movieTagsModelDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m movieTagsModelDo) Delete(models ...*model.MovieTagsModel) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *movieTagsModelDo) withDO(do gen.Dao) *movieTagsModelDo {
	m.DO = *do.(*gen.DO)
	return m
}
