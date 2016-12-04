// Copyright 2016 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package scanner_test

import (
	"testing"

	"unicode/utf8"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkString(t *testing.B) {
	data := []byte(`"hello world"`)

	for i := 0; i < t.N; i++ {
		end, err := scanner.String(data, 0)
		if err != nil {
			t.FailNow()
			return
		}

		if end == 0 {
			t.FailNow()
			return
		}
	}
}

func TestString(t *testing.T) {
	Convey("Verify String", t, func() {
		testCases := map[string]struct {
			In     string
			Out    string
			HasErr bool
		}{
			"simple": {
				In:  `"hello"`,
				Out: `"hello"`,
			},
			"array": {
				In:  `"hello", "world"`,
				Out: `"hello"`,
			},
			"escaped": {
				In:  `"hello\"\"world"`,
				Out: `"hello\"\"world"`,
			},
			"unclosed": {
				In:     `"hello`,
				HasErr: true,
			},
			"unclosed escape": {
				In:     `"hello\"`,
				HasErr: true,
			},
			"utf8": {
				In:  `"生日快乐"`,
				Out: `"生日快乐"`,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				end, err := scanner.String([]byte(tc.In), 0)
				if tc.HasErr {
					So(err, ShouldNotBeNil)
				} else {
					data := tc.In[0:end]
					So(string(data), ShouldEqual, tc.Out)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestDecode(t *testing.T) {
	Convey("Verify can decode", t, func() {
		v := ""
		_, size := utf8.DecodeRune([]byte(v))
		So(size, ShouldEqual, 0)
	})
}
