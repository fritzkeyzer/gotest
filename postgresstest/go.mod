module github.com/fritzkeyzer/test/postgresstest

go 1.16

require (
	github.com/fritzkeyzer/test/postgresstest/db v0.0.0
	github.com/fritzkeyzer/test/postgresstest/types v0.0.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect

)

replace (
	github.com/fritzkeyzer/test/postgresstest/db => ./db
	github.com/fritzkeyzer/test/postgresstest/types => ./types
)
