//insertPage.gtpl
{{define "insertPage"}}
{{template "header"}}
<div id="content">	
    <h2>Insert Page</h2>
  	<div id="post_container">
    	<p>Tell me the last 1 day</p>
      <form action="/insert/insertedValue" method="post" > 
        <input type="date" name="date"> 
        <input type="hidden" name="token" value="{{.Token}}">
        <input type="submit" value="Ok">
      </form> 
    </div>
</div>
{{template "footer"}}
{{end}}