function seeLastMessages(){
      document.getElementById("friends").innerHTML = '<div id="search"><input type="text" id="searchfield" value="Search contacts..." onkeypress="searchForPersonMessage(event)"/></div> '
    }

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
  var valueOfemail = "";
  var http = new XMLHttpRequest();
  var url = "findFriendsRequests";
  var cookies = document.cookie.split(";");
  for(i=0;i<cookies.length;i++){
    var nameOfCookie = cookies[i].split("=");
    if(nameOfCookie[0] === " email"){
      valueOfemail = nameOfCookie[1];
      break;
    }
  }
  var params = "email=" + valueOfemail;
  http.open("POST", url, true);
  //Send the proper header information along with the request
  http.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  http.onreadystatechange = function() {
    if(http.readyState == 4 && http.status == 200) {
      var jsonReceived = JSON.parse(http.responseText);
      var aux;
      var html = "";
      for(i=0;i<jsonReceived.length;i++){
        aux = '<div class="friend">'
        + '<img src="images/user.jpg" />'
        + '<p>'
        + '<strong>'+ jsonReceived[i]["FROMNAME"] +'</strong>'
        + '<br>'
        + '<span>'+ jsonReceived[i]["CELLPHONE"] +'</span>'
        //+ '</p>'
        + '<div class="status away" style="position:absolute;"></div>'
        + '<button type="button" style="height:15px;width:15px;background-color:green;border-radius:0.6em 0.6em;" placeholder="Accept"></button>'
        + '<button type="button" style="height:15px;width:15px;background-color:red;border-radius:0.6em 0.6em;" placeholder="Reject"></button>'
        + '</div>';
        html = html + aux;
      }
      document.getElementById("friends").innerHTML = html
    }
  }
  http.send(params);
}

function contactSearch(){
  html = document.getElementById("friends").innerHTML;
  document.getElementById("friends").innerHTML = '<div id="search"><input type="text" id="searchfield" value="Search for new friends..." onkeypress="searchForNewFriends(event)"/></div> '
}

function searchForNewFriends(event){
  if(e.keyCode == 13){
    alert("Searching for new friends")
  }
}
