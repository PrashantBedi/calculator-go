package service

import(
	"github.com/jinzhu/gorm"
)

type Input struct {
	gorm.Model
	Number1 int	`json:"number1`
	Number2 int	`json:"number2`
	Result int
	Operation string
}