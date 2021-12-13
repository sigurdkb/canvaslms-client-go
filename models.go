package canvaslms

// Course -
type Course struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	CourseCode string  `json:"course_code"`
	Teachers   []User  `json:"teachers"`
	TAs        []User  `json:"tas"`
	Students   []User  `json:"students"`
	Groups     []Group `json:"groups"`
}

// User -
type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	LoginId string `json:"login_id"`
}

// Group -
type Group struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Students []User `json:"students"`
}
