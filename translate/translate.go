package translate

import (
	xdata "github.com/mt1976/appFrame/dataloader"
	xlogs "github.com/mt1976/appFrame/logs"
	"golang.org/x/exp/slices"
)

// var translations xdata.Payload
//var locales []string
//var l xlogs.XLogger

type Translator struct {
	translations xdata.Payload
	locales      []string
	logger       xlogs.XLogger
}

func init() {
	//	translator.translations = New()
	//	locales = GetLocales()
	//	l = xlogs.New()
}

func New() *Translator {
	rtn := &Translator{}
	rtn.translations = *xdata.New("translate", "dat", "")
	rtn.locales = getLocales()
	rtn.logger = xlogs.New()
	return rtn
}

func (T *Translator) Get(property string) string {
	return T.translations.Get(property)
}

func (T *Translator) GetLocalised(property string, language string) string {
	if language == "" {
		language = "en"
	}
	// if language not in locales, warn and return default
	if !slices.Contains(T.GetLocales(), language) {
		T.logger.WithField("Language", language).Warn("Language not found in 'safe' locale list, using default")
		return T.Get(property)
	}
	rtn, _ := T.translations.GetStringofCategory(property, language)
	return rtn
}

func (T *Translator) GetLocales() []string {
	return T.locales
}

func getLocales() []string {
	x := []string{"en"}
	return x
}

func (T *Translator) AddLocale(locale string) {
	T.locales = append(T.locales, locale)
}

func (T *Translator) Verbose() {
	T.translations.Verbose()
}

func (T *Translator) Normal() {
	T.translations.Normal()
}
