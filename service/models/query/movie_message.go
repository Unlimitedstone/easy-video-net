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

func newMovieMessageModel(db *gorm.DB, opts ...gen.DOOption) movieMessageModel {
	_movieMessageModel := movieMessageModel{}

	_movieMessageModel.movieMessageModelDo.UseDB(db, opts...)
	_movieMessageModel.movieMessageModelDo.UseModel(&model.MovieMessageModel{})

	tableName := _movieMessageModel.movieMessageModelDo.TableName()
	_movieMessageModel.ALL = field.NewAsterisk(tableName)
	_movieMessageModel.ID = field.NewInt64(tableName, "id")
	_movieMessageModel.UniID = field.NewString(tableName, "uni_id")
	_movieMessageModel.Name = field.NewString(tableName, "name")
	_movieMessageModel.ImageURL = field.NewString(tableName, "image_url")
	_movieMessageModel.ImageSrc = field.NewString(tableName, "image_src")
	_movieMessageModel.Score = field.NewString(tableName, "score")
	_movieMessageModel.Type = field.NewString(tableName, "type")
	_movieMessageModel.Common = field.NewString(tableName, "common")
	_movieMessageModel.Source = field.NewString(tableName, "source")
	_movieMessageModel.CreateTime = field.NewTime(tableName, "create_time")
	_movieMessageModel.UpdateTime = field.NewTime(tableName, "update_time")
	_movieMessageModel.Introduce = field.NewString(tableName, "introduce")
	_movieMessageModel.Alias_ = field.NewString(tableName, "alias")
	_movieMessageModel.Duration = field.NewInt32(tableName, "duration")
	_movieMessageModel.ReleaseDate = field.NewString(tableName, "release_date")
	_movieMessageModel.Director = field.NewString(tableName, "director")
	_movieMessageModel.Actor = field.NewString(tableName, "actor")
	_movieMessageModel.Writer = field.NewString(tableName, "writer")
	_movieMessageModel.Country = field.NewString(tableName, "country")
	_movieMessageModel.Imdb = field.NewString(tableName, "imdb")

	_movieMessageModel.fillFieldMap()

	return _movieMessageModel
}

type movieMessageModel struct {
	movieMessageModelDo

	ALL         field.Asterisk
	ID          field.Int64
	UniID       field.String // 唯一id
	Name        field.String // 电影名字
	ImageURL    field.String // 电影url
	ImageSrc    field.String // 本地图片资源
	Score       field.String // 电影评分
	Type        field.String // 类型：Move:电影，TV：电视剧，Anime：动漫
	Common      field.String // 精选评论
	Source      field.String
	CreateTime  field.Time
	UpdateTime  field.Time
	Introduce   field.String // 简介
	Alias_      field.String // 别名
	Duration    field.Int32  // 时长（分钟）
	ReleaseDate field.String // 上映日期
	Director    field.String // 导演
	Actor       field.String // 主演
	Writer      field.String // 编剧
	Country     field.String // 制片国家
	Imdb        field.String // imdb id

	fieldMap map[string]field.Expr
}

