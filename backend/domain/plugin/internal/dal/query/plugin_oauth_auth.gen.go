// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/coze-dev/coze-studio/backend/domain/plugin/internal/dal/model"
)

func newPluginOauthAuth(db *gorm.DB, opts ...gen.DOOption) pluginOauthAuth {
	_pluginOauthAuth := pluginOauthAuth{}

	_pluginOauthAuth.pluginOauthAuthDo.UseDB(db, opts...)
	_pluginOauthAuth.pluginOauthAuthDo.UseModel(&model.PluginOauthAuth{})

	tableName := _pluginOauthAuth.pluginOauthAuthDo.TableName()
	_pluginOauthAuth.ALL = field.NewAsterisk(tableName)
	_pluginOauthAuth.ID = field.NewInt64(tableName, "id")
	_pluginOauthAuth.UserID = field.NewString(tableName, "user_id")
	_pluginOauthAuth.PluginID = field.NewInt64(tableName, "plugin_id")
	_pluginOauthAuth.IsDraft = field.NewBool(tableName, "is_draft")
	_pluginOauthAuth.OauthConfig = field.NewField(tableName, "oauth_config")
	_pluginOauthAuth.AccessToken = field.NewString(tableName, "access_token")
	_pluginOauthAuth.RefreshToken = field.NewString(tableName, "refresh_token")
	_pluginOauthAuth.TokenExpiredAt = field.NewInt64(tableName, "token_expired_at")
	_pluginOauthAuth.NextTokenRefreshAt = field.NewInt64(tableName, "next_token_refresh_at")
	_pluginOauthAuth.LastActiveAt = field.NewInt64(tableName, "last_active_at")
	_pluginOauthAuth.CreatedAt = field.NewInt64(tableName, "created_at")
	_pluginOauthAuth.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_pluginOauthAuth.fillFieldMap()

	return _pluginOauthAuth
}

// pluginOauthAuth Plugin OAuth Authorization Code Info
type pluginOauthAuth struct {
	pluginOauthAuthDo

	ALL                field.Asterisk
	ID                 field.Int64  // Primary Key
	UserID             field.String // User ID
	PluginID           field.Int64  // Plugin ID
	IsDraft            field.Bool   // Is Draft Plugin
	OauthConfig        field.Field  // Authorization Code OAuth Config
	AccessToken        field.String // Access Token
	RefreshToken       field.String // Refresh Token
	TokenExpiredAt     field.Int64  // Token Expired in Milliseconds
	NextTokenRefreshAt field.Int64  // Next Token Refresh Time in Milliseconds
	LastActiveAt       field.Int64  // Last active time in Milliseconds
	CreatedAt          field.Int64  // Create Time in Milliseconds
	UpdatedAt          field.Int64  // Update Time in Milliseconds

	fieldMap map[string]field.Expr
}

func (p pluginOauthAuth) Table(newTableName string) *pluginOauthAuth {
	p.pluginOauthAuthDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pluginOauthAuth) As(alias string) *pluginOauthAuth {
	p.pluginOauthAuthDo.DO = *(p.pluginOauthAuthDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pluginOauthAuth) updateTableName(table string) *pluginOauthAuth {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.UserID = field.NewString(table, "user_id")
	p.PluginID = field.NewInt64(table, "plugin_id")
	p.IsDraft = field.NewBool(table, "is_draft")
	p.OauthConfig = field.NewField(table, "oauth_config")
	p.AccessToken = field.NewString(table, "access_token")
	p.RefreshToken = field.NewString(table, "refresh_token")
	p.TokenExpiredAt = field.NewInt64(table, "token_expired_at")
	p.NextTokenRefreshAt = field.NewInt64(table, "next_token_refresh_at")
	p.LastActiveAt = field.NewInt64(table, "last_active_at")
	p.CreatedAt = field.NewInt64(table, "created_at")
	p.UpdatedAt = field.NewInt64(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *pluginOauthAuth) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pluginOauthAuth) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["plugin_id"] = p.PluginID
	p.fieldMap["is_draft"] = p.IsDraft
	p.fieldMap["oauth_config"] = p.OauthConfig
	p.fieldMap["access_token"] = p.AccessToken
	p.fieldMap["refresh_token"] = p.RefreshToken
	p.fieldMap["token_expired_at"] = p.TokenExpiredAt
	p.fieldMap["next_token_refresh_at"] = p.NextTokenRefreshAt
	p.fieldMap["last_active_at"] = p.LastActiveAt
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p pluginOauthAuth) clone(db *gorm.DB) pluginOauthAuth {
	p.pluginOauthAuthDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pluginOauthAuth) replaceDB(db *gorm.DB) pluginOauthAuth {
	p.pluginOauthAuthDo.ReplaceDB(db)
	return p
}

type pluginOauthAuthDo struct{ gen.DO }

