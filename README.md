# FETROSHOP API

###### Generate Docs File Using parvez3019/go-swagger3
`go-swagger3 --module-path . --output ./docs/OpenAPI3/source/fetroshop-api.json --schema-without-pkg`

###### Generate Docs File Using swaggo/swag For Web
`swag init -d "app/modules/web,app/model" -g web.go -o docs/web`

###### Generate Docs File Using swaggo/swag For CMS
`swag init -d "app/modules/cms,app/model" -g cms.go -o docs/cms`