$(document).bind("mobileinit", function(){
	$.mobile.defaultPageTransition = 'none';
});


function onLoad()
{
	document.addEventListener("deviceready", onDeviceReady, true);
}

function onDeviceReady()
{
	PAGES.init();
	
	var networkState = navigator.network.connection.type;

	if(networkState == Connection.NONE || networkState == undefined)
	{
		navigator.notification.alert("U heeft geen internetverbinding, maak eerst verbinding om gebruik te maken van deze app.", onDeviceReady);
	}
	if(networkState == Connection.CELL_2G || networkState == Connection.CELL_3G || networkState == Connection.CELL_4G)
	{
		navigator.notification.alert("U heeft een internetverbinding via uw provider, pas op dat u niet over uw eventuele datalimiet gaat.");
	}
}