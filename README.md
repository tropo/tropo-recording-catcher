# tropo-recording-catcher

## Overview

This is an example project for receiving an HTTP POST from a [Tropo](http://tropo.com) [record](https://www.tropo.com/docs/scripting/record) or [startCallRecording](https://www.tropo.com/docs/scripting/startcallrecording) methods. The server will then PUT those files to an AWS S3 bucket.

Note: This server accepts the POST from Tropo and immediately returns, using a Go routine to manage the longer running I/O with AWS S3.

## Configuration

These environment variables must be set:

```
export TROPO_AWS_KEY=<aws key>
export TROPO_AWS_SECRET=<aws secret>
export TROPO_AWS_BUCKET=<aws s3 bucket>
export TROPO_PORT=3000 # port to run the server on
```

## Executing

* Windows
	./tropo-recording-catcher.exe

* Linux
	./tropo-recording-catcher.linux

* OSX
	./tropo-recording-catcher.osx

## Building

Simply do ./build.sh within the package. All executables will then be in the /pkg dirctory within the project.

## License

The MIT License (MIT)

Copyright (c) 2014 Tropo, INC

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.