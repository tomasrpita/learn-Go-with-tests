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
		{
			"slices",
			[]Profile{
				{33, "London"},
				{46, "Madrid"},
			},
			[]string{"London", "Madrid"},
		},
		{
			"slices",
			[2]Profile{
				{33, "London"},
				{46, "Madrid"},
			},
			[]string{"London", "Madrid"},
		},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Foz": "Bor",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Bor")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{46, "Madrid"}
			aChannel <- Profile{35, "Valladolid"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Madrid", "Valladolid"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{46, "Madrid"}, Profile{33, "Lisboa"}
		}

		var got []string
		want := []string{"Madrid", "Lisboa"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}

}
