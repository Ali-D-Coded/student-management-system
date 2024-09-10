package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Fee represents a fee structure for a student
type Fee struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID  primitive.ObjectID `json:"studentId" bson:"studentId"` // Reference to Student
	TotalAmount float64           `json:"totalAmount" bson:"totalAmount"`
	DueDate    primitive.DateTime `json:"dueDate" bson:"dueDate"`
	Status     string             `json:"status" bson:"status"` // "paid", "unpaid", "partially_paid"
	PaidOn     *primitive.DateTime `json:"paidOn,omitempty" bson:"paidOn,omitempty"` // Optional field
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt  primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}
