//main.gtpl
{{define "main"}}
{{template "header"}}
<div id="content">	
	<h1>Hello !</h1>

	<div id="post_container">
	<div id="post_buttons">
	
	<nav id=home_buttons role="menu" >
    	<a href="/view/day" class="menu_Element" role="menuitem" > 
			<i class="icon_menu_Today"></i>
			<span class="menu_Element_text">Today</span>
    	</a>
    	<a href="/view/week" class="menu_Element" role="menuitem" >
    		<i class="icon_menu_Week"></i>
			<span class="menu_Element_text">Week</span>
    	</a>
    	<a href="/view/insertPage" class="menu_Element" role="menuitem" >
    		<i class="icon_menu_insertPage"></i>
			<span class="menu_Element_text">InsertPage</span>
    	</a>
    	<a href="/view/stats" class="menu_Element" role="menuitem" >
    		<i class="icon_menu_More"></i>
			<span class="menu_Element_text">More</span>
    	</a>
  	</nav>
 	</div> 
 	</div> 
</div>
{{template "footer"}}
{{end}}