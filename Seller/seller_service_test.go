package seller

import (
	"testing"
	"time"
)

func TestSellerService(t *testing.T) {
	seller1 := NewSeller(1, 10.0, 5, time.Now())
	seller2 := NewSeller(2, 8.0, 3, time.Now().Add(-time.Hour))
	seller3 := NewSeller(3, 12.0, 7, time.Now().Add(time.Hour))

	service := NewSellerService()

	// AddSeller tests
	service.AddSeller(seller1)
	if len(service.sellerMap) != 1 {
		t.Errorf("expected seller map length 1, got %d", len(service.sellerMap))
	}
	if _, ok := service.sellerMap[1]; !ok {
		t.Errorf("expected seller id 1 to be present in map")
	}

	// GetSeller tests
	result := service.GetSeller(1)
	if result != seller1 {
		t.Errorf("expected seller with id 1, got %#v", result)
	}
	result = service.GetSeller(99)
	if result != nil {
		t.Errorf("expected nil seller, got %#v", result)
	}

	// GetSellers tests
	sellers := service.GetSellers()
	if len(sellers) != 1 {
		t.Errorf("expected 1 seller, got %d", len(sellers))
	}
	service.AddSeller(seller2)
	service.AddSeller(seller3)
	sellers = service.GetSellers()
	if len(sellers) != 3 {
		t.Errorf("expected 3 sellers, got %d", len(sellers))
	}

	// GetSellerIds tests
	ids := service.GetSellerIds()
	if len(ids) != 3 {
		t.Errorf("expected 3 seller ids, got %d", len(ids))
	}
	if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
		t.Errorf("expected sorted ids 1, 2, 3, got %#v", ids)
	}

	// GetSellerIdsSorted tests
	ids = service.GetSellerIdsSorted()
	if len(ids) != 3 {
		t.Errorf("expected 3 seller ids, got %d", len(ids))
	}
	if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
		t.Errorf("expected sorted ids 1, 2, 3, got %#v", ids)
	}

	// GetSellerIdsSortedByTime tests
	ids = service.GetSellerIdsSortedByTime()
	if len(ids) != 3 {
		t.Errorf("expected 3 seller ids, got %d", len(ids))
	}
	if ids[0] != 2 || ids[1] != 1 || ids[2] != 3 {
		t.Errorf("expected sorted ids by time 2, 1, 3, got %#v", ids)
	}
}
