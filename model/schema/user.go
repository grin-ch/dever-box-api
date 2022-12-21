package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Comment("自增主键"),
		field.String("nickname").Unique().NotEmpty().Comment("昵称"),
		field.String("account").Unique().NotEmpty().Immutable().Comment("账号"),
		field.String("password").NotEmpty().Comment("密码"),
		field.Time("reg_time").Default(time.Now).Immutable().Comment("注册时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
