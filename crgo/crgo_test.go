package crgo

import (
	"testing"
	"regexp"
	"fmt"
	"net/http/httptest"
	"net/http"
)

const (
	REGEX_OK  = `^\[([\d,.]+s\])+([\w:\ ]+)(.*)(with\ )+(\[+\d+\])`
	REGEX_ERR = `^\[([\d,.]+s\])+([\w:\ ]+):(.*)$`
)

func TestRun_MustReturn_ErrorWithoutParameters(t *testing.T) {
	var emptyParams []string
	if Run(emptyParams) == nil {
		t.Errorf("Expected error without parameters.")
	}
}

func TestRun_MustReturn_StringWithElapsedTime(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test CRGO =)")
	}))
	defer ts.Close()

	err := Run([]string{ts.URL})
	if err != nil {
		t.Errorf("Expected no error but got %v.", err.Error())
	}
}

func TestRequest_MustReturn_StringWithElapsedTime(t *testing.T) {
	chn := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test CRGO =)")
	}))

	defer ts.Close()
	go Request(ts.URL, chn)

	matched, err := regexp.MatchString(REGEX_OK, <-chn)
	if err != nil {
		t.Error(err)
	}

	if !matched {
		t.Fail()
	}
}

func TestRequest_MustReturn_StringWithError(t *testing.T) {
	chn := make(chan string)
	url := "localhost"

	go Request(url, chn)
	matched, err := regexp.MatchString(REGEX_ERR, <-chn)
	if err != nil {
		t.Error(err)
	}

	if !matched {
		t.Fail()
	}
}
