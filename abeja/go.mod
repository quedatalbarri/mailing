module github.com/quedatalbarri/mailing/abeja

go 1.13

require (
	github.com/caarlos0/env/v6 v6.2.1
	github.com/hanzoai/gochimp3 v0.0.0-20191219204354-bad654ab6826
	github.com/joho/godotenv v1.3.0 // indirect
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/api v0.17.0
)

replace github.com/hanzoai/gochimp3 => github.com/nandanrao/gochimp3 v0.0.0-20200209211044-b930336e8cd7
