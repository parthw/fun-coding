module example.com/importing-hello-modules

go 1.15

replace example.com/hello => ../hello-modules

require (
	example.com/hello v0.0.0-00010101000000-000000000000
	rsc.io/quote/v3 v3.1.0
)
