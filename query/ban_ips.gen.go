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

func newBanIP(db *gorm.DB, opts ...gen.DOOption) banIP {
	_banIP := banIP{}

	_banIP.banIPDo.UseDB(db, opts...)
	_banIP.banIPDo.UseModel(&model.BanIP{})

	tableName := _banIP.banIPDo.TableName()
	_banIP.ALL = field.NewAsterisk(tableName)
	_banIP.IP = field.NewString(tableName, "ip")
	_banIP.Attempts = field.NewInt(tableName, "attempts")
	_banIP.ExpiredAt = field.NewInt64(tableName, "expired_at")

	_banIP.fillFieldMap()

	return _banIP
}

type banIP struct {
	banIPDo

	ALL       field.Asterisk
	IP        field.String
	Attempts  field.Int
	ExpiredAt field.Int64

	fieldMap map[string]field.Expr
}

func (b banIP) Table(newTableName string) *banIP {
	b.banIPDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b banIP) As(alias string) *banIP {
	b.banIPDo.DO = *(b.banIPDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *banIP) updateTableName(table string) *banIP {
	b.ALL = field.NewAsterisk(table)
	b.IP = field.NewString(table, "ip")
	b.Attempts = field.NewInt(table, "attempts")
	b.ExpiredAt = field.NewInt64(table, "expired_at")

	b.fillFieldMap()

	return b
}

func (b *banIP) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *banIP) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 3)
	b.fieldMap["ip"] = b.IP
	b.fieldMap["attempts"] = b.Attempts
	b.fieldMap["expired_at"] = b.ExpiredAt
}

func (b banIP) clone(db *gorm.DB) banIP {
	b.banIPDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b banIP) replaceDB(db *gorm.DB) banIP {
	b.banIPDo.ReplaceDB(db)
	return b
}

type banIPDo struct{ gen.DO }

// FirstByID Where("id=@id")
func (b banIPDo) FirstByID(id int) (result *model.BanIP, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DeleteByID update @@table set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=@id
func (b banIPDo) DeleteByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("update ban_ips set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=? ")

	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (b banIPDo) Debug() *banIPDo {
	return b.withDO(b.DO.Debug())
}

func (b banIPDo) WithContext(ctx context.Context) *banIPDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b banIPDo) ReadDB() *banIPDo {
	return b.Clauses(dbresolver.Read)
}

func (b banIPDo) WriteDB() *banIPDo {
	return b.Clauses(dbresolver.Write)
}

func (b banIPDo) Session(config *gorm.Session) *banIPDo {
	return b.withDO(b.DO.Session(config))
}

func (b banIPDo) Clauses(conds ...clause.Expression) *banIPDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b banIPDo) Returning(value interface{}, columns ...string) *banIPDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b banIPDo) Not(conds ...gen.Condition) *banIPDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b banIPDo) Or(conds ...gen.Condition) *banIPDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b banIPDo) Select(conds ...field.Expr) *banIPDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b banIPDo) Where(conds ...gen.Condition) *banIPDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b banIPDo) Order(conds ...field.Expr) *banIPDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b banIPDo) Distinct(cols ...field.Expr) *banIPDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b banIPDo) Omit(cols ...field.Expr) *banIPDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b banIPDo) Join(table schema.Tabler, on ...field.Expr) *banIPDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b banIPDo) LeftJoin(table schema.Tabler, on ...field.Expr) *banIPDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b banIPDo) RightJoin(table schema.Tabler, on ...field.Expr) *banIPDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b banIPDo) Group(cols ...field.Expr) *banIPDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b banIPDo) Having(conds ...gen.Condition) *banIPDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b banIPDo) Limit(limit int) *banIPDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b banIPDo) Offset(offset int) *banIPDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b banIPDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *banIPDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b banIPDo) Unscoped() *banIPDo {
	return b.withDO(b.DO.Unscoped())
}

func (b banIPDo) Create(values ...*model.BanIP) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b banIPDo) CreateInBatches(values []*model.BanIP, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b banIPDo) Save(values ...*model.BanIP) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b banIPDo) First() (*model.BanIP, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BanIP), nil
	}
}

func (b banIPDo) Take() (*model.BanIP, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BanIP), nil
	}
}

func (b banIPDo) Last() (*model.BanIP, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BanIP), nil
	}
}

func (b banIPDo) Find() ([]*model.BanIP, error) {
	result, err := b.DO.Find()
	return result.([]*model.BanIP), err
}

func (b banIPDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BanIP, err error) {
	buf := make([]*model.BanIP, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b banIPDo) FindInBatches(result *[]*model.BanIP, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b banIPDo) Attrs(attrs ...field.AssignExpr) *banIPDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b banIPDo) Assign(attrs ...field.AssignExpr) *banIPDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b banIPDo) Joins(fields ...field.RelationField) *banIPDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b banIPDo) Preload(fields ...field.RelationField) *banIPDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b banIPDo) FirstOrInit() (*model.BanIP, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BanIP), nil
	}
}

func (b banIPDo) FirstOrCreate() (*model.BanIP, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BanIP), nil
	}
}

func (b banIPDo) FindByPage(offset int, limit int) (result []*model.BanIP, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b banIPDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b banIPDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b banIPDo) Delete(models ...*model.BanIP) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *banIPDo) withDO(do gen.Dao) *banIPDo {
	b.DO = *do.(*gen.DO)
	return b
}
