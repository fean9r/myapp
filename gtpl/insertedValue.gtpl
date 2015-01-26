//insertedValue.gtpl
{{define "insertedValue"}}
{{template "header"}}
<div id="content">
  	<h2>You inserted </h2>
  	<div id="post_container">
    	<pre>{{.Date}}</pre>
    </div>
</div>
{{template "footer"}}
{{end}}