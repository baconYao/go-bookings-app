package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	// 驗證 POST 請求沒有帶 required payload 的 case
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	// 驗證 POST 請求有帶 required payload 的 case
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	// 驗證沒有 field 的 form
	postedData := url.Values{}
	form := New(postedData)
	if form.Has("yao") {
		t.Error("form shows has field when it doesn't")
	}

	// 驗證有 field 的 form
	postedData = url.Values{}
	postedData.Add("yao", "hi")
	postedData.Add("judy", "cute")

	// r.PostForm = postedData
	form = New(postedData)
	if !form.Has("judy") || !form.Has("yao") {
		t.Error("shows form doesn't have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	form.MinLength("yao", 3)

	if form.Valid() {
		t.Error("form shows min lenght for non-existent field")
	}

	isError := form.Errors.Get("yao")
	if isError == "" {
		t.Error("should have error but did not get one")
	}

	postedValues = url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)
	form.MinLength("some_field", 100)

	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)
	form.MinLength("another_field", 1)

	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	form.IsEmail("aaa")

	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "abc123@hi.com")
	form = New(postedValues)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "lalala")
	form = New(postedValues)
	form.IsEmail("email")

	if form.Valid() {
		t.Error("got valid for invalid email address")
	}
}
