package gormhelper

import (
	"errors"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

// ErrRecordNotFound record not found error
func IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsErrNotNilNotRecordNotFound(err error) bool {
	isError := err != nil && !errors.Is(err, gorm.ErrRecordNotFound)
	return isError
}

// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
func IsErrInvalidTransaction(err error) bool {
	return errors.Is(err, gorm.ErrInvalidTransaction)
}

// ErrNotImplemented not implemented
func IsErrNotImplemented(err error) bool {
	return errors.Is(err, gorm.ErrNotImplemented)
}

// ErrMissingWhereClause missing where clause
func IsErrMissingWhereClause(err error) bool {
	return errors.Is(err, gorm.ErrMissingWhereClause)
}

// ErrUnsupportedRelation unsupported relations
func IsErrUnsupportedRelation(err error) bool {
	return errors.Is(err, gorm.ErrUnsupportedRelation)
}

// ErrPrimaryKeyRequired primary keys required
func IsErrPrimaryKeyRequired(err error) bool {
	return errors.Is(err, gorm.ErrPrimaryKeyRequired)
}

// ErrModelValueRequired model value required
func IsErrModelValueRequired(err error) bool {
	return errors.Is(err, gorm.ErrModelValueRequired)
}

// ErrModelAccessibleFieldsRequired model accessible fields required
func IsErrModelAccessibleFieldsRequired(err error) bool {
	return errors.Is(err, gorm.ErrModelAccessibleFieldsRequired)
}

// ErrSubQueryRequired sub query required
func IsErrSubQueryRequired(err error) bool {
	return errors.Is(err, gorm.ErrSubQueryRequired)
}

// ErrInvalidData unsupported data
func IsErrInvalidData(err error) bool {
	return errors.Is(err, gorm.ErrInvalidData)
}

// ErrUnsupportedDriver unsupported driver
func IsErrUnsupportedDriver(err error) bool {
	return errors.Is(err, gorm.ErrUnsupportedDriver)
}

// ErrRegistered registered
func IsErrRegistered(err error) bool {
	return errors.Is(err, gorm.ErrRegistered)
}

// ErrInvalidField invalid field
func IsErrInvalidField(err error) bool {
	return errors.Is(err, gorm.ErrInvalidField)
}

// ErrEmptySlice empty slice found
func IsErrEmptySlice(err error) bool {
	return errors.Is(err, gorm.ErrEmptySlice)
}

// ErrDryRunModeUnsupported dry run mode unsupported
func IsErrDryRunModeUnsupported(err error) bool {
	return errors.Is(err, gorm.ErrDryRunModeUnsupported)
}

// ErrInvalidDB invalid db
func IsErrInvalidDB(err error) bool {
	return errors.Is(err, gorm.ErrInvalidDB)
}

// ErrInvalidValue invalid value
func IsErrInvalidValue(err error) bool {
	return errors.Is(err, gorm.ErrInvalidValue)
}

// ErrInvalidValueOfLength invalid values do not match length
func IsErrInvalidValueOfLength(err error) bool {
	return errors.Is(err, gorm.ErrInvalidValueOfLength)
}

// ErrPreloadNotAllowed preload is not allowed when count is used
func IsErrPreloadNotAllowed(err error) bool {
	return errors.Is(err, gorm.ErrPreloadNotAllowed)
}

// ErrDuplicatedKey occurs when there is a unique key constraint violation
func IsErrDuplicatedKey(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey)
}

// ErrForeignKeyViolated occurs when there is a foreign key constraint violation
func IsErrForeignKeyViolated(err error) bool {
	return errors.Is(err, gorm.ErrForeignKeyViolated)
}

func HasAffectedRows(result *gorm.DB) bool {
	return result.RowsAffected > 0
}

func ConditionMapping(query *gorm.DB, condition map[string]any) *gorm.DB {
	for field, value := range condition {
		switch reflect.ValueOf(value).Kind() {
		case reflect.Slice:
			{
				slCondition := value.([]any)
				if len(slCondition) == 2 {
					query = query.Where(fmt.Sprintf("%s %s ? ", field, slCondition[0]), slCondition[1])
				} else {
					return &gorm.DB{
						Error: fmt.Errorf("invalid condition"),
					}
				}
			}
		default:
			{
				query = query.Where(fmt.Sprintf("%s = ?", field), value)
			}
		}
	}

	return query
}
