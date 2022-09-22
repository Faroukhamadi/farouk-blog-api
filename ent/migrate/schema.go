// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContentsColumns holds the columns for the "contents" table.
	ContentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "text", Type: field.TypeString},
		{Name: "post_content", Type: field.TypeInt, Nullable: true},
	}
	// ContentsTable holds the schema information for the "contents" table.
	ContentsTable = &schema.Table{
		Name:       "contents",
		Columns:    ContentsColumns,
		PrimaryKey: []*schema.Column{ContentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contents_posts_content",
				Columns:    []*schema.Column{ContentsColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "meta_title", Type: field.TypeString},
		{Name: "foreword", Type: field.TypeString},
		{Name: "contents", Type: field.TypeString},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContentsTable,
		PostsTable,
	}
)

func init() {
	ContentsTable.ForeignKeys[0].RefTable = PostsTable
}
