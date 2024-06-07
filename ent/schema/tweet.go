package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").
			NotEmpty(),
			field.String("author_id"),
			field.String("id").Immutable().NotEmpty().Unique(),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("tweets").
			Unique().Field("author_id").Required(),
	}
}
