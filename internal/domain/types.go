package domain

import (
	"database/sql"
	"time"
)

type Insertable interface {
	Args() (cols []string, vals []interface{})
}

type NullableString struct {
	sql.NullString
}

func AutoString(s string) NullableString {
	if s != "" {
		return NullableString{
			sql.NullString{String: s, Valid: true},
		}
	}
	return NullableString{}
}

type NullableInt64 struct {
	sql.NullInt64
}

func AutoInt64(i int64) NullableInt64 {
	if i != 0 {
		return NullableInt64{
			sql.NullInt64{Int64: i, Valid: true},
		}
	}
	return NullableInt64{}
}

type NullableTime struct {
	sql.NullTime
}

func AutoTime(t time.Time) NullableTime {
	if !t.IsZero() {
		return NullableTime{
			sql.NullTime{Time: t, Valid: true},
		}
	}
	return NullableTime{}
}
