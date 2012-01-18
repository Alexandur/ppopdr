var PAGES = {
		
	defaultSearchValue: "Zoeken..",
	searchPage: 0,
	
	/**
	 * This method binds listeners to buttons etc.
	 * Is ran when application is loaded
	 */
	init: function()
	{
		//Homepage searchbox actions
		$('#searchButton').click(function(){
			PAGES.search($('#searchBox').attr('value'));
		});
		$('#searchBox').focus(function(){
			if($(this).attr('value') == PAGES.defaultSearchValue)
			{
				$(this).attr('value', "");
			}
		});
		$('#searchBox').blur(function(){
			if($(this).attr('value') == "")
			{
				$(this).attr('value', PAGES.defaultSearchValue);
			}
		});
		
		//bind all buttons with classes 'button' and 'previous'
		PAGES.rebindButtons();
	},
	
	/**
	 * Unbinds all buttons with classes 'button' and 'previous', then rebinds all buttons including new ones.
	 * To avoid double actions.
	 */
	rebindButtons: function()
	{
		$('.previous').unbind('click');
		$('.previous').click(PAGES.goBack);
		
		$('.button').unbind('click');
		$('.button').click(function()
		{
			PAGES.getNewPage($(this));
		});
	},

	/**
	 * This method makes the app go one page back in history
	 */
	goBack: function()
	{
		history.back();
		return false;
	},
	
	/**
	 * This method searches all articles
	 */
	search: function(search)
	{
		$.ajax({
			url: "http://aglansbeek.appspot.com/search/" + search,
			success: function(data){
				PAGES.createNewSearchPage(data, search);
			},
			error: function(){
				PAGES.failed();
			}
		});
	},
	
	createNewSearchPage: function(jsonData, search)
	{
		var newPage = '<div data-role="page" id="search' + PAGES.searchPage + '"><div data-role="header"><a class="previous" data-icon="arrow-l">Vorige</a><h1>' + search + '</h1><a href="#home" data-icon="home">Home</a></div><div id="content" data-role="content"><ul data-role="listview">';
		
		for(var i in jsonData.D)
		{
			newPage += '<li><a class="button" title="' + jsonData.D[i].Id + '"><h3>' + jsonData.D[i].Title + '</h3><p>in de categorie ' + jsonData.D[i].Category + '</p></a></li>';
		}
		
		newPage += '</ul></div></div>';
		
		$('body').append(newPage);
		
		PAGES.rebindButtons();
		$.mobile.changePage("#search" + PAGES.searchPage);
		
		PAGES.searchPage++;
	},
	
	/**
	 * Gets a new page. Connect with webservice for data.
	 * Via an ajax request json data is loaded and depending on the new page
	 *  - insertNewCategoryPage
	 *  - insertNewAricleListPage
	 *  - insertNewArticlePage
	 *  is called.
	 */
	getNewPage: function(o)
	{
		var what = o.attr('title')

		if(what == 'hcats')
		{
			$.ajax({
				url: "http://aglansbeek.appspot.com/category",
				success: function(data){
					PAGES.insertNewHCategoryPage(data, what);
				},
				error: function(){
					PAGES.failed();
				}
			});
		}
		else if(what == 'categorie')
		{
			$.ajax({
				url: "http://aglansbeek.appspot.com/article/category/" + o.attr('rel'),
				success: function(data){
					PAGES.insertNewCategoryPage(data, o.attr('rel'));
				},
				error: function(){
					PAGES.failed();
				}
			});
		}
		else if(parseInt(what))
		{
			$.ajax({
				url: "http://aglansbeek.appspot.com/article/" + what,
				success: function(data){
					PAGES.insertNewArticlePage(data, what);
				},
				error: function(){
					PAGES.failed();
				}
			});
		}
		else
		{
			$.ajax({
				url: "http://aglansbeek.appspot.com/article/category/" + what,
				success: function(data){
					PAGES.insertNewArticleListPage(data, what);
				},
				error: function(){
					PAGES.failed();
				}
			});
		}
	},
	
	insertNewHCategoryPage: function(jsonData, id)
	{
		var newPage = '<div data-role="page" id="' + id + '"><div data-role="header"><a class="previous" data-icon="arrow-l">Vorige</a><h1>Categorien</h1><a href="#home" data-icon="home">Home</a></div><div id="content" data-role="content">';
		
		for(var i in jsonData.D)
		{
			newPage += '<a class="button" title="categorie" rel="' + jsonData.D[i].Id + '" data-role="button">' + jsonData.D[i].Title + '</a>';
		}
		
		newPage += '</div></div>';
		
		$('body').append(newPage);
		$.mobile.changePage('#' + id);
		
		//rebind buttons
		PAGES.init();
	},
	
	/**
	 * Insert new category page. Lists all categories as buttons in the app gotten from the webservice
	 */
	insertNewCategoryPage: function(jsonData, id)
	{
		var newPage = '<div data-role="page" id="c' + id + '"><div data-role="header"><a class="previous" data-icon="arrow-l">Vorige</a><h1>Categorien</h1><a href="#home" data-icon="home">Home</a></div><div id="content" data-role="content">';
		
		for(var name in jsonData.D)
		{
			newPage += '<div class="collapsible" data-icon="plus" data-role="collapsible" data-theme="b" data-content-theme="d"><h3>' + name + '</h3>';
			for(var article in jsonData.D[name])
			{
				newPage += '<a class="button" title="' + jsonData.D[name][article].Id + '" data-role="button">' + jsonData.D[name][article].Title + '</a>';
			}
			newPage += '</div>'
		}
		
		newPage += '</div></div>';
		
		$('body').append(newPage);
		$.mobile.changePage('#c' + id);
		
		//rebind buttons
		PAGES.init();
	},
	
	/**
	 * Insert new article page. Displays an article gotten from the webservice.
	 */
	insertNewArticlePage: function(jsonData, id)
	{
		var newPage = '<div data-role="page" id="p' + id + '"><div data-role="header"><a class="previous" data-icon="arrow-l">Vorige</a><h1>' + jsonData.D.Title + '</h1><a href="#home" data-icon="home">Home</a></div><div id="content" data-role="content">';
		newPage += jsonData.D.Content;
		newPage += '<h3>Tips</h3>';
		for(var tip in jsonData.D.Tips)
		{
			newPage += '<h4>' + jsonData.D.Tips[tip].Author + '</h4><p>' + jsonData.D.Tips[tip].Content + '</p>';
		}
		newPage += '</div></div>';
		
		$('body').append(newPage);
		$.mobile.changePage('#p' + id);
		
		//rebind buttons
		PAGES.rebindButtons();
	},
	
	/**
	 * This method is fired when an ajax request fails.
	 */
	failed: function()
	{
		navigator.notification.alert("Het is niet gelukt de gewenste data op te halen.");
	}
}