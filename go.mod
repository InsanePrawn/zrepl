module github.com/zrepl/zrepl

go 1.12

require (
	github.com/DATA-DOG/go-sqlmock v1.3.3 // indirect
	github.com/fatih/color v1.7.0
	github.com/gdamore/tcell v1.4.0
	github.com/gdamore/tcell/v2 v2.1.0
	github.com/gitchander/permutation v0.0.0-20181107151852-9e56b92e9909
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-sql-driver/mysql v1.4.1-0.20190907122137-b2c03bcae3d4
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/jinzhu/copier v0.0.0-20170922082739-db4671f3a9b8
	github.com/kisielk/gotool v1.0.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kr/pretty v0.1.0
	github.com/lib/pq v1.2.0
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.8
	github.com/montanaflynn/stats v0.5.0
	github.com/onsi/ginkgo v1.10.2 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pkg/profile v1.2.1
	github.com/problame/go-netssh v0.0.0-20200601114649-26439f9f0dc5
	github.com/prometheus/client_golang v1.2.1
	github.com/prometheus/common v0.7.0
	github.com/rivo/tview v0.0.0-20210125085121-dbc1f32bb1d0
	github.com/sergi/go-diff v1.0.1-0.20180205163309-da645544ed44 // go1.12 thinks it needs this
	github.com/spf13/cobra v0.0.2
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.5.1
	github.com/willf/bitset v1.1.10
	github.com/yudai/gojsondiff v0.0.0-20170107030110-7b1b7adf999d
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect; go1.12 thinks it needs this
	github.com/yudai/pp v2.0.1+incompatible // indirect
	github.com/zrepl/yaml-config v0.0.0-20191220194647-cbb6b0cf4bdd
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135
	gonum.org/v1/gonum v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20210122163508-8081c04a3579 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

// https://github.com/rivo/tview/pull/523#issuecomment-778837580
replace github.com/rivo/tview => github.com/berrange/tview v0.0.0-20201105134140-da9053a5590a
