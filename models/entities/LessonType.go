package entities

type LessonType struct {
	Name  string  `json:"name" gorm:"not null"`
}

var DefaultLessonTypes = []string {
	"lesson",
	"homework",
	"exam",
}
