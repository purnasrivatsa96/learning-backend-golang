package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/purnasrivatsa96/learning-backend-golang/tree/main/section-1_working-with-database-postgres-sqlc/simplebank/db/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {

	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
