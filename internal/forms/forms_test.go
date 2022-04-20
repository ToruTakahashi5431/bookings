package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/example.com", nil)

	// toru ↓これでもいける
	// r, _ := http.NewRequest("POST", "/example.com", nil)

	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/example.com", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form show valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shwos does not have required field when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it dose not")
	}

	postedData = url.Values{}
	postedData.Add("toru", "takahashi")
	form = New(postedData)

	has = form.Has("toru")
	if !has {
		t.Error("showa from does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("toru", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("toru")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValue := url.Values{}
	postedValue.Add("some_field", "some value")
	form = New(postedValue)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("show minlength of 100 when data shorter")
	}

	postedValue = url.Values{}
	postedValue.Add("another_field", "torutorutoru")
	form = New(postedValue)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("show minlength of 1 not met when it is")
	}

	isError = form.Errors.Get("toru")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValue := url.Values{}
	form := New(postedValue)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValue = url.Values{}
	postedValue.Add("email", "torunomichi5431@gmail.com")
	form = New(postedValue)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedValue = url.Values{}
	postedValue.Add("email", "toru")
	form = New(postedValue)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email address")
	}
}
