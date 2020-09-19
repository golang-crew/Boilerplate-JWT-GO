// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set UsersQuerySet

// UsersQuerySet is an queryset type for Users
type UsersQuerySet struct {
	db *gorm.DB
}

// NewUsersQuerySet constructs new UsersQuerySet
func NewUsersQuerySet(db *gorm.DB) UsersQuerySet {
	return UsersQuerySet{
		db: db.Model(&Users{}),
	}
}

func (qs UsersQuerySet) w(db *gorm.DB) UsersQuerySet {
	return NewUsersQuerySet(db)
}

func (qs UsersQuerySet) Select(fields ...UsersDBSchemaField) UsersQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *Users) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Users) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) All(ret *[]Users) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Delete is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) Delete() error {
	return qs.db.Delete(Users{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(Users{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(Users{})
	return db.RowsAffected, db.Error
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) GetUpdater() UsersUpdater {
	return NewUsersUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDEq(ID uint64) UsersQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDGt(ID uint64) UsersQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDGte(ID uint64) UsersQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDIn(ID ...uint64) UsersQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDLt(ID uint64) UsersQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDLte(ID uint64) UsersQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDNe(ID uint64) UsersQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) IDNotIn(ID ...uint64) UsersQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) Limit(limit int) UsersQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) Offset(offset int) UsersQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs UsersQuerySet) One(ret *Users) error {
	return qs.db.First(ret).Error
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) OrderAscByID() UsersQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByPassword is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) OrderAscByPassword() UsersQuerySet {
	return qs.w(qs.db.Order("password ASC"))
}

// OrderAscByUsername is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) OrderAscByUsername() UsersQuerySet {
	return qs.w(qs.db.Order("username ASC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) OrderDescByID() UsersQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByPassword is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) OrderDescByPassword() UsersQuerySet {
	return qs.w(qs.db.Order("password DESC"))
}

// OrderDescByUsername is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) OrderDescByUsername() UsersQuerySet {
	return qs.w(qs.db.Order("username DESC"))
}

// PasswordEq is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordEq(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password = ?", password))
}

// PasswordGt is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordGt(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password > ?", password))
}

// PasswordGte is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordGte(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password >= ?", password))
}

// PasswordIn is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordIn(password ...string) UsersQuerySet {
	if len(password) == 0 {
		qs.db.AddError(errors.New("must at least pass one password in PasswordIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("password IN (?)", password))
}

// PasswordLike is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordLike(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password LIKE ?", password))
}

// PasswordLt is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordLt(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password < ?", password))
}

// PasswordLte is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordLte(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password <= ?", password))
}

// PasswordNe is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordNe(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password != ?", password))
}

// PasswordNotIn is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordNotIn(password ...string) UsersQuerySet {
	if len(password) == 0 {
		qs.db.AddError(errors.New("must at least pass one password in PasswordNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("password NOT IN (?)", password))
}

// PasswordNotlike is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) PasswordNotlike(password string) UsersQuerySet {
	return qs.w(qs.db.Where("password NOT LIKE ?", password))
}

// UsernameEq is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameEq(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username = ?", username))
}

// UsernameGt is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameGt(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username > ?", username))
}

// UsernameGte is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameGte(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username >= ?", username))
}

// UsernameIn is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameIn(username ...string) UsersQuerySet {
	if len(username) == 0 {
		qs.db.AddError(errors.New("must at least pass one username in UsernameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("username IN (?)", username))
}

// UsernameLike is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameLike(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username LIKE ?", username))
}

// UsernameLt is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameLt(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username < ?", username))
}

// UsernameLte is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameLte(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username <= ?", username))
}

// UsernameNe is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameNe(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username != ?", username))
}

// UsernameNotIn is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameNotIn(username ...string) UsersQuerySet {
	if len(username) == 0 {
		qs.db.AddError(errors.New("must at least pass one username in UsernameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("username NOT IN (?)", username))
}

// UsernameNotlike is an autogenerated method
// nolint: dupl
func (qs UsersQuerySet) UsernameNotlike(username string) UsersQuerySet {
	return qs.w(qs.db.Where("username NOT LIKE ?", username))
}

// SetID is an autogenerated method
// nolint: dupl
func (u UsersUpdater) SetID(ID uint64) UsersUpdater {
	u.fields[string(UsersDBSchema.ID)] = ID
	return u
}

// SetPassword is an autogenerated method
// nolint: dupl
func (u UsersUpdater) SetPassword(password string) UsersUpdater {
	u.fields[string(UsersDBSchema.Password)] = password
	return u
}

// SetUsername is an autogenerated method
// nolint: dupl
func (u UsersUpdater) SetUsername(username string) UsersUpdater {
	u.fields[string(UsersDBSchema.Username)] = username
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u UsersUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u UsersUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set UsersQuerySet

// ===== BEGIN of Users modifiers

// UsersDBSchemaField describes database schema field. It requires for method 'Update'
type UsersDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f UsersDBSchemaField) String() string {
	return string(f)
}

// UsersDBSchema stores db field names of Users
var UsersDBSchema = struct {
	ID       UsersDBSchemaField
	Username UsersDBSchemaField
	Password UsersDBSchemaField
}{

	ID:       UsersDBSchemaField("id"),
	Username: UsersDBSchemaField("username"),
	Password: UsersDBSchemaField("password"),
}

// Update updates Users fields by primary key
// nolint: dupl
func (o *Users) Update(db *gorm.DB, fields ...UsersDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":       o.ID,
		"username": o.Username,
		"password": o.Password,
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

		return fmt.Errorf("can't update Users %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// UsersUpdater is an Users updates manager
type UsersUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewUsersUpdater creates new Users updater
// nolint: dupl
func NewUsersUpdater(db *gorm.DB) UsersUpdater {
	return UsersUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Users{}),
	}
}

// ===== END of Users modifiers

// ===== END of all query sets