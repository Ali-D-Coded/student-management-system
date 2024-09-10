package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Payment represents a payment transaction for a student's fee
type Payment struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID    primitive.ObjectID `json:"studentId" bson:"studentId"` // Reference to Student
	AmountPaid   float64            `json:"amountPaid" bson:"amountPaid"`
	PaymentDate  primitive.DateTime `json:"paymentDate" bson:"paymentDate"`
	PaymentMethod string            `json:"paymentMethod" bson:"paymentMethod"` // e.g., "Credit Card", "Bank Transfer", "Cash"
	ReceiptURL   string             `json:"receiptUrl" bson:"receiptUrl"`       // URL to the payment receipt document if needed
	CreatedAt    primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt    primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}
