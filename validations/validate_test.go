package validations

import (
	"reflect"
	"testing"
)

func TestPassword(t *testing.T) {
	data := struct {
		Password string
	}{Password: "123Abc78"}
	t.Run("invalid", func(t *testing.T) {
		errs := map[string]string{}
		ok := New(data, Fields{
			"Password": Rules(Password),
		}).Validate(errs)
		assertFalse(t, ok)
		asserteq(t, 1, len(errs))
	})
	t.Run("valid", func(t *testing.T) {
		errs := map[string]string{}
		data.Password = "123Abc78."
		ok := New(data, Fields{
			"Password": Rules(Password),
		}).Validate(errs)
		assertTrue(t, ok)
		asserteq(t, 0, len(errs))
	})
}

func TestCustomMessage(t *testing.T) {
	data := struct {
		Name string
	}{Name: ""}
	errs := map[string]string{}
	ok := New(data, Fields{
		"Name": Rules(Required, Message("custom name")),
	}).Validate(errs)
	assertFalse(t, ok)
	asserteq(t, 1, len(errs))
	asserteq(t, "custom name", errs["Name"])
}

func TestRequired(t *testing.T) {
	data := struct {
		Name string
	}{Name: ""}
	t.Run("invalid", func(t *testing.T) {
		errs := map[string]string{}
		ok := New(data, Fields{
			"Name": Rules(Required),
		}).Validate(errs)
		assertFalse(t, ok)
		asserteq(t, 1, len(errs))
	})
	t.Run("valid", func(t *testing.T) {
		errs := map[string]string{}
		data.Name = "foo"
		ok := New(data, Fields{
			"Name": Rules(Required),
		}).Validate(errs)
		assertTrue(t, ok)
		asserteq(t, 0, len(errs))
	})
}

func TestUrl(t *testing.T) {
	data := struct {
		Url string
	}{Url: "http://foocom"}
	t.Run("invalid", func(t *testing.T) {
		errs := map[string]string{}
		ok := New(data, Fields{
			"Url": Rules(Url),
		}).Validate(errs)
		assertFalse(t, ok)
		asserteq(t, 1, len(errs))
	})
	t.Run("valid", func(t *testing.T) {
		errs := map[string]string{}
		data.Url = "http://foo.com"
		ok := New(data, Fields{
			"Url": Rules(Url),
		}).Validate(errs)
		assertTrue(t, ok)
		asserteq(t, 0, len(errs))
	})
}

func TestEmail(t *testing.T) {
	data := struct {
		Email string
	}{Email: "foo.com"}
	t.Run("invalid", func(t *testing.T) {
		errs := map[string]string{}
		ok := New(data, Fields{
			"Email": Rules(Email),
		}).Validate(errs)
		assertFalse(t, ok)
		asserteq(t, 1, len(errs))
	})
	t.Run("valid", func(t *testing.T) {
		errs := map[string]string{}
		data.Email = "foo@test.com"
		ok := New(data, Fields{
			"Email": Rules(Email),
		}).Validate(errs)
		assertTrue(t, ok)
		asserteq(t, 0, len(errs))
	})
}

func TestEqual(t *testing.T) {
	data := struct {
		Equal string
	}{Equal: "foobarbaz"}
	t.Run("invalid", func(t *testing.T) {
		errs := map[string]string{}
		ok := New(data, Fields{
			"Equal": Rules(Equal("bar")),
		}).Validate(errs)
		assertFalse(t, ok)
		asserteq(t, 1, len(errs))
	})
	t.Run("valid", func(t *testing.T) {
		errs := map[string]string{}
		data.Equal = "foobarbaz"
		ok := New(data, Fields{
			"Equal": Rules(Equal("foobarbaz")),
		}).Validate(errs)
		assertTrue(t, ok)
		asserteq(t, 0, len(errs))
	})
}

func TestMax(t *testing.T) {
	data := struct {
		Name string
	}{Name: "123456789"}
	t.Run("invalid", func(t *testing.T) {
		errs := map[string]string{}
		ok := New(data, Fields{
			"Name": Rules(Max(5)),
		}).Validate(errs)
		assertFalse(t, ok)
		asserteq(t, 1, len(errs))
	})
	t.Run("valid", func(t *testing.T) {
		errs := map[string]string{}
		foo := struct {
			Name string
		}{Name: "123"}
		ok := New(foo, Fields{
			"Name": Rules(Max(5)),
		}).Validate(errs)
		assertTrue(t, ok)
		asserteq(t, 0, len(errs))
	})
}

func TestMin(t *testing.T) {
	data := struct {
		Name string
	}{Name: "1234"}
	t.Run("invalid", func(t *testing.T) {
		errs := map[string]string{}
		ok := New(data, Fields{
			"Name": Rules(Min(5)),
		}).Validate(errs)
		assertFalse(t, ok)
		asserteq(t, 1, len(errs))
	})
	t.Run("valid", func(t *testing.T) {
		errs := map[string]string{}
		data.Name = "123456789"
		ok := New(data, Fields{
			"Name": Rules(Min(5)),
		}).Validate(errs)
		assertTrue(t, ok)
		asserteq(t, 0, len(errs))
	})
}

func assertTrue(t *testing.T, con bool) {
	if !con {
		t.Fatalf("expected true")
	}
}

func assertFalse(t *testing.T, con bool) {
	if con {
		t.Fatalf("expected false")
	}
}

func asserteq(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("expected %v to equal %v", a, b)
	}
}
