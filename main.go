package main

import (
	"flag"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

func main() {
	var count int
	var lang string
	var name string
	flag.IntVar(&count, "count", 0, "number of items to buy")
	flag.StringVar(&lang, "lang", "en", "for accept language")
	flag.StringVar(&name, "name", "One", "for name parameter")
	flag.Parse()
	fmt.Printf("You're buying %d cookies\n", count)
	// new localizer

	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("active.en.toml")
	bundle.LoadMessageFile("active.zh.toml")

	//generate file
	localizer := i18n.NewLocalizer(bundle, lang)
	localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "PersonCats",
			One:   "{{.Name}} has {{.PluralCount}} cat.",
			Other: "{{.Name}} has {{.PluralCount}} cats.",
		},
	})

	// print localizer
	{
		localizer := i18n.NewLocalizer(bundle, lang)
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID:   "PersonCats",
			PluralCount: count,
			DefaultMessage: &i18n.Message{
				ID:    "PersonCats",
				Hash:  "sha1-e0aace783982e635707d2e4da79d4962ead0ade3",
				One:   "Smith has {{.PluralCount}} cat.",
				Other: "Smith has {{.PluralCount}} cats.",
			},
		}))
	}

}
