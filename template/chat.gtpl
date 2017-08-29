<!DOCTYPE html>
<html >
<head>
  <meta charset="UTF-8">:
  <title>Chat Messenger</title>
      <link rel="stylesheet" href="css/style-chat.css"> 
</head>

<body>
  <link href='https://fonts.googleapis.com/css?family=Open+Sans:400,600,700' rel='stylesheet' type='text/css'>

<div id="chatbox">
	<div id="friendslist">
    	<div id="topmenu">
            <span class="history" onclick='loadFriendRequests()' ></span>
            <span class="chats" onclick='seeLastMessages()' ></span>
            <span class="friends" onclick='contactSearch()' ></span>
        </div>  
        <div id="friends">
		<div id ="contacts">
			<div class="friend" onclick='' >
				<img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/245657/1_copy.jpg" />
				<p>
				<strong>miro badev</strong>
				<span>mirobadev@gmail.com</span>
				</p>
				<div class="status available"></div>
			</div>
			<div class="friend">
				<img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/245657/2_copy.jpg" />
				<p>
				<strong>Martin Joseph</strong>
				<span>marjoseph@gmail.com</span>
				</p>
				<div class="status away"></div>
			</div>
			<div class="friend">
				<img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/245657/3_copy.jpg" />
				<p>
				<strong>Tomas Kennedy</strong>
				<span>tomaskennedy@gmail.com</span>
				</p>
				<div class="status inactive"></div>
			</div>
		</div> 
		<!--<div id="search">
			<input type="text" id="searchfield" value="Search contacts..." onkeypress="includeHtml(event)"/>
		</div>--> 
	</div>		
	</div>	
  </div>        
</div>

<div id="chat-box">

</div>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
  
    <script src="js/index-chat.js"></script>
    <script>
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
			
	}

	function contactSearch(){
		document.getElementById("friends").innerHTML = '<div id="search"><input type="text" id="searchfield" value="Search for new friends..." onkeypress="searchForNewFriends(event)"/></div> '

	}

	function searchForNewFriends(event){
		if(e.keyCode == 13){
			alert("Searching for new friends")
		}
	}
    </script>
</body>
</html>
