package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Student represents a student in the system
type Student struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Age            int                `json:"age" bson:"age"`
	Gender         string             `json:"gender" bson:"gender"`
	Address        Address            `json:"address" bson:"address"`
	Contact        string             `json:"contact" bson:"contact"`
	ParentContact  string             `json:"parentContact" bson:"parentContact"`
	EnrollmentDate primitive.DateTime `json:"enrollmentDate" bson:"enrollmentDate"`
	Courses        []CourseEnrollment `json:"courses" bson:"courses"`
	CreatedAt      primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt      primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

// Address represents the address of a student
type Address struct {
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	ZipCode string `json:"zipCode" bson:"zipCode"`
}

// CourseEnrollment represents a student's enrollment in a course
type CourseEnrollment struct {
	CourseID       primitive.ObjectID `json:"courseId" bson:"courseId"`
	EnrollmentDate primitive.DateTime `json:"enrollmentDate" bson:"enrollmentDate"`
	CompletionDate primitive.DateTime `json:"completionDate" bson:"completionDate"`
}
