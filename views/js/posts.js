function toggleTab(tabId) {
	if(tabId == 1){
		document.getElementById("friendsTab").classList.add("active");
		document.getElementById("newsFeedTab").classList.remove("active");
		document.getElementById("statusTab").classList.remove("active");

		document.getElementById("friendsForm").style.display = "block";
		document.getElementById("newsFeedForm").style.display = "none";
		document.getElementById("statusForm").style.display = "none";
		
	} else if(tabId == 2){
		document.getElementById("friendsTab").classList.remove("active");
		document.getElementById("newsFeedTab").classList.add("active");
		document.getElementById("statusTab").classList.remove("active");

		document.getElementById("friendsForm").style.display = "none";
		document.getElementById("newsFeedForm").style.display = "block";
		document.getElementById("statusForm").style.display = "none";

	} else if(tabId == 3){
		document.getElementById("friendsTab").classList.remove("active");
		document.getElementById("newsFeedTab").classList.remove("active");
		document.getElementById("statusTab").classList.add("active");

		document.getElementById("friendsForm").style.display = "none";
		document.getElementById("newsFeedForm").style.display = "none";
		document.getElementById("statusForm").style.display = "block";
	}
}