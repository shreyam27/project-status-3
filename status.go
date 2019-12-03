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
    <h3>Final Status Report</h3>
    <span>Decemeber 5</span><br>
</div>
<div>
<h3>Burndown chart</h3>
<img src="https://i.imgur.com/CAwpbfI.png" height="850" width="1000">
<img src="https://i.imgur.com/q0ZVbGV.png" height="450" width="1000">
<h3>Completed Stories across all sprints</h3>
<ul>
<li>View about page - Enhancement</li>
<li>Create event - Enhancement</li>
<li>RSVP related database update - DB Schema Update</li>
<li>RSVP for the event - Enhancement</li>
<li>View event details - Enhancement</li>
<li>Images in events - Bug</li>
<li>Logo in header and footer - Enhancement</li>
<li>Upload app on hosting platform</li>
<li>AB testing (via programming) - Enhancement</li>
<li>Track analytics - Google analytics</li>
<li>Error handling for create page - Enhancement</li>
<li>Error handling for RSVP - Enhancement</li>
<li>Confirmation code for RSVP - Enhancement</li>
<li>Time incorrectly parsed in new events - Bug</li>
<li>Bootstrap CSS inclusion - Enhancement</li>
<li>APIs for the app</li>
<li>Bug fixing</li>
<li>Aesthetics of the app</li>
<li>Prject status report</li>
</ul>
<p>We improved with each spirnt restrospective. In the first Sprint we familarized ourseleves with the starter code and completed simple functionalities. In the second sprint we completed the key stories. In the third sprint we focused on AB testing and additional enhancements  </p>
</div>
<section>
<h3>Traffic analysis - Google analytics</h3>
<p>The percentage of traffic from each of the referres follows: </p>
<img src="https://i.imgur.com/CAwpbfI.png" height="850" width="1000">
</section>
<section>
<h3>AB testing</h3>
<p>"Donate" is more effective than "Support". The percentage of users clicking on "Donate" of the total is 75% of the total clicks on the link whereas 'Support' is 25% of the percentage of the total clicks. We tracked the database counts programmatically. The count of clicks tracked in the database follows:</p>
<img src="https://i.imgur.com/oxbltrV.png" height="275" width="450">
</section>

{{end}}
`

var layoutTemplate = template.Must(template.New("layout").Parse(layoutMarkup))
var statusTemplate = template.Must(template.Must(layoutTemplate.Clone()).Parse(statusMarkup))
