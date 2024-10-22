// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set TodoQuerySet

// TodoQuerySet is an queryset type for Todo
type TodoQuerySet struct {
	db *gorm.DB
}

// NewTodoQuerySet constructs new TodoQuerySet
func NewTodoQuerySet(db *gorm.DB) TodoQuerySet {
	return TodoQuerySet{
		db: db.Model(&Todo{}),
	}
}

func (qs TodoQuerySet) w(db *gorm.DB) TodoQuerySet {
	return NewTodoQuerySet(db)
}

func (qs TodoQuerySet) Select(fields ...TodoDBSchemaField) TodoQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *Todo) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Todo) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) All(ret *[]Todo) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Delete is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) Delete() error {
	return qs.db.Delete(Todo{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Todo{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Todo{})
	return db.RowsAffected, db.Error
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) GetUpdater() TodoUpdater {
	return NewTodoUpdater(qs.db)
}

// Limit is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) Limit(limit int) TodoQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) Offset(offset int) TodoQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs TodoQuerySet) One(ret *Todo) error {
	return qs.db.First(ret).Error
}

// OrderAscByTitle is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) OrderAscByTitle() TodoQuerySet {
	return qs.w(qs.db.Order("title ASC"))
}

// OrderAscByUserID is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) OrderAscByUserID() TodoQuerySet {
	return qs.w(qs.db.Order("user_id ASC"))
}

// OrderDescByTitle is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) OrderDescByTitle() TodoQuerySet {
	return qs.w(qs.db.Order("title DESC"))
}

// OrderDescByUserID is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) OrderDescByUserID() TodoQuerySet {
	return qs.w(qs.db.Order("user_id DESC"))
}

// TitleEq is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleEq(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title = ?", title))
}

// TitleGt is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleGt(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title > ?", title))
}

// TitleGte is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleGte(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title >= ?", title))
}

// TitleIn is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleIn(title ...string) TodoQuerySet {
	if len(title) == 0 {
		qs.db.AddError(errors.New("must at least pass one title in TitleIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("title IN (?)", title))
}

// TitleLike is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleLike(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title LIKE ?", title))
}

// TitleLt is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleLt(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title < ?", title))
}

// TitleLte is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleLte(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title <= ?", title))
}

// TitleNe is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleNe(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title != ?", title))
}

// TitleNotIn is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleNotIn(title ...string) TodoQuerySet {
	if len(title) == 0 {
		qs.db.AddError(errors.New("must at least pass one title in TitleNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("title NOT IN (?)", title))
}

// TitleNotlike is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) TitleNotlike(title string) TodoQuerySet {
	return qs.w(qs.db.Where("title NOT LIKE ?", title))
}

// UserIDEq is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDEq(userID uint64) TodoQuerySet {
	return qs.w(qs.db.Where("user_id = ?", userID))
}

// UserIDGt is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDGt(userID uint64) TodoQuerySet {
	return qs.w(qs.db.Where("user_id > ?", userID))
}

// UserIDGte is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDGte(userID uint64) TodoQuerySet {
	return qs.w(qs.db.Where("user_id >= ?", userID))
}

// UserIDIn is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDIn(userID ...uint64) TodoQuerySet {
	if len(userID) == 0 {
		qs.db.AddError(errors.New("must at least pass one userID in UserIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_id IN (?)", userID))
}

// UserIDLt is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDLt(userID uint64) TodoQuerySet {
	return qs.w(qs.db.Where("user_id < ?", userID))
}

// UserIDLte is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDLte(userID uint64) TodoQuerySet {
	return qs.w(qs.db.Where("user_id <= ?", userID))
}

// UserIDNe is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDNe(userID uint64) TodoQuerySet {
	return qs.w(qs.db.Where("user_id != ?", userID))
}

// UserIDNotIn is an autogenerated method
// nolint: dupl
func (qs TodoQuerySet) UserIDNotIn(userID ...uint64) TodoQuerySet {
	if len(userID) == 0 {
		qs.db.AddError(errors.New("must at least pass one userID in UserIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("user_id NOT IN (?)", userID))
}

// SetTitle is an autogenerated method
// nolint: dupl
func (u TodoUpdater) SetTitle(title string) TodoUpdater {
	u.fields[string(TodoDBSchema.Title)] = title
	return u
}

// SetUserID is an autogenerated method
// nolint: dupl
func (u TodoUpdater) SetUserID(userID uint64) TodoUpdater {
	u.fields[string(TodoDBSchema.UserID)] = userID
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u TodoUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u TodoUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set TodoQuerySet

// ===== BEGIN of Todo modifiers

// TodoDBSchemaField describes database schema field. It requires for method 'Update'
type TodoDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f TodoDBSchemaField) String() string {
	return string(f)
}

// TodoDBSchema stores db field names of Todo
var TodoDBSchema = struct {
	UserID TodoDBSchemaField
	Title  TodoDBSchemaField
}{

	UserID: TodoDBSchemaField("user_id"),
	Title:  TodoDBSchemaField("title"),
}

// Update updates Todo fields by primary key
// nolint: dupl
func (o *Todo) Update(db *gorm.DB, fields ...TodoDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"user_id": o.UserID,
		"title":   o.Title,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Todo %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// TodoUpdater is an Todo updates manager
type TodoUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewTodoUpdater creates new Todo updater
// nolint: dupl
func NewTodoUpdater(db *gorm.DB) TodoUpdater {
	return TodoUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Todo{}),
	}
}

// ===== END of Todo modifiers

// ===== END of all query sets
