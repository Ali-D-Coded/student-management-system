package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Course represents a course in the system
type Course struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CourseName  string             `json:"courseName" bson:"courseName"`
	Description string             `json:"description" bson:"description"`
	DurationInWeeks int            `json:"durationInWeeks" bson:"durationInWeeks"`
	TeacherID   primitive.ObjectID `json:"teacherId" bson:"teacherId"` // Reference to User (Teacher)
	CreatedAt   primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt   primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}
