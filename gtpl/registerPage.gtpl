//registerPage.gtpl
{{define "registerPage"}}
{{template "header"}}
<div id="content">
  	<h2>Register yourself to access MyApp!</h2>
  	<div id="post_container">
    	<form action="/insert/register" method="post">
    		Email:<input type="text" name="username">
    		Username:<input type="text" name="username">
    		Password:<input type="password" name="password">
    		<input type="hidden" name="token" value="{{.Token}}">
    		<input type="submit" value="Register">
		</form>
    </div>
</div>
{{template "footer"}}
{{end}}