func (m movieMessageModel) Table(newTableName string) *movieMessageModel {
	m.movieMessageModelDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m movieMessageModel) As(alias string) *movieMessageModel {
	m.movieMessageModelDo.DO = *(m.movieMessageModelDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *movieMessageModel) updateTableName(table string) *movieMessageModel {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.UniID = field.NewString(table, "uni_id")
	m.Name = field.NewString(table, "name")
	m.ImageURL = field.NewString(table, "image_url")
	m.ImageSrc = field.NewString(table, "image_src")
	m.Score = field.NewString(table, "score")
	m.Type = field.NewString(table, "type")
	m.Common = field.NewString(table, "common")
	m.Source = field.NewString(table, "source")
	m.CreateTime = field.NewTime(table, "create_time")
	m.UpdateTime = field.NewTime(table, "update_time")
	m.Introduce = field.NewString(table, "introduce")
	m.Alias_ = field.NewString(table, "alias")
	m.Duration = field.NewInt32(table, "duration")
	m.ReleaseDate = field.NewString(table, "release_date")
	m.Director = field.NewString(table, "director")
	m.Actor = field.NewString(table, "actor")
	m.Writer = field.NewString(table, "writer")
	m.Country = field.NewString(table, "country")
	m.Imdb = field.NewString(table, "imdb")

	m.fillFieldMap()

	return m
}

func (m *movieMessageModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *movieMessageModel) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 19)
	m.fieldMap["id"] = m.ID
	m.fieldMap["uni_id"] = m.UniID
	m.fieldMap["name"] = m.Name
	m.fieldMap["image_url"] = m.ImageURL
	m.fieldMap["image_src"] = m.ImageSrc
	m.fieldMap["score"] = m.Score
	m.fieldMap["type"] = m.Type
	m.fieldMap["common"] = m.Common
	m.fieldMap["source"] = m.Source
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
	m.fieldMap["introduce"] = m.Introduce
	m.fieldMap["alias"] = m.Alias_
	m.fieldMap["duration"] = m.Duration
	m.fieldMap["release_date"] = m.ReleaseDate
	m.fieldMap["director"] = m.Director
	m.fieldMap["actor"] = m.Actor
	m.fieldMap["writer"] = m.Writer
	m.fieldMap["country"] = m.Country
	m.fieldMap["imdb"] = m.Imdb
}

func (m movieMessageModel) clone(db *gorm.DB) movieMessageModel {
	m.movieMessageModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m movieMessageModel) replaceDB(db *gorm.DB) movieMessageModel {
	m.movieMessageModelDo.ReplaceDB(db)
	return m
}

type movieMessageModelDo struct{ gen.DO }

type IMovieMessageModelDo interface {
	gen.SubQuery
	Debug() IMovieMessageModelDo
	WithContext(ctx context.Context) IMovieMessageModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMovieMessageModelDo
	WriteDB() IMovieMessageModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMovieMessageModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMovieMessageModelDo
	Not(conds ...gen.Condition) IMovieMessageModelDo
	Or(conds ...gen.Condition) IMovieMessageModelDo
	Select(conds ...field.Expr) IMovieMessageModelDo
	Where(conds ...gen.Condition) IMovieMessageModelDo
	Order(conds ...field.Expr) IMovieMessageModelDo
	Distinct(cols ...field.Expr) IMovieMessageModelDo
	Omit(cols ...field.Expr) IMovieMessageModelDo
	Join(table schema.Tabler, on ...field.Expr) IMovieMessageModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMovieMessageModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMovieMessageModelDo
	Group(cols ...field.Expr) IMovieMessageModelDo
	Having(conds ...gen.Condition) IMovieMessageModelDo
	Limit(limit int) IMovieMessageModelDo
	Offset(offset int) IMovieMessageModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMovieMessageModelDo
	Unscoped() IMovieMessageModelDo
	Create(values ...*model.MovieMessageModel) error
	CreateInBatches(values []*model.MovieMessageModel, batchSize int) error
	Save(values ...*model.MovieMessageModel) error
	First() (*model.MovieMessageModel, error)
	Take() (*model.MovieMessageModel, error)
	Last() (*model.MovieMessageModel, error)
	Find() ([]*model.MovieMessageModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MovieMessageModel, err error)
	FindInBatches(result *[]*model.MovieMessageModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MovieMessageModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMovieMessageModelDo
	Assign(attrs ...field.AssignExpr) IMovieMessageModelDo
	Joins(fields ...field.RelationField) IMovieMessageModelDo
	Preload(fields ...field.RelationField) IMovieMessageModelDo
	FirstOrInit() (*model.MovieMessageModel, error)
	FirstOrCreate() (*model.MovieMessageModel, error)
	FindByPage(offset int, limit int) (result []*model.MovieMessageModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMovieMessageModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m movieMessageModelDo) Debug() IMovieMessageModelDo {
	return m.withDO(m.DO.Debug())
}

func (m movieMessageModelDo) WithContext(ctx context.Context) IMovieMessageModelDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m movieMessageModelDo) ReadDB() IMovieMessageModelDo {
	return m.Clauses(dbresolver.Read)
}

func (m movieMessageModelDo) WriteDB() IMovieMessageModelDo {
	return m.Clauses(dbresolver.Write)
}

func (m movieMessageModelDo) Session(config *gorm.Session) IMovieMessageModelDo {
	return m.withDO(m.DO.Session(config))
}

func (m movieMessageModelDo) Clauses(conds ...clause.Expression) IMovieMessageModelDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m movieMessageModelDo) Returning(value interface{}, columns ...string) IMovieMessageModelDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m movieMessageModelDo) Not(conds ...gen.Condition) IMovieMessageModelDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m movieMessageModelDo) Or(conds ...gen.Condition) IMovieMessageModelDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m movieMessageModelDo) Select(conds ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m movieMessageModelDo) Where(conds ...gen.Condition) IMovieMessageModelDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m movieMessageModelDo) Order(conds ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m movieMessageModelDo) Distinct(cols ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m movieMessageModelDo) Omit(cols ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m movieMessageModelDo) Join(table schema.Tabler, on ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m movieMessageModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m movieMessageModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m movieMessageModelDo) Group(cols ...field.Expr) IMovieMessageModelDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m movieMessageModelDo) Having(conds ...gen.Condition) IMovieMessageModelDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m movieMessageModelDo) Limit(limit int) IMovieMessageModelDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m movieMessageModelDo) Offset(offset int) IMovieMessageModelDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m movieMessageModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMovieMessageModelDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m movieMessageModelDo) Unscoped() IMovieMessageModelDo {
	return m.withDO(m.DO.Unscoped())
}

