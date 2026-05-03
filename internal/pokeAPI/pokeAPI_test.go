package pokeAPI

import "testing"

func TestCommandMapIncreasesOffset(t *testing.T) {
	mapOffset = 0
	mapLimit = 20

	err := CommandMap()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mapOffset != 20 {
		t.Fatalf("expected mapOffset to be 20, got %d", mapOffset)
	}
}

func TestCommandMapbFromSecondPageGoesBackToFirstPage(t *testing.T) {
	mapOffset = 40
	mapLimit = 20

	err := CommandMapb()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mapOffset != 20 {
		t.Fatalf("expected mapOffset to be 20 after going back, got %d", mapOffset)
	}
}

func TestCommandMapbDoesNotGoBelowZero(t *testing.T) {
	mapOffset = 20
	mapLimit = 20

	err := CommandMapb()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mapOffset != 20 {
		t.Fatalf("expected mapOffset to be 20 after fetching first page, got %d", mapOffset)
	}
}