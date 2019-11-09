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
    <h3>Project Status Report 2</h3>
    <span>November 9</span><br>
</div>
<div>
<h3>Burndown chart</h3>
<img src="/Burndown chart.PNG" height="850" width="1000">
<img src="/Completed stories.PNG" height="450" width="1000">
<h3> Sprint Completed Stories</h3>
<ul>
<li>View about page - Enhancement</li>
<li>Create event - Enhancement</li>
<li>RSVP related database update - DB Schema Update</li>
<li>RSVP for the event - Enhancement</li>
<li>View event details - Enhancement</li>
<li>Images in events - Bug</li>
</ul>
<p>We worked on the functionalities that add most value to the user including creating events, viewing event details, RSVP for the events and about page. We also worked on bug fixing. We faced challenges related to parsing of time on the create events page which we will fix in the next sprint. In this sprint, we worked on separate branches and merged the pull requests. This enabled the devlopers to work on the code simultaneously.  </p>
</div>
<div>
<h3>Product Backlog</h3>
<p>While we have not determined the final contents of the sprint 3, we will focus on analytics tracking. We will also spruce up the app with error handling, improving look and feel and fixing bugs.
<li>Tracking analytics</li>
<li>Error handling</li>
<li>Aesthetics of the app</li>
</div>
{{end}}
`

var layoutTemplate = template.Must(template.New("layout").Parse(layoutMarkup))
var statusTemplate = template.Must(template.Must(layoutTemplate.Clone()).Parse(statusMarkup))
