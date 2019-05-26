# testlog

The **testlog** package provides a convenient way to log to the test output when
testing projects that log to a logger (such as the standard libraries
[`"net/http".Server`]).
This ensures that log output doesn't pollute test output and is shown under the
correct test only if the test that generated the log output fails.

```go
import (
	"code.soquee.net/testlog"
)
```

[`"net/http".Server`]: https://godoc.org/net/http#Server.ErrorLog


## License

The package may be used under the terms of the BSD 2-Clause License a copy of
which may be found in the [`LICENSE`] file.

Unless you explicitly state otherwise, any contribution submitted for inclusion
in the work by you shall be licensed as above, without any additional terms or
conditions.

[`LICENSE`]: ./LICENSE
