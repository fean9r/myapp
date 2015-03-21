//login.gtpl
{{define "login"}}
{{template "header"}}
<div id="content">
  	<h2>Login Success </h2>
  	<div id="post_container">
    	<p> Ciao</p> <pre>{{.Username}}</pre>  
    </div>
</div>
{{template "footer"}}
{{end}}