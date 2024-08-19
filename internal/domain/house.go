package domain

type House struct {
	ID        NullableInt64  `db:"id"`
	Address   string         `db:"address"`
	Year      int            `db:"year"`
	Developer NullableString `db:"developer"`
	CreatedAt NullableTime   `db:"created_at"`
	UpdatedAt NullableTime   `db:"updated_at"`
}

var _ Insertable = (*House)(nil)

func (h *House) Args() (cols []string, vals []interface{}) {
	for col, val := range h.attrs() {
		vals = append(vals, val)
		cols = append(cols, col)
	}
	return
}

func (h *House) attrs() map[string]interface{} {
	return map[string]interface{}{
		"address":   h.Address,
		"year":      h.Year,
		"developer": h.Developer,
	}
}
