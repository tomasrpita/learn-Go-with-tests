package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Tomás"},
			[]string{"Tomás"},
		},
		{
			"struct with tow string fields",
			struct {
				Name string
				City string
			}{"Tomás", "Madrid"},
			[]string{"Tomás", "Madrid"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Tomás", 46},
			[]string{"Tomás"},
		},
		{
			"nested fields",
			Person{
				"Tomás",
				Profile{46, "Madrid"},
			},
			[]string{"Tomás", "Madrid"},
		},
		{
			"pointers to things",
			&Person{
				"Tomás",
				Profile{46, "Madrid"},
			},
			[]string{"Tomás", "Madrid"},
		},
		// {
		// 	"slices",
		// 	[]Profile{
		// 		{33, "London"},
		// 		{46, "Madrid"},
		// 	},
		// 	[]string{"London", "Madrid"},
		// },
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
