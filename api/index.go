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
			<title>Message Routes</title>
		</head>
		<body>
			<h1>Message Routes</h1>
			<p><strong>Message:</strong> Hello!</p>
			<ul>
				<li><code>/api</code>: You are here.</li>
				<li><code>/api/date</code>: Current date.</li>
				<li><code>/api/dbget</code>: Return some all USERS table information from a connected postgres db.</li>
				<li><code>/api/dbinsert/x</code>: Insert into USERS table with name (x).</li>
				<li><code>/api/dbdelete/y</code>: Delete from USERS table with id (y).</li>
				<li><code>/api/dbupdate/y/x</code>: Update from USERS table with id (y) and name (x).</li>
			</ul>
		</body>
		</html>
        `))
}
