package view

func MainPage() string {
	return `
	<!DOCTYPE html>
<html>
<head>
<title>Welcome !</title>
</head>
<body>
	<form action="/createuser" method="get">
	<input type="submit" value="Create a user">
	</form>

	<form action="/user" method="get">
	<input type="submit" value="Get a list of all the users">
	</form>

	<form action="/removeuser" method="get">
	<input type="submit" value="Remove a user">
	</form>

</body>
</html>
	`
}

//The form for the user creation
func CreateUser() string {
	return `
	<!DOCTYPE html>
<html>
<head>
<title>Create a user</title>
</head>
<body>
	<form action="/user" method="post">

  <label for="username">Username:</label>
  <input type="text" id="username" name="username"><br><br>

  <label for="password">Password:</label>
  <input type="password" id="password" name="password"><br><br>

  <label for="email">Email:</label>
  <input type="text" id="email" name="email"><br><br>

  <input type="submit" value="Submit">
</form>
</body>
</html>
	`
}

//The page after the user creation
func PostUser() string {
	return `
	<!DOCTYPE html>
<html>
<head>
<title>Post user</title>
</head>
<body>
	Your user have been successfully created !
	<form action="/" method="get">
  <input type="submit" value="Back to the main page">
</form>
</body>
</html>
	`
}

//The form
func RemoveUser() string {
	return `
	<!DOCTYPE html>
<html>
<head>
<title>Remove a user</title>
</head>
<body>
	Delete a user: 
	<form action="/deleteuser" method="delete">
  <label for="user">Username:</label>
  <input type="text" id="username" name="username"><br><br>
  <input type="submit" value="Remove">
</form>
</body>
</html>
	`
}

//The action
func DeleteUser() string {
	return `
	<!DOCTYPE html>
<html>
<head>
<title>Delete a user</title>
</head>
<body>
	The user was (maybe) removed ! 
	<form action="/" method="get">
  <input type="submit" value="Back to the main page">
</form>
</body>
</html>
	`
}
