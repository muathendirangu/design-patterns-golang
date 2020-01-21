package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected an in instance of singleton from GetInstance but got nil")
	}

	expectedCounter := counter1

	currentCounter := counter1.AddOne()

	if currentCounter != 1 {
		t.Errorf("Expected the counter value to be 1 but it is %d", currentCounter)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	currentCounter = counter2.AddOne()
	if currentCounter != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCounter)
	}
}
