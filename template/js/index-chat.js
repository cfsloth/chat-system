$(document).ready(function(){
  var preloadbg = document.createElement("img");
  preloadbg.src = "https://s3-us-west-2.amazonaws.com/s.cdpn.io/245657/timeline1.png";

	$("#search").focus(function(){
		if($(this).val() == "Search contacts..."){
			$(this).val("");
		}
	});
	$("#search").focusout(function(){
		if($(this).val() == ""){
			$(this).val("Search contacts...");

		}
	});

	$("#sendmessage input").focus(function(){
		if($(this).val() == "Send message..."){
			$(this).val("");
		}
	});
	$("#sendmessage input").focusout(function(){
		if($(this).val() == ""){
			$(this).val("Send message...");

		}
	});
/*	$(".friend").each(function(){
		$(this).click(function(){
			var childOffset = $(this).offset();
			var parentOffset = $(this).parent().parent().offset();
			var childTop = childOffset.top - parentOffset.top;
			var clone = $(this).find('img').eq(0).clone();
			var top = childTop+12+"px";

			$(clone).css({'top': top}).addClass("floatingImg").appendTo("#chatbox");

			setTimeout(function(){$("#profile p").addClass("animate");$("#profile").addClass("animate");}, 100);
			setTimeout(function(){
				$("#chat-messages").addClass("animate");
				$('.cx, .cy').addClass('s1');
				setTimeout(function(){$('.cx, .cy').addClass('s2');}, 100);
				setTimeout(function(){$('.cx, .cy').addClass('s3');}, 200);
			}, 150);

			$('.floatingImg').animate({
				'width': "68px",
				'left':'108px',
				'top':'20px'
			}, 200);

			var name = $(this).find("p strong").html();
			var email = $(this).find("p span").html();
			$("#profile p").html(name);
			$("#profile span").html(email);

			$(".message").not(".right").find("img").attr("src", $(clone).attr("src"));
			$('#friendslist').fadeOut();
			$('#chatview').fadeIn();


			$('#close').unbind("click").click(function(){
				$("#chat-messages, #profile, #profile p").removeClass("animate");
				$('.cx, .cy').removeClass("s1 s2 s3");
				$('.floatingImg').animate({
					'width': "40px",
					'top':top,
					'left': '12px'
				}, 200, function(){$('.floatingImg').remove()});

				setTimeout(function(){
					$('#chatview').fadeOut();
					$('#friendslist').fadeIn();
				}, 50);
			});

		});
	});	*/
});

function seeLastMessages(){
  document.getElementById("friends").innerHTML = '<div id="search"><input type="text" id="searchfield" value="Search contacts..." onkeypress="searchForPersonMessage(event)"/></div> '
}

//Ajax call to the search the database for messages of a person. Move this to other directory
function searchForPersonMessage(e){
  //Enter key
  if(e.keyCode == 13){
    alert('Go to database and update')
  }
}

function testAlert(e){
  if (e.keyCode == 13) {
    alert("it works")
  }
}

function loadFriendRequest(){
  //Ajax to load friend requests
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function () {
    if(this.readyState == 4 && this.status == 200){
      var response = xhttp.responseText;
      console.log(response);
    }
  };
  xhttp.open("GET","/findFriendsRequests",true);
  xhttp.send();

}

function contactSearch(){
  document.getElementById("friends").innerHTML = '<div id="search"><input type="text" id="searchfield" value="Search for new friends..." onkeypress="searchForNewFriends(event)"/></div> '
}

function searchForNewFriends(event){
  if(e.keyCode == 13){
    alert("Searching for new friends")
  }
}
