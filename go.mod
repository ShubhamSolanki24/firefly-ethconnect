module github.com/hyperledger/firefly-ethconnect

require (
	github.com/Masterminds/semver v1.5.0
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/Shopify/sarama v1.33.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bketelsen/crypt v0.0.4 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/ethereum/go-ethereum v1.10.17
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-openapi/jsonreference v0.20.0
	github.com/go-openapi/spec v0.20.6
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/gorilla/websocket v1.5.0
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/hyperledger/firefly-common v0.1.4
	github.com/icza/dyno v0.0.0-20210726202311-f1bafe5d9996
	github.com/julienschmidt/httprouter v1.3.0
	github.com/kaleido-io/ethbinding v0.0.0-20220405144420-999853435d9e
	github.com/klauspost/compress v1.15.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/microcosm-cc/bluemonday v1.0.18 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/nwaples/rardecode v1.1.3 // indirect
	github.com/oklog/ulid/v2 v2.0.2
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.11.0 // indirect
	github.com/stretchr/testify v1.7.1
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/tidwall/gjson v1.14.1
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.5.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/x-cray/logrus-prefixed-formatter v0.5.2
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/crypto v0.0.0-20220513210258-46612604a0f9 // indirect
	golang.org/x/net v0.0.0-20220516155154-20f960328961 // indirect
	golang.org/x/sys v0.0.0-20220513210249-45d2b4557a2a // indirect
	golang.org/x/term v0.0.0-20220411215600-e5f449aeb171 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20220512140231-539c8e751b99 // indirect
)

replace github.com/kaleido-io/ethbinding => ../ethbinding

go 1.16
