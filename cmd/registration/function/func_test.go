package function

import "testing"

func TestId(t *testing.T) {

	total := Id(12)

	if total != 12 {
		t.Errorf("Error")
		
	}

}
func TestName(t *testing.T) {

	total := Name("rafael")

	if total != "rafael"{
		t.Errorf("Error")
	}

}

func TestTime(t *testing.T){

	total := Time(7.35)

	if total != 7.35 {

		t.Errorf("Error")

	}

}