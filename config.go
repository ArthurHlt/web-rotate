package main

import (
	"fmt"
	"time"
)

// vertical-align: middle;text-align: center;border: 1px solid black;
var defaultBoxStyle = Style{
	"position":       "absolute",
	"top":            0,
	"left":           0,
	"font-size":      "xx-large",
	"border-radius":  "5px",
	"margin":         "10px",
	"width":          "30%",
	"height":         "120px",
	"z-index":        50000,
	"opacity":        0.7,
	"background":     "white",
	"line-height":    "120px",
	"vertical-align": "middle",
	"text-align":     "center",
	"border":         "1px solid black",
}

var defaultTextStyle = Style{
	"color":   "blue",
	"opacity": 1,
}

type Config struct {
	Pages      []*Page `yaml:"pages"`
	Fullscreen bool    `yaml:"fullscreen"`
}

type StyleCartridge struct {
	Box  Style `yaml:"box"`
	Text Style `yaml:"text"`
}

func (c *StyleCartridge) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain StyleCartridge
	var err error
	if err = unmarshal((*plain)(c)); err != nil {
		return err
	}
	c.Text = mergeMap(defaultTextStyle, c.Text)
	c.Box = mergeMap(defaultBoxStyle, c.Box)
	return nil
}

type Style map[string]interface{}

func (s Style) ToCss() string {
	css := ""
	for k, v := range s {
		css += fmt.Sprintf("%s: %s;", k, fmt.Sprint(v))
	}
	return css
}

type Page struct {
	Cartridge      string                 `yaml:"cartridge"`
	Url            string                 `yaml:"url"`
	Duration       Duration               `yaml:"duration"`
	StyleCartridge *StyleCartridge        `yaml:"style_cartridge,omitempty"`
	Headers        map[string]interface{} `yaml:"headers,omitempty"`
}

func (c *Page) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Page
	var err error
	if err = unmarshal((*plain)(c)); err != nil {
		return err
	}
	if c.Cartridge == "" {
		return fmt.Errorf("You must define cartridge in your page")
	}
	if c.Url == "" {
		return fmt.Errorf("You must define an url in your page")
	}
	if c.StyleCartridge == nil {
		c.StyleCartridge = &StyleCartridge{
			Box:  defaultBoxStyle,
			Text: defaultTextStyle,
		}
	}
	return nil
}

type Duration time.Duration

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	var err error
	if err = unmarshal(&s); err != nil {
		return err
	}
	if s == "" {
		s = "20s"
	}
	duration, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(duration)

	return nil
}
