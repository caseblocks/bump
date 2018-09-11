# bump

Increases a counter in VERSION files. Output is to STDOUT and the original VERSION file is left unmodified.

# Usage

Assuming a VERSION file containting `1.2.3.` in the current folder:

    bump major          # Output 2.0.0

    bump minor          # Output 1.3.0

    bump patch          # Output 1.2.4

    bump somethingelse  # Output 1.2.3

To modify a VERSION file:

    bump major > VERSION  # overwrites current file

# Thanks

Special thanks to https://github.com/giantswarm/semver-bump for the inspiration.

# License

MIT License

Copyright (c) 2017-2018 EmergeAdapt Ltd.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
