package main

import (
	"testing"
	"time"
)

type MockShopModel struct{
}

func (m MockShopModel) CountCustomers(t time.Time) (int, error){
	return 500, nil
}

func (m MockShopModel) CountSales(t time.Time) (int, error){
	return 500, nil
}

func TestCalculateSalesRate(t *testing.T) {
	msm := MockShopModel{}
	sr, err := calculateSalesRate(msm)
	if err != nil {
		t.Fatal(err)
	}
	exp := "1.00"
	if sr != exp {
		t.Fatalf("Expected : %v ---- Actual : %v \n",exp,sr)
	}
}