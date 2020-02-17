# go-debouncer

Simple thread safe debouncer in Go

## Install

```
go get github.com/skydive-project/go-debouncer
```

## Usage

```
package main

import (
        "fmt"
        "time"
        "github.com/skydive-project/go-debouncer"
)

func myCallback() {
        fmt.Printf("Let's do something\n")
}

func main() {
        // Creates and returns a new debouncer that will postpone
        // the execution of a function myCallback after 3 seconds
        // have elapsed since the last time it was invoked.
        debouncer := debouncer.New(time.Second, myCallback)

        // Start the debouncer
        debouncer.Start()

        // Trigger the execution of the callback
        debouncer.Call()

        // The following calls will not cause the callback to
        // be executed more than once
        for i:=0; i<10; i++ {
                debouncer.Call()
        }

        // Wait a bit to let the debouncer
        time.Sleep(3*time.Second)

        // Stop the debouncer
        debouncer.Stop()
}
```

## License

This software is licensed under the Apache License, Version 2.0 (the
"License"); you may not use this software except in compliance with the
License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

