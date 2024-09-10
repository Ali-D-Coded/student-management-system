package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Attendance represents an attendance record for a student
type Attendance struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID primitive.ObjectID `json:"studentId" bson:"studentId"` // Reference to Student
	CourseID  primitive.ObjectID `json:"courseId" bson:"courseId"`   // Reference to Course
	Date      primitive.DateTime `json:"date" bson:"date"`
	Status    string             `json:"status" bson:"status"` // "present", "absent"
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}
