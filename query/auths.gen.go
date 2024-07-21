// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/0xJacky/Nginx-UI/model"
)

func newAuth(db *gorm.DB, opts ...gen.DOOption) auth {
	_auth := auth{}

	_auth.authDo.UseDB(db, opts...)
	_auth.authDo.UseModel(&model.Auth{})

	tableName := _auth.authDo.TableName()
	_auth.ALL = field.NewAsterisk(tableName)
	_auth.ID = field.NewInt(tableName, "id")
	_auth.CreatedAt = field.NewTime(tableName, "created_at")
	_auth.UpdatedAt = field.NewTime(tableName, "updated_at")
	_auth.DeletedAt = field.NewField(tableName, "deleted_at")
	_auth.Name = field.NewString(tableName, "name")
	_auth.Password = field.NewString(tableName, "password")
	_auth.Status = field.NewBool(tableName, "status")

	_auth.fillFieldMap()

	return _auth
}

type auth struct {
	authDo

	ALL       field.Asterisk
	ID        field.Int
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Password  field.String
	Status    field.Bool

	fieldMap map[string]field.Expr
}

func (a auth) Table(newTableName string) *auth {
	a.authDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a auth) As(alias string) *auth {
	a.authDo.DO = *(a.authDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *auth) updateTableName(table string) *auth {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Name = field.NewString(table, "name")
	a.Password = field.NewString(table, "password")
	a.Status = field.NewBool(table, "status")

	a.fillFieldMap()

	return a
}

func (a *auth) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *auth) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["name"] = a.Name
	a.fieldMap["password"] = a.Password
	a.fieldMap["status"] = a.Status
}

func (a auth) clone(db *gorm.DB) auth {
	a.authDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a auth) replaceDB(db *gorm.DB) auth {
	a.authDo.ReplaceDB(db)
	return a
}

type authDo struct{ gen.DO }

// FirstByID Where("id=@id")
func (a authDo) FirstByID(id int) (result *model.Auth, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DeleteByID update @@table set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=@id
func (a authDo) DeleteByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("update auths set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a authDo) Debug() *authDo {
	return a.withDO(a.DO.Debug())
}

func (a authDo) WithContext(ctx context.Context) *authDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authDo) ReadDB() *authDo {
	return a.Clauses(dbresolver.Read)
}

func (a authDo) WriteDB() *authDo {
	return a.Clauses(dbresolver.Write)
}

func (a authDo) Session(config *gorm.Session) *authDo {
	return a.withDO(a.DO.Session(config))
}

func (a authDo) Clauses(conds ...clause.Expression) *authDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authDo) Returning(value interface{}, columns ...string) *authDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authDo) Not(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authDo) Or(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authDo) Select(conds ...field.Expr) *authDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authDo) Where(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authDo) Order(conds ...field.Expr) *authDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authDo) Distinct(cols ...field.Expr) *authDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authDo) Omit(cols ...field.Expr) *authDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authDo) Join(table schema.Tabler, on ...field.Expr) *authDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authDo) LeftJoin(table schema.Tabler, on ...field.Expr) *authDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authDo) RightJoin(table schema.Tabler, on ...field.Expr) *authDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authDo) Group(cols ...field.Expr) *authDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authDo) Having(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authDo) Limit(limit int) *authDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authDo) Offset(offset int) *authDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *authDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authDo) Unscoped() *authDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authDo) Create(values ...*model.Auth) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authDo) CreateInBatches(values []*model.Auth, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authDo) Save(values ...*model.Auth) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authDo) First() (*model.Auth, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) Take() (*model.Auth, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) Last() (*model.Auth, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) Find() ([]*model.Auth, error) {
	result, err := a.DO.Find()
	return result.([]*model.Auth), err
}

func (a authDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Auth, err error) {
	buf := make([]*model.Auth, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authDo) FindInBatches(result *[]*model.Auth, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authDo) Attrs(attrs ...field.AssignExpr) *authDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authDo) Assign(attrs ...field.AssignExpr) *authDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authDo) Joins(fields ...field.RelationField) *authDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authDo) Preload(fields ...field.RelationField) *authDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authDo) FirstOrInit() (*model.Auth, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) FirstOrCreate() (*model.Auth, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) FindByPage(offset int, limit int) (result []*model.Auth, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a authDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authDo) Delete(models ...*model.Auth) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authDo) withDO(do gen.Dao) *authDo {
	a.DO = *do.(*gen.DO)
	return a
}
