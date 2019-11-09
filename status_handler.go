package main

import (
	"html/template"
)

var layoutMarkup = `
	{{define "layout"}}
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body {
					max-width: 800px;
					margin: 2rem auto;
				}

				header {
					background-color: #8ac5c3;
				}

				footer {
					background-color: #edece8;
				}
			</style>
			<title>Project status report</title>
		</head>
		<body>
			<header></header>
			{{block "content" .}}

			{{end}}
			<footer>This is the footer</footer>
		</body>
	{{end}}
`

var indexMarkup = `
{{define "content"}}
	<p>
		
    </p>
{{end}}
`
