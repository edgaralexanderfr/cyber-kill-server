package physics

import (
	"testing"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/util"
)

const EPSILON = 0.1

func TestVector2GetDirectionTo1(t *testing.T) {
	var origin1 Vector2 = Vector2{150, 150}
	var target1 Vector2 = Vector2{256.1, 256.1}
	var result1 Vector2 = Vector2{42.4, 42.4}
	var speed1 float32 = 60

	resultInterface := origin1.GetDirectionTo(target1, speed1)

	if resultInterface == nil {
		t.Fatalf(`{150 150}.GetDirectionTo({256.1 256.1}, 60) = nil, want {42.4 42.4}`)
	}

	result := resultInterface.(Vector2)

	if !util.CompareEpsilon(result.X, result1.X, EPSILON) ||
		!util.CompareEpsilon(result.Y, result1.Y, EPSILON) {
		t.Fatalf(`{150 150}.GetDirectionTo({256.1 256.1}, 60) = {%f %f}, want {42.4 42.4}`, result.X, result.Y)
	}
}

func TestVector2GetDirectionTo2(t *testing.T) {
	var origin2 Vector2 = Vector2{380, 80}
	var target2 Vector2 = Vector2{341.8, 296.7}
	var result2 Vector2 = Vector2{-6.9, 39.4}
	var speed2 float32 = 40

	resultInterface := origin2.GetDirectionTo(target2, speed2)

	if resultInterface == nil {
		t.Fatalf(`{380 80}.GetDirectionTo({341.8 296.7}, 40) = nil, want {-6.9 39.4}`)
	}

	result := resultInterface.(Vector2)

	if !util.CompareEpsilon(result.X, result2.X, EPSILON) ||
		!util.CompareEpsilon(result.Y, result2.Y, EPSILON) {
		t.Fatalf(`{380 80}.GetDirectionTo({341.8 296.7}, 40) = {%f %f}, want {-6.9 39.4}`, result.X, result.Y)
	}
}

func TestVector2GetDirectionTo3(t *testing.T) {
	var origin3 Vector2 = Vector2{400, 300}
	var target3 Vector2 = Vector2{246.8, 171.4}
	var result3 Vector2 = Vector2{-15.3, -12.9}
	var speed3 float32 = 20

	resultInterface := origin3.GetDirectionTo(target3, speed3)

	if resultInterface == nil {
		t.Fatalf(`{400 300}.GetDirectionTo({246.8 171.4}, 20) = nil, want {-15.3 -12.9}`)
	}

	result := resultInterface.(Vector2)

	if !util.CompareEpsilon(result.X, result3.X, EPSILON) ||
		!util.CompareEpsilon(result.Y, result3.Y, EPSILON) {
		t.Fatalf(`{400 300}.GetDirectionTo({246.8 171.4}, 20) = {%f %f}, want {-15.3 -12.9}`, result.X, result.Y)
	}
}

func TestVector2GetDirectionTo4(t *testing.T) {
	var origin4 Vector2 = Vector2{120, 250}
	var target4 Vector2 = Vector2{170, 163.4}
	var result4 Vector2 = Vector2{100, -173.2}
	var speed4 float32 = 200

	resultInterface := origin4.GetDirectionTo(target4, speed4)

	if resultInterface == nil {
		t.Fatalf(`{120 250}.GetDirectionTo({170 163.4}, 200) = nil, want {100 -173.2}`)
	}

	result := resultInterface.(Vector2)

	if !util.CompareEpsilon(result.X, result4.X, EPSILON) ||
		!util.CompareEpsilon(result.Y, result4.Y, EPSILON) {
		t.Fatalf(`{120 250}.GetDirectionTo({170 163.4}, 200) = {%f %f}, want {100 -173.2}`, result.X, result.Y)
	}
}

func TestVector2GetDirectionTo5(t *testing.T) {
	var origin5 Vector2 = Vector2{200, 100}
	var target5 Vector2 = Vector2{250, 186.6}
	var result5 Vector2 = Vector2{0, 0}
	var speed5 float32 = 0

	resultInterface := origin5.GetDirectionTo(target5, speed5)

	if resultInterface == nil {
		t.Fatalf(`{200 100}.GetDirectionTo({250 186.6}, 0) = nil, want {0 0}`)
	}

	result := resultInterface.(Vector2)

	if !util.CompareEpsilon(result.X, result5.X, EPSILON) ||
		!util.CompareEpsilon(result.Y, result5.Y, EPSILON) {
		t.Fatalf(`{200 100}.GetDirectionTo({250 186.6}, 0) = {%f %f}, want {0 0}`, result.X, result.Y)
	}
}

func TestVector2GetDirectionTo6(t *testing.T) {
	var origin6 Vector2 = Vector2{240, 180}
	var target6 Vector2 = Vector2{240, 180}
	var result6 Vector2 = Vector2{0, 0}
	var speed6 float32 = 0

	resultInterface := origin6.GetDirectionTo(target6, speed6)

	if resultInterface == nil {
		t.Fatalf(`{240 180}.GetDirectionTo({240 180}, 0) = nil, want {0 0}`)
	}

	result := resultInterface.(Vector2)

	if !util.CompareEpsilon(result.X, result6.X, EPSILON) ||
		!util.CompareEpsilon(result.Y, result6.Y, EPSILON) {
		t.Fatalf(`{240 180}.GetDirectionTo({240 180}, 0) = {%f %f}, want {0 0}`, result.X, result.Y)
	}
}

func TestVector2GetDirectionTo7(t *testing.T) {
	var origin7 Vector2 = Vector2{240, 180}
	var target7 Vector2 = Vector2{240, 180}
	var result7 Vector2 = Vector2{0, 0}
	var speed7 float32 = 65

	resultInterface := origin7.GetDirectionTo(target7, speed7)

	if resultInterface == nil {
		t.Fatalf(`{240 180}.GetDirectionTo({240 180}, 65) = nil, want {0 0}`)
	}

	result := resultInterface.(Vector2)

	if !util.CompareEpsilon(result.X, result7.X, EPSILON) ||
		!util.CompareEpsilon(result.Y, result7.Y, EPSILON) {
		t.Fatalf(`{240 180}.GetDirectionTo({240 180}, 65) = {%f %f}, want {0 0}`, result.X, result.Y)
	}
}
