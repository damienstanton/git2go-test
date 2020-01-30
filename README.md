### Steps to reproduce

- Clone this repo
- In `libgit-src` dir, extract the zipfile, then
    - run  `mkdir build && cmake -B build -S . && cmake --build build` to generate buildfiles
    - cd into `build` and run `make install`. Now you have `libgit 0.27` installed on your system
- In your Go module-based project, run `go get gopkg.in/libgit2/git2go.v27`
- Include the source code similar to this one, and point it at a known local git repo for testing.

Given the above steps have been followed, the test should pass
```sh
$ make test
# === RUN   TestMain
# --- PASS: TestMain (0.00s)
# PASS
# coverage: 0.0% of statements
# ok      github.com/damienstanton/git2go-test    0.112s
```
© 2020 Damien Stanton

See LICENSE for details.

