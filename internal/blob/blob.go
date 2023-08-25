package blob

import (
	"bytes"
	"embed"
	htmltmpl "html/template"
	texttmpl "text/template"
)

// content holds our static web server content.
//
//go:embed template/*
var content embed.FS

type RenderArtifact struct {
	HTML string
	Text string
}

func RenderHTML(filename string, data any) (string, error) {
	t, err := htmltmpl.ParseFS(content, "template/"+filename+".html.tmpl")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func RenderText(filename string, data any) (string, error) {
	t, err := texttmpl.ParseFS(content, "template/"+filename+".txt.tmpl")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func RenderPair(name string, data any) (*RenderArtifact, error) {
	html, err := RenderHTML(name, data)
	if err != nil {
		return nil, err
	}

	text, err := RenderText(name, data)
	if err != nil {
		return nil, err
	}

	return &RenderArtifact{
		HTML: html,
		Text: text,
	}, nil
}
