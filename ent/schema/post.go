package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/farouk-blog-api/ent/schema/mixin/timestamps"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("meta_title"),
		field.String("foreword"),
		field.String("contents"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contents", Content.Type),
	}
}
