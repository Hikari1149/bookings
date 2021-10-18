package forms

import (
	"net/url"
	"testing"
)

func TestForm_valid(t *testing.T) {
	postData := url.Values{}

	form := New(postData)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	postData := url.Values{}

	form := New(postData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error ,but did not get one")
	}

	postData = url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "a")
	postData.Add("c", "a")

	form = New(postData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error ,but did get one")
	}
}

func TestForm_MinLength(t *testing.T) {

	postData := url.Values{}
	form := New(postData)
	form.MinLength("a", 3)
	if form.Valid() {
		t.Error("pass min length when it does not")
	}

	postData = url.Values{}
	postData.Add("a", "123123")
	form = New(postData)
	form.MinLength("a", 3)
	if !form.Valid() {
		t.Error("not pass min length when it should")
	}
}

func TestForm_Has(t *testing.T) {

	postData := url.Values{}
	form := New(postData)

	form.Has("a")
	if form.Valid() {
		t.Error("pass has when it does not")
	}

	postData = url.Values{}
	postData.Add("a", "123")
	form = New(postData)
	form.Has("a")
	if !form.Valid() {
		t.Error("not pass has when it should")
	}

}

func TestForm_IsEmail(t *testing.T) {

	postData := url.Values{}
	form := New(postData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("pass isEmail when it does not")
	}

	postData = url.Values{}
	postData.Add("a", "123@qq.com")
	form = New(postData)
	form.IsEmail("a")
	if !form.Valid() {
		t.Error("not pass isEmail when it should")
	}

}
