// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: GPL-3.0-only

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: suggest.sql

package db

import (
	"context"
)

const suggestAttribute = `-- name: SuggestAttribute :many
SELECT
  drvattr.attr
FROM
  derivationattr AS drvattr
WHERE
  drvattr.attr LIKE ?
ORDER BY
  drvattr.attr
LIMIT 100
`

func (q *Queries) SuggestAttribute(ctx context.Context, attr string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, suggestAttribute, attr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var attr string
		if err := rows.Scan(&attr); err != nil {
			return nil, err
		}
		items = append(items, attr)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
