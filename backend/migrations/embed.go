// Package migrations embeds the goose SQL migrations.
package migrations

import "embed"

// FS contains all goose migration files in this directory.
//
//go:embed *.sql
var FS embed.FS
