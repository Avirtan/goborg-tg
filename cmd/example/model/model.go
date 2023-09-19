package model

type ApiCollege struct {
	Date  string `json:"date"`
	Items []struct {
		Group        string `json:"group"`
		Lesson1Start string `json:"lesson1_start"`
		Lesson1End   string `json:"lesson1_end"`
		Lesson2Start string `json:"lesson2_start"`
		Lesson2End   string `json:"lesson2_end"`
		Discipline   string `json:"discipline"`
		Teacher      string `json:"teacher"`
		Cabinet      string `json:"cabinet"`
	} `json:"items"`
}

type Student struct {
	UserId    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Grout     string `db:"grp"`
	IsSub     bool   `db:"issub"`
}
