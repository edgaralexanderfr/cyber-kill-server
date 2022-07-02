package input

import "testing"

var inputFactory InputFactoryInterface = &InputFactory{}
var input InputInterface = inputFactory.NewInput()

func TestInputAction(t *testing.T) {
	if input == nil {
		t.Fatalf(`1. inputFactory.NewInput() = nil, want InputInterface`)
	}

	if input.GetAction("MoveUp") {
		t.Fatalf(`2. input.GetAction("MoveUp") = true, want false`)
	}

	input.SetAction("MoveUp", true)

	if !input.GetAction("MoveUp") {
		t.Fatalf(`3. input.GetAction("MoveUp") = false, want true`)
	}

	input.SetAction("MoveUp", false)

	if input.GetAction("MoveUp") {
		t.Fatalf(`4. input.GetAction("MoveUp") = true, want false`)
	}

	input.SetAction("MoveUp", true)

	if input.GetAction("JUMP") {
		t.Fatalf(`5. input.GetAction("JUMP") = true, want false`)
	}

	input.SetAction("JUMP", true)

	if !input.GetAction("JUMP") {
		t.Fatalf(`6. input.GetAction("JUMP") = false, want true`)
	}

	if !input.GetAction("MoveUp") {
		t.Fatalf(`7. input.GetAction("MoveUp") = false, want true`)
	}
}

func TestInputValue(t *testing.T) {
	var x, y float32

	if input == nil {
		t.Fatalf(`1. inputFactory.NewInput() = nil, want InputInterface`)
	}

	x, y = input.GetValue("Aim")

	if x != 0.0 || y != 0.0 {
		t.Fatalf(`8. input.GetValue("Aim") = %f, %f; want 0.0, 0.0`, x, y)
	}

	input.SetValue("Aim", 5.0, -2.1)
	x, y = input.GetValue("Aim")

	if x != 5.0 || y != -2.1 {
		t.Fatalf(`9. input.GetValue("Aim") = %f, %f; want 5.0, -2.1`, x, y)
	}
}
