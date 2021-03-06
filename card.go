package main

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

// Card 表示一个卡片的实体类。
type Card struct {
	ID         string   `yaml:"id"`
	Title      string   `yaml:"title"`
	Tags       []string `yaml:"tags"`
	Draft      bool     `yaml:"draft"`
	Question   string
	Answer     string
	ReviewTime time.Time
	Hash       string
	Level      int
}

// GenerateAnEmptyCard 生成一张空的卡片的内容
func GenerateAnEmptyCard(id string) []byte {
	tmpl, _ := template.New("card").Parse(
		`---
id: {{.id}}
title: null
draft: true
tags:
    - null
---

<!--front-->
todo

<!--back-->
todo
`)
	data := map[string]string{"id": id}
	buf := bytes.NewBuffer(nil)
	tmpl.Execute(buf, data)
	return buf.Bytes()
}

func (card *Card) String() string {
	s := ""
	s = s + "tags: " + "[" + strings.Join(card.Tags, ",") + "]\n"
	s = s + "title: " + card.Title + "\n"
	return s
}
