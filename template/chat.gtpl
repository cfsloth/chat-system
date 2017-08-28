<!DOCTYPE html>
<html >
<head>
  <meta charset="UTF-8">:
  <title>Chat Messenger</title>
      <link rel="stylesheet" href="css/style-chat.css"> 
</head>

<body>
  <link href='https://fonts.googleapis.com/css?family=Open+Sans:400,600,700' rel='stylesheet' type='text/css'>

  <!--<a id="view-code" href="https://codepen.io/virgilpana/pen/ZYZXgP" target="_blank">VIEW CODE</a> -->

<div id="chatbox">
	<div id="friendslist">
    	<div id="topmenu">
            <span class="friends"></span>
            <span class="chats" onclick='seeLastMessages()' ></span>
            <span class="history"></span>
        </div>      
        <div id="friends">
		<div id ="friends1">
			<div class="friend">
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
		<div id="search">
			<input type="text" id="searchfield" value="Search contacts..." />
		</div>
	</div>		
	</div>	
    
  <!--  <div id="chatview" class="p1">    	
  -->
  </div>        
</div>	
	

	
		

<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
  
    <script src="js/index-chat.js"></script>
    <script>
	function seeLastMessages(){
		document.getElementById("friends1").innerHTML = "";
	}
    </script>
</body>
</html>
