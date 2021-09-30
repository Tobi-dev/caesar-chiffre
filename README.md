# Introduction

This is just a simple project for university for the lesson Cryptography.
This programs takes a string and can encode + decode a caesar chiffre.

# What do you need?
You need a working Golang installation. Tested with Go 1.17.1
You can download it here [here](https://golang.org/)

# How to use

Encode:
````
go run main.go --string "something to encode" --rounds 3
````

Decode:
````
go run main.go --string "some encoded string" --decode --rounds 3
````