func (m movieMessageModelDo) Create(values ...*model.MovieMessageModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m movieMessageModelDo) CreateInBatches(values []*model.MovieMessageModel, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m movieMessageModelDo) Save(values ...*model.MovieMessageModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m movieMessageModelDo) First() (*model.MovieMessageModel, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieMessageModel), nil
	}
}

func (m movieMessageModelDo) Take() (*model.MovieMessageModel, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieMessageModel), nil
	}
}

func (m movieMessageModelDo) Last() (*model.MovieMessageModel, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieMessageModel), nil
	}
}

func (m movieMessageModelDo) Find() ([]*model.MovieMessageModel, error) {
	result, err := m.DO.Find()
	return result.([]*model.MovieMessageModel), err
}

func (m movieMessageModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MovieMessageModel, err error) {
	buf := make([]*model.MovieMessageModel, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m movieMessageModelDo) FindInBatches(result *[]*model.MovieMessageModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m movieMessageModelDo) Attrs(attrs ...field.AssignExpr) IMovieMessageModelDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m movieMessageModelDo) Assign(attrs ...field.AssignExpr) IMovieMessageModelDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m movieMessageModelDo) Joins(fields ...field.RelationField) IMovieMessageModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m movieMessageModelDo) Preload(fields ...field.RelationField) IMovieMessageModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m movieMessageModelDo) FirstOrInit() (*model.MovieMessageModel, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieMessageModel), nil
	}
}

func (m movieMessageModelDo) FirstOrCreate() (*model.MovieMessageModel, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MovieMessageModel), nil
	}
}

func (m movieMessageModelDo) FindByPage(offset int, limit int) (result []*model.MovieMessageModel, count int64, err error) {
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

func (m movieMessageModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m movieMessageModelDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m movieMessageModelDo) Delete(models ...*model.MovieMessageModel) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *movieMessageModelDo) withDO(do gen.Dao) *movieMessageModelDo {
	m.DO = *do.(*gen.DO)
	return m
}
