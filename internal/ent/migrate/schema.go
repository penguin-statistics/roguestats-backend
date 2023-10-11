// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "event_id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_agent", Type: field.TypeString},
		{Name: "content", Type: field.TypeJSON},
		{Name: "research_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_researches_events",
				Columns:    []*schema.Column{EventsColumns[4]},
				RefColumns: []*schema.Column{ResearchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "events_users_events",
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "event_created_at",
				Unique:  false,
				Columns: []*schema.Column{EventsColumns[1]},
			},
			{
				Name:    "event_user_agent",
				Unique:  false,
				Columns: []*schema.Column{EventsColumns[2]},
			},
			{
				Name:    "event_user_id",
				Unique:  false,
				Columns: []*schema.Column{EventsColumns[5]},
			},
			{
				Name:    "event_research_id",
				Unique:  false,
				Columns: []*schema.Column{EventsColumns[4]},
			},
		},
	}
	// QueryPresetsColumns holds the columns for the "query_presets" table.
	QueryPresetsColumns = []*schema.Column{
		{Name: "query_preset_id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "where", Type: field.TypeJSON},
		{Name: "mapping", Type: field.TypeString},
		{Name: "research_id", Type: field.TypeString},
	}
	// QueryPresetsTable holds the schema information for the "query_presets" table.
	QueryPresetsTable = &schema.Table{
		Name:       "query_presets",
		Columns:    QueryPresetsColumns,
		PrimaryKey: []*schema.Column{QueryPresetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "query_presets_researches_query_presets",
				Columns:    []*schema.Column{QueryPresetsColumns[4]},
				RefColumns: []*schema.Column{ResearchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ResearchesColumns holds the columns for the "researches" table.
	ResearchesColumns = []*schema.Column{
		{Name: "research_id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "schema", Type: field.TypeBytes},
	}
	// ResearchesTable holds the schema information for the "researches" table.
	ResearchesTable = &schema.Table{
		Name:       "researches",
		Columns:    ResearchesColumns,
		PrimaryKey: []*schema.Column{ResearchesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Size: 64},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "credential", Type: field.TypeString, Size: 64},
		{Name: "attributes", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		QueryPresetsTable,
		ResearchesTable,
		UsersTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = ResearchesTable
	EventsTable.ForeignKeys[1].RefTable = UsersTable
	QueryPresetsTable.ForeignKeys[0].RefTable = ResearchesTable
}
