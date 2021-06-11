package benchmarks

import (
	"io/ioutil"
	"time"

	"github.com/Telefonica/govice"
)

func newGovice() *govice.Logger {
	logger := govice.NewLogger()
	logger.SetLevel("debug")
	logger.SetWriter(ioutil.Discard)
	return logger
}
func newDisabledGovice() *govice.Logger {
	logger := govice.NewLogger()
	logger.SetLevel("fatal")
	return logger
}
func setGoviceContext(l *govice.Logger) {
	type FakeContext struct {
		Int int `json"int"`
		Ints []int `json"int"`
		Str string `json"string"`
		Strs []string `json"strings"`
		Time time.Time `json"time"`
		Times []time.Time `json"times"`
		User1 *user `json"user1"`
		User2 *user `json"user2"`
		Users users `json"users"`
		Error error `json"err"`
	}
	context := FakeContext{
		Int: _tenInts[0],
		Ints: _tenInts,
		Str: _tenStrings[0],
		Strs: _tenStrings,
		Time: _tenTimes[0],
		Times: _tenTimes,
		User1: _oneUser,
		User2: _oneUser,
		Users: _tenUsers,
		Error: errExample,

	}
	l.SetLogContext(context)
}
func fakeGoviceContext() interface{} {
	type FakeContext struct {
		Int int `json"int"`
		Ints []int `json"int"`
		Str string `json"string"`
		Strs []string `json"strings"`
		Time time.Time `json"time"`
		Times []time.Time `json"times"`
		User1 *user `json"user1"`
		User2 *user `json"user2"`
		Users users `json"users"`
		Error error `json"err"`
	}
	context := FakeContext{
		Int: _tenInts[0],
		Ints: _tenInts,
		Str: _tenStrings[0],
		Strs: _tenStrings,
		Time: _tenTimes[0],
		Times: _tenTimes,
		User1: _oneUser,
		User2: _oneUser,
		Users: _tenUsers,
		Error: errExample,

	}
	return context
}
