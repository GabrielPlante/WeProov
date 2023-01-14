package view

//The form for the user creation
func CreateUser() string {
	return `
	<!DOCTYPE html>
<html>
<head>
<title>Create a user</title>
</head>
<body>
	<form action="/postuser" method="post">

  <label for="username">Username:</label>
  <input type="text" id="username" name="username"><br><br>

  <label for="password">Password:</label>
  <input type="text" id="password" name="password"><br><br>

  <label for="email">Email:</label>
  <input type="text" id="Email" name="Email"><br><br>

  <input type="submit" value="Submit">
</form>
</body>
</html>
	`
}
