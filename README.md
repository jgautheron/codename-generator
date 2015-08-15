# codename-generator [![Circle CI](https://circleci.com/gh/jgautheron/codename-generator.svg?style=svg)](https://circleci.com/gh/jgautheron/codename-generator)

This library written in Golang generates a random code name meant for naming software releases if you run short of inspiration.

Currently based on the pattern "[Superb] [Superhero]".

The words were taken from:
- https://github.com/sindresorhus/superb
- https://github.com/sindresorhus/superheroes

The words are included in the binary with go-bindata, to regenerate the dictionary, run in the codename folder:
```
# go-bindata -o words.go -pkg codename data/
```

## Getting started
First, download the project.
```
go get -u github.com/jgautheron/codename-generator
```

Then in your code:
```go
import "github.com/jgautheron/codename-generator"

// Sanitized returns a safe string, ex. "awesome-hero".
cn, err := codename.Get(codename.Sanitized)
```

## Contributing
Contributions are most welcome, especially for enlarging the words pool.

## License
MIT

## Author
Jonathan Gautheron - jgautheron [A-T] neverblend.in  
https://twitter.com/jgautheron
