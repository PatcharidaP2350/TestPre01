package entity

import (
	"gorm.io/gorm"
	"time"
)
type ExerciseActivity struct {   //กิจกรรมที่ User ออก
	gorm.Model
	ActivityName string    //ชื่อกิจกรรมการออกกำลังกาย
	ExerciseID   uint 
	UserID uint
	Date time.Time          //วันที่
	Duration int  //เวลาที่ User ออกกำลังกาย
	CaloriesBurnd int   //(คำนวณ = calories_burned_per_minute * duration) แคลที่เผา
}
