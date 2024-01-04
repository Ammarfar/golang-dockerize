package models

type Mark struct {
	Id     *uint  `json:"id"`
	UserId *int64 `json:"user_id"`
	Mark   *int   `json:"mark"`
}
