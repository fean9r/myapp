//Error.gtpl
{{define "Error"}}
{{template "header"}}
<div id="content">
  	<h2>General Error </h2>
  	<div id="post_container">
    	<p> An error {{.}} occurred!</p>
    </div>
</div>
{{template "footer"}}
{{end}}