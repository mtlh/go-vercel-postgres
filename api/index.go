package api

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		`
        <!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>API Routes</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					margin: 0;
					padding: 0;
					background-color: #f4f4f4;
					color: #333;
				}
				.container {
					max-width: 800px;
					margin: 20px auto;
					padding: 20px;
					background-color: #fff;
					border-radius: 5px;
					box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
				}
				h1 {
					color: #007bff;
				}
				ul {
					list-style: none;
					padding: 0;
				}
				li {
					margin-bottom: 10px;
				}
				code {
					background-color: #f8f9fa;
					padding: 2px 5px;
					border-radius: 3px;
				}
				a {
					color: #007bff;
					text-decoration: none;
				}
				a:hover {
					text-decoration: underline;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>API Routes</h1>
				<ul>
					<li><code><a href="/api">/api</a></code></li>
					<li>You are here.</li>
					<li><code><a href="/api/date">/api/date</a></code></li>
					<li>Current date.</li>
					<li><code><a href="/api/dbget">/api/dbget</a></code></li>
					<li>Return some all USERS table information from a connected postgres db.</li>
					<li><code><a href="/api/dbinsert/demo">/api/dbinsert/x</a></code></li>
					<li>Insert into USERS table with name (x).</li>
					<li><code><a href="/api/dbdelete/4">/api/dbdelete/y</a></code></li>
					<li>Delete from USERS table with id (y).</li>
					<li><code><a href="/api/dbupdate/6/demo">/api/dbupdate/y/x</a></code></li>
					<li>Update from USERS table with id (y) and name (x).</li>
				</ul>
			</div>
		</body>
		</html>
        `))
}
