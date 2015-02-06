//header.gtpl
{{define "header"}}
<html ng-app="schedulerApp">
<head >
	<meta charset="utf-8">
	<link rel="stylesheet" href="/stylesheets/myapp.css">
	<link rel="stylesheet" href="/lib/scheduler/dhtmlxscheduler.css">
 
	<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
 
	<script src="/lib/angular/angular.min.js"></script>
	<script src="/lib/scheduler/dhtmlxscheduler.js"></script>
	<script src="/scripts/app.js"></script>
	<script src="/scripts/app.scheduler.js"></script>
	<script src="/scripts/myscripts.js"></script>
	<script type="text/javascript">
		// some javascript script 
	</script>

</head>
  <body class="app" ng-controller="MainSchedulerCtrl">
  	<div class="header">
  		<h1 class="logo " id="logo">
  			<a class="logo_anchor" href="/view/main">MyApp</a>
  		</h1>
  		<a class="login" href="/view/loginPage">Login</a>
  	</div> <!-- End header-->
  	<div class="container"> <!-- Start Container-->
{{end}}