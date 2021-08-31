module hello

go 1.17

replace example.com/greetings => ../greetings

replace kchen.com/math => ../math

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	kchen.com/math v0.0.0-00010101000000-000000000000
)
