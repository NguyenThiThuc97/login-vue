<template>
<div class="hold-transition sidebar-mini">
	<!-- Site wrapper -->
	<div class="wrapper">
	<!-- Navbar -->
	<nav class="main-header navbar navbar-expand navbar-white navbar-light">
		<!-- Left navbar links -->
		<ul class="navbar-nav">
		<li class="nav-item">
			<a class="nav-link" data-widget="pushmenu" href="#" role="button"><font-awesome-icon icon="bars" /></a>
		</li>
		<li class="nav-item d-none d-sm-inline-block">
			<a href="../../index3.html" class="nav-link">Home</a>
		</li>
		<li class="nav-item d-none d-sm-inline-block">
			<a href="#" class="nav-link">Contact</a>
		</li>
		</ul>

		<!-- Right navbar links -->
		<ul class="navbar-nav ml-auto">
		<!-- Navbar Search -->
		<li class="nav-item">
			<a class="nav-link" data-widget="navbar-search" href="#" role="button">
			<font-awesome-icon icon="search"/>
			</a>
			<div class="navbar-search-block">
			<form class="form-inline">
				<div class="input-group input-group-sm">
				<input class="form-control form-control-navbar" type="search" placeholder="Search" aria-label="Search">
				<div class="input-group-append">
					<button class="btn btn-navbar" type="submit">
					<font-awesome-icon icon="search"/>
					</button>
					<button class="btn btn-navbar" type="button" data-widget="navbar-search">
					<font-awesome-icon icon="times"/>
					</button>
				</div>
				</div>
			</form>
			</div>
		</li>

		<!-- Messages Dropdown Menu -->
		<li class="nav-item dropdown">
			<a class="nav-link" data-toggle="dropdown" href="#">
			<font-awesome-icon icon="comments" />
			<span class="badge badge-danger navbar-badge">3</span>
			</a>
			<div class="dropdown-menu dropdown-menu-lg dropdown-menu-right">
			<a href="#" class="dropdown-item">
				<!-- Message Start -->
				<div class="media">
				<img :src="require('@/assets/dist/img/user1-128x128.jpg')" alt="User Avatar" class="img-size-50 mr-3 img-circle">
				<div class="media-body">
					<h3 class="dropdown-item-title">
					Brad Diesel
					<span class="float-right text-sm text-danger"><font-awesome-icon icon="star" /></span>
					</h3>
					<p class="text-sm">Call me whenever you can...</p>
					<p class="text-sm text-muted"><font-awesome-icon icon="clock" class="mr-1" /> 4 Hours Ago</p>
				</div>
				</div>
				<!-- Message End -->
			</a>
			<div class="dropdown-divider"></div>
			<a href="#" class="dropdown-item">
				<!-- Message Start -->
				<div class="media">
				<img :src="require('@/assets/dist/img/user8-128x128.jpg')" alt="User Avatar" class="img-size-50 img-circle mr-3">
				<div class="media-body">
					<h3 class="dropdown-item-title">
					John Pierce
					<span class="float-right text-sm text-muted"><font-awesome-icon icon="star" /></span>
					</h3>
					<p class="text-sm">I got your message bro</p>
					<p class="text-sm text-muted"><font-awesome-icon icon="clock" class="mr-1" /> 4 Hours Ago</p>
				</div>
				</div>
				<!-- Message End -->
			</a>
			<div class="dropdown-divider"></div>
			<a href="#" class="dropdown-item">
				<!-- Message Start -->
				<div class="media">
				<img :src="require('@/assets/dist/img/user3-128x128.jpg')" alt="User Avatar" class="img-size-50 img-circle mr-3">
				<div class="media-body">
					<h3 class="dropdown-item-title">
					Nora Silvester
					<span class="float-right text-sm text-warning"><font-awesome-icon icon="star" /></span>
					</h3>
					<p class="text-sm">The subject goes here</p>
					<p class="text-sm text-muted"><font-awesome-icon icon="clock" class="mr-1" /> 4 Hours Ago</p>
				</div>
				</div>
				<!-- Message End -->
			</a>
			<div class="dropdown-divider"></div>
			<a href="#" class="dropdown-item dropdown-footer">See All Messages</a>
			</div>
		</li>
		<!-- Notifications Dropdown Menu -->
		<li class="nav-item dropdown">
			<a class="nav-link" data-toggle="dropdown" href="#">
			<font-awesome-icon icon="bell" />
			<span class="badge badge-warning navbar-badge">15</span>
			</a>
			<div class="dropdown-menu dropdown-menu-lg dropdown-menu-right">
			<span class="dropdown-item dropdown-header">15 Notifications</span>
			<div class="dropdown-divider"></div>
			<a href="#" class="dropdown-item">
				<font-awesome-icon icon="envelope" class="mr-2" /> 4 new messages
				<span class="float-right text-muted text-sm">3 mins</span>
			</a>
			<div class="dropdown-divider"></div>
			<a href="#" class="dropdown-item">
				<font-awesome-icon icon="users" class="mr-2" /> 8 friend requests
				<span class="float-right text-muted text-sm">12 hours</span>
			</a>
			<div class="dropdown-divider"></div>
			<a href="#" class="dropdown-item">
				<font-awesome-icon icon="file" class="mr-2" /> 3 new reports
				<span class="float-right text-muted text-sm">2 days</span>
			</a>
			<div class="dropdown-divider"></div>
			<a href="#" class="dropdown-item dropdown-footer">See All Notifications</a>
			</div>
		</li>
		<li class="nav-item">
			<a class="nav-link" data-widget="fullscreen" href="#" role="button">
			<font-awesome-icon icon="expand-arrows-alt" />
			</a>
		</li>
		<li class="nav-item">
			<a class="nav-link" data-widget="control-sidebar" data-slide="true" href="#" role="button">
			<font-awesome-icon icon="th-large" />
			</a>
		</li>
		</ul>
	</nav>
	<!-- /.navbar -->

	<!-- Main Sidebar Container -->
	<aside class="main-sidebar sidebar-dark-primary elevation-4">
		<!-- Brand Logo -->
		<a href="../../index3.html" class="brand-link">
		<img :src="require('@/assets/dist/img/AdminLTELogo.png')" alt="AdminLTE Logo" class="brand-image img-circle elevation-3" style="opacity: .8">
		<span class="brand-text font-weight-light">AdminLTE 3</span>
		</a>

		<!-- Sidebar -->
		<div class="sidebar">
		<!-- Sidebar user (optional) -->
		<div class="user-panel mt-3 pb-3 mb-3 d-flex">
			<div class="image">
			<img :src="require('@/assets/dist/img/user2-160x160.jpg')" class="img-circle elevation-2" alt="User Image">
			</div>
			<div class="info">
			<a href="#" class="d-block">Alexander Pierce</a>
			</div>
		</div>

		<!-- SidebarSearch Form -->
		<div class="form-inline">
			<div class="input-group" data-widget="sidebar-search">
			<input class="form-control form-control-sidebar" type="search" placeholder="Search" aria-label="Search">
			<div class="input-group-append">
				<button class="btn btn-sidebar">
				<font-awesome-icon icon="search" />
				</button>
			</div>
			</div>
		</div>

		<!-- Sidebar Menu -->
		<nav class="mt-2">
			<ul class="nav nav-pills nav-sidebar flex-column" data-widget="treeview" role="menu" data-accordion="false">
			<!-- Add icons to the links using the .nav-icon class
				with font-awesome or any other icon font library -->
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="tachometer-alt" />
				<p>
					Dashboard
					<font-awesome-icon class="right" icon="angle-left" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../../index.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Dashboard v1</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../../index2.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Dashboard v2</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../../index3.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Dashboard v3</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="../widgets.html" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="th" />
				<p>
					Widgets
					<span class="right badge badge-danger">New</span>
				</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="copy" />
				<p>
					Layout Options
					<font-awesome-icon icon="angle-left" class="right" />
					<span class="badge badge-info right">6</span>
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../layout/top-nav.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Top Navigation</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../layout/top-nav-sidebar.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Top Navigation + Sidebar</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../layout/boxed.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Boxed</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../layout/fixed-sidebar.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Fixed Sidebar</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../layout/fixed-sidebar-custom.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Fixed Sidebar <small>+ Custom Area</small></p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../layout/fixed-topnav.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Fixed Navbar</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../layout/fixed-footer.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Fixed Footer</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../layout/collapsed-sidebar.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Collapsed Sidebar</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="chart-pie" />
				<p>
					Charts
					<font-awesome-icon class="right" icon="angle-left" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../charts/chartjs.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>ChartJS</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../charts/flot.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Flot</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../charts/inline.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Inline</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../charts/uplot.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>uPlot</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="tree" />
				<p>
					UI Elements
					<font-awesome-icon icon="angle-left" class="right" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../UI/general.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>General</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../UI/icons.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Icons</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../UI/buttons.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Buttons</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../UI/sliders.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Sliders</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../UI/modals.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Modals & Alerts</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../UI/navbar.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Navbar & Tabs</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../UI/timeline.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Timeline</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../UI/ribbons.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Ribbons</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="edit" />
				<p>
					Forms
					<font-awesome-icon icon="angle-left" class="right" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../forms/general.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>General Elements</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../forms/advanced.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Advanced Elements</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../forms/editors.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Editors</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../forms/validation.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Validation</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="table" />
				<p>
					Tables
					<font-awesome-icon icon="angle-left" class="right" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../tables/simple.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Simple Tables</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../tables/data.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>DataTables</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../tables/jsgrid.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>jsGrid</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-header">EXAMPLES</li>
			<li class="nav-item">
				<a href="../calendar.html" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="calendar-alt" />
				<p>
					Calendar
					<span class="badge badge-info right">2</span>
				</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="../gallery.html" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="image" />
				<p>
					Gallery
				</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="../kanban.html" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="columns" />
				<p>
					Kanban Board
				</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="envelope" />
				<p>
					Mailbox
					<font-awesome-icon icon="angle-left" class="right" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../mailbox/mailbox.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Inbox</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../mailbox/compose.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Compose</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../mailbox/read-mail.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Read</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item menu-open">
				<a href="#" class="nav-link active">
				<font-awesome-icon class="nav-icon" icon="book" />
				<p>
					Pages
					<font-awesome-icon icon="angle-left" class="right" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../examples/invoice.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Invoice</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/profile.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Profile</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/e-commerce.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>E-commerce</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/projects.html" class="nav-link active">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Projects</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/project-add.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Project Add</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/project-edit.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Project Edit</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/project-detail.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Project Detail</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/contacts.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Contacts</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/faq.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>FAQ</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/contact-us.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Contact us</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="plus-square" />
				<p>
					Extras
					<font-awesome-icon icon="angle-left" class="right" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="#" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>
						Login & Register v1
						<font-awesome-icon icon="angle-left" class="right" />
					</p>
					</a>
					<ul class="nav nav-treeview">
					<li class="nav-item">
						<a href="../examples/login.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Login v1</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="../examples/register.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Register v1</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="../examples/forgot-password.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Forgot Password v1</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="../examples/recover-password.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Recover Password v1</p>
						</a>
					</li>
					</ul>
				</li>
				<li class="nav-item">
					<a href="#" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>
						Login & Register v2
						<font-awesome-icon icon="angle-left" class="right" />
					</p>
					</a>
					<ul class="nav nav-treeview">
					<li class="nav-item">
						<a href="../examples/login-v2.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Login v2</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="../examples/register-v2.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Register v2</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="../examples/forgot-password-v2.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Forgot Password v2</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="../examples/recover-password-v2.html" class="nav-link">
						<font-awesome-icon icon="circle" class="nav-icon" />
						<p>Recover Password v2</p>
						</a>
					</li>
					</ul>
				</li>
				<li class="nav-item">
					<a href="../examples/lockscreen.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Lockscreen</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/legacy-user-menu.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Legacy User Menu</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/language-menu.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Language Menu</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/404.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Error 404</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/500.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Error 500</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/pace.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Pace</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../examples/blank.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Blank Page</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../../starter.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Starter Page</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="search" />
				<p>
					Search
					<font-awesome-icon icon="angle-left" class="right" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="../search/simple.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Simple Search</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="../search/enhanced.html" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Enhanced</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-header">MISCELLANEOUS</li>
			<li class="nav-item">
				<a href="../../iframe.html" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="ellipsis-h" />
				<p>Tabbed IFrame Plugin</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="https://adminlte.io/docs/3.1/" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="file" />
				<p>Documentation</p>
				</a>
			</li>
			<li class="nav-header">MULTI LEVEL EXAMPLE</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon icon="circle" class="nav-icon" />
				<p>Level 1</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon" icon="circle" />
				<p>
					Level 1
					<font-awesome-icon class="right" icon="angle-left" />
				</p>
				</a>
				<ul class="nav nav-treeview">
				<li class="nav-item">
					<a href="#" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Level 2</p>
					</a>
				</li>
				<li class="nav-item">
					<a href="#" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>
						Level 2
						<font-awesome-icon class="right" icon="angle-left" />
					</p>
					</a>
					<ul class="nav nav-treeview">
					<li class="nav-item">
						<a href="#" class="nav-link">
						<font-awesome-icon icon="dot-circle" class="nav-icon" />
						<p>Level 3</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="#" class="nav-link">
						<font-awesome-icon icon="dot-circle" class="nav-icon" />
						<p>Level 3</p>
						</a>
					</li>
					<li class="nav-item">
						<a href="#" class="nav-link">
						<font-awesome-icon icon="dot-circle" class="nav-icon" />
						<p>Level 3</p>
						</a>
					</li>
					</ul>
				</li>
				<li class="nav-item">
					<a href="#" class="nav-link">
					<font-awesome-icon icon="circle" class="nav-icon" />
					<p>Level 2</p>
					</a>
				</li>
				</ul>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon icon="circle" class="nav-icon" />
				<p>Level 1</p>
				</a>
			</li>
			<li class="nav-header">LABELS</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon text-danger" icon="circle" />
				<p class="text">Important</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon text-warning" icon="circle" />
				<p>Warning</p>
				</a>
			</li>
			<li class="nav-item">
				<a href="#" class="nav-link">
				<font-awesome-icon class="nav-icon text-info" icon="circle" />
				<p>Informational</p>
				</a>
			</li>
			</ul>
		</nav>
		<!-- /.sidebar-menu -->
		</div>
		<!-- /.sidebar -->
	</aside>

	<!-- Content Wrapper. Contains page content -->
	<div class="content-wrapper">
		<!-- Content Header (Page header) -->
		<section class="content-header">
		<div class="container-fluid">
			<div class="row mb-2">
			<div class="col-sm-6">
				<h1>Projects</h1>
			</div>
			<div class="col-sm-6">
				<ol class="breadcrumb float-sm-right">
				<li class="breadcrumb-item"><a href="#">Home</a></li>
				<li class="breadcrumb-item active">Projects</li>
				</ol>
			</div>
			</div>
		</div><!-- /.container-fluid -->
		</section>

		<!-- Main content -->
		<section class="content">

		<!-- Default box -->
		<div class="card">
			<div class="card-header">
			<h3 class="card-title">Projects</h3>

			<div class="card-tools">
				<button type="button" class="btn btn-tool" data-card-widget="collapse" title="Collapse">
				<font-awesome-icon icon="minus" />
				</button>
				<button type="button" class="btn btn-tool" data-card-widget="remove" title="Remove">
				<font-awesome-icon icon="times" />
				</button>
			</div>
			</div>
			<div class="card-body p-0">
			<table class="table table-striped projects">
				<thead>
					<tr>
						<th style="width: 1%">
							#
						</th>
						<th style="width: 20%">
							Project Name
						</th>
						<th style="width: 30%">
							Team Members
						</th>
						<th>
							Project Progress
						</th>
						<th style="width: 8%" class="text-center">
							Status
						</th>
						<th style="width: 20%">
						</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar2.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar3.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar4.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="57" aria-valuemin="0" aria-valuemax="100" style="width: 57%">
								</div>
							</div>
							<small>
								57% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar2.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="47" aria-valuemin="0" aria-valuemax="100" style="width: 47%">
								</div>
							</div>
							<small>
								47% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar2.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar3.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="77" aria-valuemin="0" aria-valuemax="100" style="width: 77%">
								</div>
							</div>
							<small>
								77% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar2.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar3.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar4.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%">
								</div>
							</div>
							<small>
								60% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar4.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar5.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="12" aria-valuemin="0" aria-valuemax="100" style="width: 12%">
								</div>
							</div>
							<small>
								12% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar2.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar3.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar4.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="35" aria-valuemin="0" aria-valuemax="100" style="width: 35%">
								</div>
							</div>
							<small>
								35% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar4.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar5.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="87" aria-valuemin="0" aria-valuemax="100" style="width: 87%">
								</div>
							</div>
							<small>
								87% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar3.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar4.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="77" aria-valuemin="0" aria-valuemax="100" style="width: 77%">
								</div>
							</div>
							<small>
								77% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
					<tr>
						<td>
							#
						</td>
						<td>
							<a>
								AdminLTE v3
							</a>
							<br/>
							<small>
								Created 01.01.2019
							</small>
						</td>
						<td>
							<ul class="list-inline">
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar3.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar4.png')">
								</li>
								<li class="list-inline-item">
									<img alt="Avatar" class="table-avatar" :src="require('@/assets/dist/img/avatar5.png')">
								</li>
							</ul>
						</td>
						<td class="project_progress">
							<div class="progress progress-sm">
								<div class="progress-bar bg-green" role="progressbar" aria-valuenow="77" aria-valuemin="0" aria-valuemax="100" style="width: 77%">
								</div>
							</div>
							<small>
								77% Complete
							</small>
						</td>
						<td class="project-state">
							<span class="badge badge-success">Success</span>
						</td>
						<td class="project-actions text-right">
							<a class="btn btn-primary btn-sm" href="#">
								<font-awesome-icon icon="folder" />
								View
							</a>
							<a class="btn btn-info btn-sm" href="#">
								<font-awesome-icon icon="pencil-alt" />
								Edit
							</a>
							<a class="btn btn-danger btn-sm" href="#">
								<font-awesome-icon icon="trash" />
								Delete
							</a>
						</td>
					</tr>
				</tbody>
			</table>
			</div>
			<!-- /.card-body -->
		</div>
		<!-- /.card -->

		</section>
		<!-- /.content -->
	</div>
	<!-- /.content-wrapper -->

	<footer class="main-footer">
		<div class="float-right d-none d-sm-block">
		<b>Version</b> 3.2.0-rc
		</div>
		<strong>Copyright &copy; 2014-2021 <a href="https://adminlte.io">AdminLTE.io</a>.</strong> All rights reserved.
	</footer>

	<!-- Control Sidebar -->
	<aside class="control-sidebar control-sidebar-dark">
		<!-- Control sidebar content goes here -->
	</aside>
	<!-- /.control-sidebar -->
	</div>
	<!-- ./wrapper -->
</div>
</template>
<script>
export default {
	computed: {
		currentUser() {
			return this.$store.state.auth.user;
		},
	},
	methods: {
		logOut() {
		this.$store.dispatch('auth/logout');
		this.$router.push('/login');
		}
	}
}
</script>