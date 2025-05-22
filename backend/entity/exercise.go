

package entity

import (

	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	ExerciseName string
	ExerciseType uint
	CaloriesBurndPerMinute int
}	