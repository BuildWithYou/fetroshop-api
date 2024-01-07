package gormhelper

import (
	"errors"

	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"gorm.io/gorm"
)

// ErrRecordNotFound record not found error
func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsRecordNotFoundAndNotNil(err error) bool {
	return validatorhelper.IsNotNil(err) && errors.Is(err, gorm.ErrRecordNotFound)
}

// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
func IsInvalidTransaction(err error) bool {
	return errors.Is(err, gorm.ErrInvalidTransaction)
}

// ErrNotImplemented not implemented
func IsNotImplemented(err error) bool {
	return errors.Is(err, gorm.ErrNotImplemented)
}

// ErrMissingWhereClause missing where clause
func IsMissingWhereClause(err error) bool {
	return errors.Is(err, gorm.ErrMissingWhereClause)
}

// ErrUnsupportedRelation unsupported relations
func IsUnsupportedRelation(err error) bool {
	return errors.Is(err, gorm.ErrUnsupportedRelation)
}

// ErrPrimaryKeyRequired primary keys required
func IsPrimaryKeyRequired(err error) bool {
	return errors.Is(err, gorm.ErrPrimaryKeyRequired)
}

// ErrModelValueRequired model value required
func IsModelValueRequired(err error) bool {
	return errors.Is(err, gorm.ErrModelValueRequired)
}

// ErrModelAccessibleFieldsRequired model accessible fields required
func IsModelAccessibleFieldsRequired(err error) bool {
	return errors.Is(err, gorm.ErrModelAccessibleFieldsRequired)
}

// ErrSubQueryRequired sub query required
func IsSubQueryRequired(err error) bool {
	return errors.Is(err, gorm.ErrSubQueryRequired)
}

// ErrInvalidData unsupported data
func IsInvalidData(err error) bool {
	return errors.Is(err, gorm.ErrInvalidData)
}

// ErrUnsupportedDriver unsupported driver
func IsUnsupportedDriver(err error) bool {
	return errors.Is(err, gorm.ErrUnsupportedDriver)
}

// ErrRegistered registered
func IsRegistered(err error) bool {
	return errors.Is(err, gorm.ErrRegistered)
}

// ErrInvalidField invalid field
func IsInvalidField(err error) bool {
	return errors.Is(err, gorm.ErrInvalidField)
}

// ErrEmptySlice empty slice found
func IsEmptySlice(err error) bool {
	return errors.Is(err, gorm.ErrEmptySlice)
}

// ErrDryRunModeUnsupported dry run mode unsupported
func IsDryRunModeUnsupported(err error) bool {
	return errors.Is(err, gorm.ErrDryRunModeUnsupported)
}

// ErrInvalidDB invalid db
func IsInvalidDB(err error) bool {
	return errors.Is(err, gorm.ErrInvalidDB)
}

// ErrInvalidValue invalid value
func IsInvalidValue(err error) bool {
	return errors.Is(err, gorm.ErrInvalidValue)
}

// ErrInvalidValueOfLength invalid values do not match length
func IsInvalidValueOfLength(err error) bool {
	return errors.Is(err, gorm.ErrInvalidValueOfLength)
}

// ErrPreloadNotAllowed preload is not allowed when count is used
func IsPreloadNotAllowed(err error) bool {
	return errors.Is(err, gorm.ErrPreloadNotAllowed)
}

// ErrDuplicatedKey occurs when there is a unique key constraint violation
func IsDuplicatedKey(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}

// ErrForeignKeyViolated occurs when there is a foreign key constraint violation
func IsForeignKeyViolated(err error) bool {
	return errors.Is(err, gorm.ErrForeignKeyViolated)
}

func HasAffectedRows(result *gorm.DB) bool {
	return result.RowsAffected > 0
}
