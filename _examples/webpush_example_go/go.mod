module github.com/unifiedpush/go_dbus_connector/_examples/basic_example

go 1.16

require (
	github.com/gen2brain/beeep v0.0.0-20210529141713-5586760f0cc1
	github.com/xakep666/ecego v0.1.0
	unifiedpush.org/go/dbus_connector v0.1.1
)

replace unifiedpush.org/go/dbus_connector => ../../

replace github.com/xakep666/ecego => github.com/karmanyaahm/ecego v0.1.1-0.20220403012549-48061440c9e5
