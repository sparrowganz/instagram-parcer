package instagram

import (
	"testing"
)

type caseData struct {
	data string
	err  interface{}
}

func TestGetAccountByUsername(t *testing.T) {
	cases := []caseData{
		{"sparrowganz", nil},
		{" ", InvalidStatusCode},
		{"testQEWEWR", InvalidStatusCode},
	}

	for _, testCase := range cases {
		_, err := GetAccountByUsername(testCase.data)
		if err != nil {
			if testCase.err == err.Error() {
				continue
			}
			t.Error(err)
		}
	}
}

func TestGetAccountByUrl(t *testing.T) {
	cases := []caseData{
		{"https://instagram.com/sparrowganz/", nil},
		{" ", ErrIsNotInstagram},
		{"https://vk.com", ErrIsNotInstagram},
	}
	for _, testCase := range cases {
		_, err := GetAccountByUrl(testCase.data)
		if err != nil {
			if testCase.err != err.Error() {
				t.Error(err)
			}
		}
	}
}

func TestGetAccountUrl(t *testing.T) {
	cases := []caseData{
		{"sparrowganz", "https://instagram.com/sparrowganz/"},
	}
	for _, testCase := range cases {
		url := GetAccountUrl(testCase.data)
		if url != testCase.err {
			t.Errorf("Url not equal")
		}
	}
}

func TestGetUsernameFromUrl(t *testing.T) {
	cases := []caseData{
		{"https://instagram.com/sparrowganz/", nil},
		{" ", ErrIsNotInstagram},
		{"https://vk.com", ErrIsNotInstagram},
	}
	for _, testCase := range cases {
		_, err := GetUsernameFromUrl(testCase.data)
		if err != nil {
			if err.Error() != testCase.err {
				t.Errorf("Url not equal")
			}
		}
	}
}

func TestGetLastMediasByUrl(t *testing.T) {
	cases := []caseData{
		{"https://instagram.com/sparrowganz/", nil},
		{" ", ErrIsNotInstagram},
		{"https://vk.com", ErrIsNotInstagram},
	}

	for _, testCase := range cases {
		_, err := GetLastMediasByUrl(testCase.data)
		if err != nil {
			if err.Error() != testCase.err {
				t.Errorf("Url not equal")
			}
		}
	}
}

func TestGetLastMediasByUsername(t *testing.T) {
	cases := []caseData{
		{"sparrowganz", nil},
		{" ", InvalidStatusCode},
		{"testQEWEWR", InvalidStatusCode},
	}
	for _, testCase := range cases {
		_, err := GetLastMediasByUsername(testCase.data)
		if err != nil {
			if err.Error() != testCase.err {
				t.Errorf("Url not equal")
			}
		}
	}
}

func TestGetMediaByShortCode(t *testing.T) {
	cases := []caseData{
		{"Bj4Ggzig8WP", nil},
		{" ", InvalidStatusCode},
		{"Bj-Hh6THT", InvalidStatusCode},
	}
	for _, testCase := range cases {
		_, err := GetMediaByShortCode(testCase.data)
		if err != nil {
			if testCase.err != err.Error() {
				t.Error(err)
			}
		}
	}
}

func TestGetMediaByUrl(t *testing.T) {
	cases := []caseData{
		{"https://instagram.com/p/Bj4Ggzig8WP", nil},
		{" ", ErrIsNotInstagram},
		{"https://instagram.com/p/Bj-Hh6THT", InvalidStatusCode},
	}
	for _, testCase := range cases {
		_, err := GetMediaByUrl(testCase.data)
		if err != nil {
			if testCase.err != err.Error() {
				t.Error(err)
			}
		}
	}
}
