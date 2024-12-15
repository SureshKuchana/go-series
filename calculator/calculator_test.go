package calculator_test

import (
	"_/home/sureshkoochana/tech/go-lang/go-series/calculator"
	"testing"
)

func TestAdd(t *testing.T){
	t.Parallel()
	type testCase struct{
		a,b float64
		want float64
	}
	testCases := []testCase{
		{a:2, b:2, want: 4},
		{a:4, b:2, want: 6},
		{a:5, b:0, want: 5},
		{a:10, b:10, want: 20},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a,tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}	
	}	
}

func TestSubtract(t *testing.T){
	t.Parallel()
	type testCase struct {
		a,b float64
		want float64
	}
	testCases := []testCase{
		{a:2, b:2, want: 0},
		{a:4, b:2, want: 2},
		{a:5, b:0, want: 5},
		{a:20, b:10, want: 10},
	}

	for _, tc := range testCases{
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}	
}

func TestMultiply(t *testing.T){
	t.Parallel()
	type testCase struct {
		a,b float64
		want float64
	}
	testCases := []testCase{
		{a:2, b:2, want: 4},
		{a:4, b:2, want: 8},
		{a:5, b:0, want: 0},
		{a:20, b:10, want: 200},
	}

	for _, tc := range testCases{
		got := calculator.Mutitply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Mutitply(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T){
	t.Parallel()
	type testCase struct {
		a,b float64
		want float64
	}
	testCases := []testCase{
		{a:2, b:2, want: 1},
		{a:4, b:2, want: 2},		
		{a:20, b:10, want: 2},
	}

	for _, tc := range testCases{
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}