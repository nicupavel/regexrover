// Copyright 2024 Nicu Pavel <npavel@linuxconsulting.ro>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package regexrover

import "log"

type Logger func(format string, v ...interface{})

var infoLog Logger = log.Printf
var debugLog Logger = log.Printf
var errorLog Logger = log.Printf
var miscLog Logger = EmptyLogger

func EmptyLogger(format string, v ...interface{}) {
	_, _ = format, v // disable unused warning
}

func SetLogger(info Logger, debug Logger, misc Logger, error Logger) {
	infoLog = info
	debugLog = debug
	errorLog = error
	miscLog = misc
}
