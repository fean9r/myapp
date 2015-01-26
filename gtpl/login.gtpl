//login.gtpl
{{define "login"}}
{{template "header"}}
<div id="content">
  	<h2>Login </h2>
  	<div id="post_container">
    	<form action="/insert/login" method="post">
    		Username:<input type="text" name="username">
    		Password:<input type="password" name="password">
    		<input type="hidden" name="token" value="{{.Token}}">
    		<input type="submit" value="Login">
		</form>
    </div>
</div>
{{template "footer"}}
{{end}}