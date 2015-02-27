//day.gtpl
{{define "day"}}
{{template "header"}}
<div id="content">	
	<h2>Day</h2>
  	<div id="post_container">
    	<p>Today is:</p>
    	<pre>{{.Day }} {{.Month }} {{.Year }} </pre>
    	<pre>{{.Dates }} </pre>
    	<p>Advises are:</p>
    </div>
</div> 
{{template "footer"}}
{{end}}