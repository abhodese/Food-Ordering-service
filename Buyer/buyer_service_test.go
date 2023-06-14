package buyer

import (
	"net/http"
	"testing"
)

func TestSearchProducts(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SearchProducts(tt.args.w, tt.args.r)
		})
	}
}


func TestPlaceOrder(t *testing.T){
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PlaceOrder(tt.args.w, tt.args.r)
		})
	}
}


