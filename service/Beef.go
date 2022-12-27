package service

import (
	"PieFireDire/domain"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type beef struct {
}

func NewBeef() beef {
	return beef{}
}

func (beef) CleanText(s string) string {
	regSt := regexp.MustCompile(`[.,\n]|[ ]{2,}`)
	regRd := regexp.MustCompile(`[ ]{2,}`)
	s = regSt.ReplaceAllString(s, " ")
	s = regRd.ReplaceAllString(s, " ")
	return s
}

func (beef) CountWord(sl []string) domain.Beefs {
	mapWord := make(map[string]uint)
	for _, w := range sl {
		if w != "" {
			_, ok := mapWord[strings.ToLower(w)]
			if ok {
				mapWord[strings.ToLower(w)] += 1
			} else {
				mapWord[strings.ToLower(w)] = 1
			}
		}
	}
	return domain.Beefs{Beef: mapWord}
}

func (beef) GetText() (string, error) {
	var client http.Client
	resp, err := client.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	}
	return "", errors.New("404")
}
