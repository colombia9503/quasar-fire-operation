package location

import "testing"

func TestGetLocationWithMoreThatOneValue(t *testing.T) {
	if _, _, err := GetLocation(100, 115.5, 142.7, 100, 100); err == nil {
		t.Error("Expected error")
	} else if err.Error() != "distances array out of bounds, expected exactly 3 distances" {
		t.Error("Unexpected error")
	}
}

func TestGetLocationWithLessThatThreeValues(t *testing.T) {
	if _, _, err := GetLocation(100, 115.5); err == nil {
		t.Error("Expected error")
	} else if err.Error() != "distances array out of bounds, expected exactly 3 distances" {
		t.Error("Unexpected error")
	}
}

func TestGetLocationOK(t *testing.T) {
	if x, y, err := GetLocation(100, 115.5, 142.7); err != nil {
		t.Error("Unexpected error")
	} else if int(x) != 301 && int(y) != 12140444 {
		t.Fatalf("Received %d %d instead of %d %d", int(x), int(y), 301, 12140444)
	}
}
