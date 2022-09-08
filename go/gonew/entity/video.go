package entity

type Video struct {
	Id          int    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string `json:"title" gorm:type:varchar(100)`
	Description string `json:"description" gorm:type:varchar(255)`
	URL         string `json:"url" gorm:type:varchar(255)`
	Author      Person `json:"author" gorm:"foreignkey:PersonID`
	PersionID   int    `json:"-"`
}

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
