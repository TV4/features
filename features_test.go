package features

import "testing"

func TestNew(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		f := New(nil)

		if f.Enabled("foo") == true {
			t.Fatalf(`f.Enabled("foo") == true, want false`)
		}

		f.Enable("foo")

		if f.Enabled("foo") == false {
			t.Fatalf(`f.Enabled("foo") == false, want true`)
		}

		f.Disable("foo")

		if f.Enabled("foo") == true {
			t.Fatalf(`f.Enabled("foo") == true, want false`)
		}
	})

	t.Run("normal", func(t *testing.T) {
		f := New(map[string]bool{
			"foo": true,
			"bar": false,
		})

		if f.Enabled("foo") == false {
			t.Fatalf(`f.Enabled("foo") == false, want true`)
		}

		if f.Enabled("bar") == true {
			t.Fatalf(`f.Enabled("bar") == true, want false`)
		}

		if f.Enabled("baz") == true {
			t.Fatalf(`f.Enabled("baz") == true, want false`)
		}

		f.Enable("baz")

		if f.Enabled("baz") == false {
			t.Fatalf(`f.Enabled("baz") == false, want true`)
		}

		f.Disable("baz")

		if f.Enabled("baz") == true {
			t.Fatalf(`f.Enabled("baz") == true, want false`)
		}
	})
}
