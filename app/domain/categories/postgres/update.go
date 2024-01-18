package postgres

import (
	"fmt"
	"reflect"

	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"gorm.io/gorm"
)

func (p *PostgreSQL) Update(data *categories.Category, condition map[string]any) *gorm.DB {
	// return p.DB.Where(condition).Updates(data)
	query := p.DB
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

	return query.Updates(data)
}