type IPluginOauthAuthDo interface {
	gen.SubQuery
	Debug() IPluginOauthAuthDo
	WithContext(ctx context.Context) IPluginOauthAuthDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPluginOauthAuthDo
	WriteDB() IPluginOauthAuthDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPluginOauthAuthDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPluginOauthAuthDo
	Not(conds ...gen.Condition) IPluginOauthAuthDo
	Or(conds ...gen.Condition) IPluginOauthAuthDo
	Select(conds ...field.Expr) IPluginOauthAuthDo
	Where(conds ...gen.Condition) IPluginOauthAuthDo
	Order(conds ...field.Expr) IPluginOauthAuthDo
	Distinct(cols ...field.Expr) IPluginOauthAuthDo
	Omit(cols ...field.Expr) IPluginOauthAuthDo
	Join(table schema.Tabler, on ...field.Expr) IPluginOauthAuthDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPluginOauthAuthDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPluginOauthAuthDo
	Group(cols ...field.Expr) IPluginOauthAuthDo
	Having(conds ...gen.Condition) IPluginOauthAuthDo
	Limit(limit int) IPluginOauthAuthDo
	Offset(offset int) IPluginOauthAuthDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPluginOauthAuthDo
	Unscoped() IPluginOauthAuthDo
	Create(values ...*model.PluginOauthAuth) error
	CreateInBatches(values []*model.PluginOauthAuth, batchSize int) error
	Save(values ...*model.PluginOauthAuth) error
	First() (*model.PluginOauthAuth, error)
	Take() (*model.PluginOauthAuth, error)
	Last() (*model.PluginOauthAuth, error)
	Find() ([]*model.PluginOauthAuth, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PluginOauthAuth, err error)
	FindInBatches(result *[]*model.PluginOauthAuth, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PluginOauthAuth) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPluginOauthAuthDo
	Assign(attrs ...field.AssignExpr) IPluginOauthAuthDo
	Joins(fields ...field.RelationField) IPluginOauthAuthDo
	Preload(fields ...field.RelationField) IPluginOauthAuthDo
	FirstOrInit() (*model.PluginOauthAuth, error)
	FirstOrCreate() (*model.PluginOauthAuth, error)
	FindByPage(offset int, limit int) (result []*model.PluginOauthAuth, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPluginOauthAuthDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pluginOauthAuthDo) Debug() IPluginOauthAuthDo {
	return p.withDO(p.DO.Debug())
}

func (p pluginOauthAuthDo) WithContext(ctx context.Context) IPluginOauthAuthDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pluginOauthAuthDo) ReadDB() IPluginOauthAuthDo {
	return p.Clauses(dbresolver.Read)
}

func (p pluginOauthAuthDo) WriteDB() IPluginOauthAuthDo {
	return p.Clauses(dbresolver.Write)
}

func (p pluginOauthAuthDo) Session(config *gorm.Session) IPluginOauthAuthDo {
	return p.withDO(p.DO.Session(config))
}

func (p pluginOauthAuthDo) Clauses(conds ...clause.Expression) IPluginOauthAuthDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pluginOauthAuthDo) Returning(value interface{}, columns ...string) IPluginOauthAuthDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pluginOauthAuthDo) Not(conds ...gen.Condition) IPluginOauthAuthDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pluginOauthAuthDo) Or(conds ...gen.Condition) IPluginOauthAuthDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pluginOauthAuthDo) Select(conds ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pluginOauthAuthDo) Where(conds ...gen.Condition) IPluginOauthAuthDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pluginOauthAuthDo) Order(conds ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pluginOauthAuthDo) Distinct(cols ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pluginOauthAuthDo) Omit(cols ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pluginOauthAuthDo) Join(table schema.Tabler, on ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pluginOauthAuthDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pluginOauthAuthDo) RightJoin(table schema.Tabler, on ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pluginOauthAuthDo) Group(cols ...field.Expr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pluginOauthAuthDo) Having(conds ...gen.Condition) IPluginOauthAuthDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pluginOauthAuthDo) Limit(limit int) IPluginOauthAuthDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pluginOauthAuthDo) Offset(offset int) IPluginOauthAuthDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pluginOauthAuthDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPluginOauthAuthDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pluginOauthAuthDo) Unscoped() IPluginOauthAuthDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pluginOauthAuthDo) Create(values ...*model.PluginOauthAuth) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pluginOauthAuthDo) CreateInBatches(values []*model.PluginOauthAuth, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pluginOauthAuthDo) Save(values ...*model.PluginOauthAuth) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pluginOauthAuthDo) First() (*model.PluginOauthAuth, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PluginOauthAuth), nil
	}
}

func (p pluginOauthAuthDo) Take() (*model.PluginOauthAuth, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PluginOauthAuth), nil
	}
}

func (p pluginOauthAuthDo) Last() (*model.PluginOauthAuth, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PluginOauthAuth), nil
	}
}

func (p pluginOauthAuthDo) Find() ([]*model.PluginOauthAuth, error) {
	result, err := p.DO.Find()
	return result.([]*model.PluginOauthAuth), err
}

func (p pluginOauthAuthDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PluginOauthAuth, err error) {
	buf := make([]*model.PluginOauthAuth, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pluginOauthAuthDo) FindInBatches(result *[]*model.PluginOauthAuth, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pluginOauthAuthDo) Attrs(attrs ...field.AssignExpr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pluginOauthAuthDo) Assign(attrs ...field.AssignExpr) IPluginOauthAuthDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pluginOauthAuthDo) Joins(fields ...field.RelationField) IPluginOauthAuthDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pluginOauthAuthDo) Preload(fields ...field.RelationField) IPluginOauthAuthDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pluginOauthAuthDo) FirstOrInit() (*model.PluginOauthAuth, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PluginOauthAuth), nil
	}
}

func (p pluginOauthAuthDo) FirstOrCreate() (*model.PluginOauthAuth, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PluginOauthAuth), nil
	}
}

func (p pluginOauthAuthDo) FindByPage(offset int, limit int) (result []*model.PluginOauthAuth, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pluginOauthAuthDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pluginOauthAuthDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pluginOauthAuthDo) Delete(models ...*model.PluginOauthAuth) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pluginOauthAuthDo) withDO(do gen.Dao) *pluginOauthAuthDo {
	p.DO = *do.(*gen.DO)
	return p
}
