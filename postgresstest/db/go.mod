module github.com/fritzkeyzer/test/postgresstest/db

go 1.16

require (
	github.com/fritzkeyzer/test/postgresstest/types v0.0.0
	github.com/jackc/pgx/v4 v4.14.1
	github.com/lib/pq v1.10.4

)

replace github.com/fritzkeyzer/test/postgresstest/types => ../types
