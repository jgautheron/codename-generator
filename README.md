# codename-generator [![Circle CI](https://circleci.com/gh/jgautheron/codename-generator.svg?style=svg)](https://circleci.com/gh/jgautheron/codename-generator) [![GoDoc](https://godoc.org/github.com/jgautheron/codename-generator?status.svg)](https://godoc.org/github.com/jgautheron/codename-generator)

This library written in Golang generates a random code name meant for naming software releases if you run short of inspiration.

Currently based on the pattern "[Superb] [Superhero]".

A few examples of generated codenames:
- Marvelous Meggan
- Prime Doll
- Priceless Ultimo
- Fabulous Longshot
- Wonderful Plazm

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
Instructions are documented in [CONTRIBUTING.md](https://github.com/jgautheron/codename-generator/blob/master/CONTRIBUTING.md).

## Credits
The dictionaries are from:
- https://github.com/sindresorhus/superb
- https://github.com/sindresorhus/superheroes

## License
MIT

## Author
Jonathan Gautheron - jgautheron [A-T] neverblend.in  
https://twitter.com/jgautheron
