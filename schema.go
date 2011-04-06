package main

import "strings"

type Page struct {
	Path    string
	Title   string
	Content string
	Products		[]Product
	PressReleases	[]PressRelease
}

type Product struct {
	Name            string
	Blurb           string
	FullDescription string
	ImagePath       string
}

type PressRelease struct {
	Date      string
	Title     string
	PathToPdf string
}

func (p *Page) IsMetrics() string {
	if strings.Contains(p.Path, "metrics"){
		return "selected"
	}
	return ""
}

func (p *Page) IsContact() string {
	if strings.Contains(p.Path, "contact"){
		return "selected"
	}
	return ""
}

func (p *Page) IsSupport() string {
	if strings.Contains(p.Path, "support"){
		return "selected"
	}
	return ""
}

func (p *Page) IsAboutPress() string {
	if strings.Contains(p.Path, "about"){
		return "selected"
	}
	if strings.Contains(p.Path,"press"){
		return "selected"
	}
	return ""
}

func (p *Page) IsProducts() string {
	if strings.Contains(p.Path, "products"){
		return "selected"
	}
	return ""
}