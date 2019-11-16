package main

import (
	"html/template"
)

var layoutMarkup = `
	{{define "layout"}}
		<!DOCTYPE html>
		<html>
		<head>
			<title>Project status report</title>
		</head>
		<body>
			<header></header>
			{{block "content" .}}

			{{end}}
			<footer></footer>
		</body>
	{{end}}
`

var statusMarkup = `
{{define "content"}}
	<div align="right">
        Jack Ly (jack.ly@yale.edu)<br>
        Shreya Manjunath (shreya.manjunath@yale.edu) <br>
        Nikhita Gupta(nikhita.gupta@yale.edu)<br>
        Lira Lourdes(lourdes.lira@yale.edu)<br>
</div>

<div style="text-align:center">
    <h1>Inner Surf Coding Team</h1>
    <h3>Project Status Report 3</h3>
    <span>November 16</span><br>
</div>
<div>
<h3>Burndown chart</h3>
<img src="https://i.imgur.com/CAwpbfI.png" height="775" width="1000">
<img src="https://i.imgur.com/q0ZVbGV.png" height="450" width="1000">
<h3> Sprint Completed Stories</h3>
<ul>
<li>AB testing (via programming) - Enhancement</li>
<li>Track analytics - Google analytics</li>
<li>Error handling for create page - Enhancement</li>
<li>Error handling for RSVP - Enhancement</li>
<li>Confirmation code for RSVP - Enhancement</li>
<li>Time incorrectly parsed in new events - Bug</li>
<li>Bootstrap CSS inclusion - Enhancement</li>
</ul>
<p>We worked on setting up Google analytics and programtically performed AB testing. Additionally, we worked on error handling for create and RSVP page. We also added the hash based confirmation code for the email after RSVP. We also fixed bugs. We faced challenges related to Google Optimize for AB testing and therefore decided to use a programming approach.</p>
</div>
<div>
<h3>Product Backlog</h3>
<p>We will focus on APIs, bug fixing and aesthetic aspects of the app in the remaining time period. We will check the user preference as part of AB testing via tracking the contents of the Heroku DB where we are storing this data.
<ul>
<li>APIs for the app</li>
<li>Bug fixing</li>
<li>Aesthetics of the app</li>
</ul>
</div>
<div>
<h3>Google Analytics</h3>
<p> Google analytics data on the traffic flow from different groups of users follows: </p>
https://i.imgur.com/zf6eG2k.png
<img src="https://i.imgur.com/zf6eG2k.png" height="400" width="1000">
{{end}}
`

var layoutTemplate = template.Must(template.New("layout").Parse(layoutMarkup))
var statusTemplate = template.Must(template.Must(layoutTemplate.Clone()).Parse(statusMarkup))
