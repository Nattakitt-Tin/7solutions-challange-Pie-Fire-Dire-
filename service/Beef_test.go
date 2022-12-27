package service_test

import (
	"PieFireDire/domain"
	"PieFireDire/service"
	"log"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCleanText(t *testing.T) {

	type testCase struct {
		name     string
		text     string
		expected string
	}

	cases := []testCase{
		{name: "have space", text: "Fatback t-bone t-bone", expected: "Fatback t-bone t-bone"},
		{name: "have comma and dot", text: "Fatback.t-bone,t-bone", expected: "Fatback t-bone t-bone"},
		{name: "have newline", text: "Fatback\nt-bone\nt-bone", expected: "Fatback t-bone t-bone"},
		{name: "have multi space", text: "Fatback  t-bone  t-bone", expected: "Fatback t-bone t-bone"},
		{name: "have all", text: "Fatback. \nt-bone,  t-bone", expected: "Fatback t-bone t-bone"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := service.NewBeef()
			result := b.CleanText(c.text)
			assert.Equal(t, c.expected, result)
		})
	}

}

func TestGetText(t *testing.T) {

	t.Run("test get text", func(t *testing.T) {
		ishaveText := true
		b := service.NewBeef()
		result, err := b.GetText()
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, ishaveText, len(result) > 0)
	})

}

func TestCountBeef(t *testing.T) {

	type testCase struct {
		name     string
		text     string
		expected domain.Beefs
	}

	cases := []testCase{
		{name: "have space", text: "Fatback t-bone t-bone", expected: domain.Beefs{Beef: map[string]uint{"fatback": 1, "t-bone": 2}}},
		{name: "have comma and dot", text: "Fatback.t-bone,t-bone", expected: domain.Beefs{Beef: map[string]uint{"fatback": 1, "t-bone": 2}}},
		{name: "have newline", text: "Fatback\nt-bone\nt-bone", expected: domain.Beefs{Beef: map[string]uint{"fatback": 1, "t-bone": 2}}},
		{name: "have multi space", text: "Fatback  t-bone  t-bone", expected: domain.Beefs{Beef: map[string]uint{"fatback": 1, "t-bone": 2}}},
		{name: "have all", text: "Fatback. \nt-bone,  t-bone", expected: domain.Beefs{Beef: map[string]uint{"fatback": 1, "t-bone": 2}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := service.NewBeef()
			clean := strings.Split(b.CleanText(c.text), " ")
			result := b.CountWord(clean)
			assert.Equal(t, c.expected.Beef, result.Beef)
		})
	}

}
