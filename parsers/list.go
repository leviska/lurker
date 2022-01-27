package parsers

import "github.com/leviska/lurker/base"

var ParserMap = map[string]base.Parser{
	"genius.com": &Genius{},
}
