<!DOCTYPE html>
<html >
<head>
  <meta charset="UTF-8">
  <title>Login</title>

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/meyer-reset/2.0/reset.min.css">

  <link rel='stylesheet prefetch' href='https://fonts.googleapis.com/css?family=Roboto:400,100,300,500,700,900'>
<link rel='stylesheet prefetch' href='https://fonts.googleapis.com/css?family=Montserrat:400,700'>
<link rel='stylesheet prefetch' href='https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css'>

      <link rel="stylesheet" href="css/style.css">


</head>

<body>

<div class="container">
  <div class="info">
	  <h1>Login</h1>
	  <!--<span>Made with <i class="fa fa-heart"></i> by <a href="http://andytran.me">Andy Tran</a></span>-->
  </div>
</div>
<div class="form">
	<div class="thumbnail"><img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/169963/hat.svg"/>
	</div>
  <form method="POST" class="register-form">
    <input type="text" name= "name" placeholder="Name" required />
    <input type="password" name="password" placeholder="Password" required />
    <input type="text" name="email" placeholder="Email Address" required/>
    <input type="text" name="cellPhone" placeholder="Phone Number" required/>
    <button type="submit" >Create</button>
    <p class="message">Already registered? <a href="#">Sign In</a></p>
  </form>
  <form method="POST" class="login-form">
    <input type="text" name="email" placeholder="Username" required />
    <input type="password" name="password" placeholder="Password" required/>
    <button type="submit">Login</button>
    <p class="message">Not registered? <a href="#">Create an account</a></p>
  </form>
</div>
<video id="video" autoplay="autoplay" loop="loop" poster="polina.jpg">
  <source src="http://andytran.me/A%20peaceful%20nature%20timelapse%20video.mp4" type="video/mp4"/>
</video>
  <script src='http://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js'></script>

    <script src="js/index.js"></script>

</body>
</html>
