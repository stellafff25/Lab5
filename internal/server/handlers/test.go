package handlers

import (
	"context"
	"testing"

	db "github.com/stellafff25/Lab5/db/sqlc"
)

func TestCreateOrder_Unit(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	order, err := handler.CreateOrderHandler(context.Background(), db.CreateOrderParams{
		Name:   "TestOrder",
		Amount: 100,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.Name != "TestOrder" {
		t.Errorf("expected name 'TestOrder', got %v", order.Name)
	}

	if order.Amount != 100 {
		t.Errorf("expected amount 100, got %v", order.Amount)
	}
}

func TestGetOrder_Unit(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	order, err := handler.GetOrderHandler(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.ID != 1 {
		t.Errorf("expected id 1, got %v", order.ID)
	}

	if order.Name != "TestOrder" {
		t.Errorf("expected name 'TestOrder', got %v", order.Name)
	}
}

func TestUpdateOrder_Unit(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	order, err := handler.UpdateOrderHandler(context.Background(), db.UpdateOrderParams{
		ID:     1,
		Name:   "UpdatedOrder",
		Amount: 150,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.Name != "UpdatedOrder" {
		t.Errorf("expected name 'UpdatedOrder', got %v", order.Name)
	}

	if order.Amount != 150 {
		t.Errorf("expected amount 150, got %v", order.Amount)
	}
}

func TestDeleteOrder_Unit(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	err := handler.DeleteOrderHandler(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
