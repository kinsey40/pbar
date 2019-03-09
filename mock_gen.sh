# Script to auto-generate the mocks used in the testing of this library

${GOPATH}/bin/mockgen -source=render/clock.go -destination=mocks/clock_mock.go -package=mocks
${GOPATH}/bin/mockgen -source=render/write.go -destination=mocks/write_mock.go -package=mocks
${GOPATH}/bin/mockgen -source=render/settings.go -destination=mocks/settings_mock.go -package=mocks
${GOPATH}/bin/mockgen -source=render/values.go -destination=mocks/values_mock.go -package=mocks
