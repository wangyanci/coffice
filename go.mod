module github.com/wangyanci/coffice

go 1.13

replace (
	github.com/wangyanci/coffice v0.0.0-20190903051234-4220604777aa => ./
	gopkg.in/urfave/cli.v2 v2.3.0 => github.com/urfave/cli/v2 v2.3.0
)

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/lib/pq v1.9.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	gopkg.in/urfave/cli.v2 v2.3.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
