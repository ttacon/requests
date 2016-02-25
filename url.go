package requests

import "strings"

type URLBuilder struct {
	components []func(vals map[string]string) string
}

func ParseURLTemplate(templ string) URLBuilder {
	pieces := strings.Split(strings.TrimPrefix(templ, "/"), "/")
	p := URLBuilder{
		components: make([]func(map[string]string) string, len(pieces)),
	}
	for i, piece := range pieces {
		if len(piece) > 0 && piece[0] == ':' {
			p.components[i] = func(piece string) func(vals map[string]string) string {
				return func(vals map[string]string) string {
					return vals[piece[1:]]
				}
			}(piece)
		} else {
			p.components[i] = func(piece string) func(_ map[string]string) string {
				return func(_ map[string]string) string {
					return piece
				}
			}(piece)
		}
	}

	return p
}

func (p URLBuilder) Build(vals map[string]string) string {
	pieces := make([]string, len(p.components))
	for i, fn := range p.components {
		pieces[i] = fn(vals)
	}

	return "/" + strings.Join(pieces, "/")
}
