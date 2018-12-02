# Gofile, the idiomatic way to manage go packages

## Alpha

`gf` is in alpha and as such will likely receive changes and bug fixes over time. Early adopters welcome, but buyer beware!

## What is a Gofile?

Gofile is the revolutionary new way to manage your packages, but removing all the un-idiomatic use of distinct files, and compressing all that code into one single Gofile!

How many times have you been looking for that special file in your package, only to spend hours scrolling on your trackpad and or mouse in vain?

My hypothosis is: Too many!

That's why I wrote `gf`, to support the community in embracing the time saving power of Gofile!

Golang is all about writing code, not searching for files! We need to get things done, not give "Big File System" more of our hard earned money!

Please enjoy `gf` as your gateway to Golang productivity!

## Usage

* Install the package with your choice of go get invocations
* Pass the fully qualified path to your go package
    * To be Module-Ready, Gofile does not assume something lives under `$GOPATH`

`gf` is intended to be used in build pipelines, where raw go source code can conveniently be bundled into a single Gofile.

Truly idiomatic gophers will want to use `gf` before each commit, when they may have accidently added files without realizing it.

### Flags

* a - reserved for future use
* b - reserved for future use
* c - reserved for future use
* d - dry run, do not delete the original package files (however this leaves the package unable to compile)
* e - reserved for future use
* f - reserved for future use
* g - reserved for future use
* h - reserved for future use
* i - reserved for future use
* j - reserved for future use
* k - reserved for future use
* l - reserved for future use
* m - reserved for future use
* n - reserved for future use
* o - organize certain package level fixtures at the top of the resulting Gofile, such as package level constants and variables
* p - (required) The fully qualified path to the package. To be module-ready, this stands neither for path nor package, but an original word that I invented called `pathkage`
* q - reserved for future use
* r - reserved for future use
* s - reserved for future use
* t - reserved for future use
* u - reserved for future use
* v - reserved for future use
* w - reserved for future use
* x - reserved for future use
* y - reserved for future use
* z - reserved for future use

## Roadmap

* Delete all the unidiomatic files after generating Gofile
* Perform same work, but for idiomatic Gofile_test testing results
* Adoption into the go tool itself