module main

go 1.17

replace kchen.com/greetings => ../greetings

replace kchen.com/math => ../math

require (
	kchen.com/greetings v0.0.0-00010101000000-000000000000
	kchen.com/math v0.0.0-00010101000000-000000000000
)
