// Package codename generates a random code name meant for naming software releases
// if you run short of inspiration.
//
// Currently based on the pattern "[Superb] [Superhero]".
//
// The words are included in the binary with go-bindata, to regenerate
// the dictionary, run in the codename folder:
//   $ go get -u github.com/jteeuwen/go-bindata/...
//   $ go-bindata -o words.go -pkg codename data/
package codename

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	// SpacedString returns "Awesome Hero".
	SpacedString FormatType = 1
	// Sanitized returns a safe string, "awesome-hero".
	Sanitized FormatType = 2

	SuperbFilePath      = "data/superb.json"
	SuperheroesFilePath = "data/superheroes.json"
)

type FormatType uint
type JSONData []interface{}

// GetCodename generates a code name with the given format.
func Get(format FormatType) (string, error) {
	// Load Superb words
	sp, err := Asset(SuperbFilePath)
	if err != nil {
		return "", err
	}

	var spa JSONData
	json.Unmarshal(sp, &spa)

	// Load Superheroes words
	sh, err := Asset(SuperheroesFilePath)
	if err != nil {
		return "", err
	}
	var sha JSONData
	json.Unmarshal(sh, &sha)

	// Seed the randomizer
	rand.Seed(time.Now().UTC().UnixNano())

	// Pick random items from each list
	fw := spa[rand.Intn(len(spa))]
	sw := sha[rand.Intn(len(sha))]

	codename := fmt.Sprintf("%s %s", fw, sw)

	switch format {
	case SpacedString:
		upperCaseFirst := func(s string) string {
			if len(s) == 0 {
				return ""
			}
			r, n := utf8.DecodeRuneInString(s)
			return string(unicode.ToUpper(r)) + s[n:]
		}

		cs, csn := strings.Split(codename, " "), []string{}
		for _, csi := range cs {
			csn = append(csn, upperCaseFirst(csi))
		}
		codename = strings.Join(csn, " ")
		break
	case Sanitized:
		codename = strings.ToLower(strings.TrimSpace(codename))
		codename = strings.Replace(codename, " ", "-", -1)
		reg, err := regexp.Compile("[^a-z0-9-]+")
		if err != nil {
			return "", err
		}
		codename = reg.ReplaceAllString(codename, "")
		break
	}

	return codename, nil
}
