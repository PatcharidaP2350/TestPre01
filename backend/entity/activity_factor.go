
package entity

import (
	"gorm.io/gorm"

)

type ActivityFactor struct {
	gorm.Model
	ActivityLevel string  //ระดับการออกกำลังกาย เช่น "กิจกรรมเบา", "สูงมาก"
	Duration int     //ช่วงเวลาที่ออกกำลังกายในแต่ละวัน (เช่น 15-30 นาที)
	Frequency int     //ความถี่ต่อสัปดาห์ (เช่น 3-4 ครั้ง/สัปดาห์)
	EstimatedCalories int     //ปริมาณแคลอรี่ที่คาดว่าจะออกกำลังกาย, ค่าพลังงานที่ใช้ประมาณต่อวัน (kcal)
	
}