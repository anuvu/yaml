package yaml

import (
	"github.com/gosexy/to"
	"testing"
)

func TestRead(t *testing.T) {
	_, err := Open("examples/input/settings.yaml")
	if err != nil {
		t.Errorf("Test failed.")
	}
}

func TestGet(t *testing.T) {
	settings, err := Open("examples/input/settings.yaml")

	if err != nil {
		t.Errorf("Test failed.")
	}

	test1 := "Hello World!"
	val1 := to.String(settings.Get("test_string"))

	if val1 != test1 {
		t.Errorf("Got %t expecting %t.", val1, test1)
	}

	val2 := to.Int(settings.Get("non_defined_int"))

	if val2 != 0 {
		t.Errorf("Test failed.")
	}

	test3 := "Third"

	val3 := to.List(settings.Get("test_map/element_3/test_sequence"))

	if val3[2] != test3 {
		t.Errorf("Got %t expecting %t.", val3[2], test3)
	}

	test4 := "HaS CaSe"
	val4 := to.String(settings.Get("test_caseless"))

	if val4 != test4 {
		t.Errorf("Got %t expecting %t.", val4, test4)
	}

	test5 := "Hello World!"
	val5 := to.String(settings.Get("test_string"))

	if test5 != val5 {
		t.Errorf("Got %t expecting %t.", test5, val5)
	}

	test6 := 1234
	val6 := to.Int(settings.Get("test_int"))

	if test6 != val6 {
		t.Errorf("Got %t expecting %t.", test6, val6)
	}

	test7 := float64(1.2)
	val7 := to.Float64(settings.Get("test_float"))

	if test7 != val7 {
		t.Errorf("Got %t expecting %t.", test7, val7)
	}

	test8 := true
	val8 := to.Bool(settings.Get("test_bool"))

	if test8 != val8 {
		t.Errorf("Got %t expecting %t.", test8, val8)
	}

}

func TestSet(t *testing.T) {
	settings, err := Open("examples/input/settings.yaml")

	if err != nil {
		t.Errorf("Test failed.")
	}

	settings.Set("test_map/element_3/test_bool", true)

	test1 := true
	val1 := to.Bool(settings.Get("test_map/element_3/test_bool"))

	if val1 != test1 {
		t.Errorf("Got %t expecting %t.", val1, test1)
	}

}

func TestWrite(t *testing.T) {
	settings := New()

	settings.Set("test_map/element_3/test_bool", true)

	err := settings.Write("examples/input/settings2.yaml")

	if err != nil {
		t.Errorf("Test failed.")
	}
}
