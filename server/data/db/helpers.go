package db

import "github.com/mbaraa/apollo-music/models"

// AllowedModels defines allowed models to be used in the db
type AllowedModels interface {
	models.User | models.Verification | models.Subscription
	GetId() uint
}

func InitTables() {
	if instance != nil {
		err := instance.AutoMigrate(
			new(models.User),
			new(models.Verification),
			new(models.Subscription),
		)
		if err != nil {
			panic(err)
		}
	} else {
		panic("No DB connection was initialized")
	}
}

// checkConds reports whether the provided conditions are valid or not
func checkConds(conds ...any) bool {
	return len(conds) > 1 && checkCondsMeaning(conds...)
}

func checkCondsMeaning(conds ...any) bool {
	ok := false

	switch conds[0].(type) {
	case string:
		ok = true
	default:
		return false
	}

	for _, cond := range conds[1:] {
		switch cond.(type) {
		case bool,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128,
			string:
			ok = true
		default:
			return false
		}
	}

	return ok
}
