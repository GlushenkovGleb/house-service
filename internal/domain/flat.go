package domain

const (
	CreatedStatus      = "created"
	ApprovedStatus     = "approved"
	DeclinedStatus     = "declined"
	OnModerationStatus = "on_moderation"
)

var FlatStatuses = []FlatStatus{
	CreatedStatus,
	ApprovedStatus,
	DeclinedStatus,
	OnModerationStatus,
}

type FlatStatus string

func IsValidFlatStatus(status string) bool {
	for _, flatStatus := range FlatStatuses {
		if string(flatStatus) == status {
			return true
		}
	}
	return false
}

type Flat struct {
	ID      NullableInt64 `db:"id"`
	Number  int           `db:"number"`
	HouseID int           `db:"house_id"`
	Price   int           `db:"price"`
	Rooms   int           `db:"rooms"`
	Status  FlatStatus    `db:"status"`
}

var _ Insertable = (*Flat)(nil)

func (f *Flat) Args() (cols []string, vals []interface{}) {
	for col, val := range f.attrs() {
		cols = append(cols, col)
		vals = append(vals, val)
	}

	return
}

func (f *Flat) attrs() map[string]interface{} {
	return map[string]interface{}{
		"number":   f.Number,
		"house_id": f.HouseID,
		"price":    f.Price,
		"rooms":    f.Rooms,
	}
}
