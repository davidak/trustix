// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: GPL-3.0-only

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: derivations.sql

package db

import (
	"context"
)

const createDerivation = `-- name: CreateDerivation :one
INSERT INTO derivation (drv, system) VALUES (?, ?) RETURNING id, drv, system
`

type CreateDerivationParams struct {
	Drv    string
	System string
}

func (q *Queries) CreateDerivation(ctx context.Context, arg CreateDerivationParams) (Derivation, error) {
	row := q.db.QueryRowContext(ctx, createDerivation, arg.Drv, arg.System)
	var i Derivation
	err := row.Scan(&i.ID, &i.Drv, &i.System)
	return i, err
}

const createDerivationAttr = `-- name: CreateDerivationAttr :exec
INSERT OR IGNORE INTO derivationattr (attr, derivation_id) VALUES (?, ?)
`

type CreateDerivationAttrParams struct {
	Attr         string
	DerivationID int64
}

func (q *Queries) CreateDerivationAttr(ctx context.Context, arg CreateDerivationAttrParams) error {
	_, err := q.db.ExecContext(ctx, createDerivationAttr, arg.Attr, arg.DerivationID)
	return err
}

const createDerivationOutput = `-- name: CreateDerivationOutput :exec
INSERT INTO derivationoutput (output, store_path, derivation_id) VALUES (?, ?, ?)
`

type CreateDerivationOutputParams struct {
	Output       string
	StorePath    string
	DerivationID int64
}

func (q *Queries) CreateDerivationOutput(ctx context.Context, arg CreateDerivationOutputParams) error {
	_, err := q.db.ExecContext(ctx, createDerivationOutput, arg.Output, arg.StorePath, arg.DerivationID)
	return err
}

const createDerivationRefDirect = `-- name: CreateDerivationRefDirect :exec
INSERT OR IGNORE INTO derivationrefdirect (drv_id, referrer_id) VALUES (?, ?)
`

type CreateDerivationRefDirectParams struct {
	DrvID      int64
	ReferrerID int64
}

func (q *Queries) CreateDerivationRefDirect(ctx context.Context, arg CreateDerivationRefDirectParams) error {
	_, err := q.db.ExecContext(ctx, createDerivationRefDirect, arg.DrvID, arg.ReferrerID)
	return err
}

const createDerivationRefRecursive = `-- name: CreateDerivationRefRecursive :exec
INSERT OR IGNORE INTO derivationrefrecursive (drv_id, referrer_id) VALUES (?, ?)
`

type CreateDerivationRefRecursiveParams struct {
	DrvID      int64
	ReferrerID int64
}

func (q *Queries) CreateDerivationRefRecursive(ctx context.Context, arg CreateDerivationRefRecursiveParams) error {
	_, err := q.db.ExecContext(ctx, createDerivationRefRecursive, arg.DrvID, arg.ReferrerID)
	return err
}

const createEval = `-- name: CreateEval :one
INSERT INTO evaluation (commit_sha) VALUES (?) RETURNING id, commit_sha, timestamp
`

func (q *Queries) CreateEval(ctx context.Context, commitSha string) (Evaluation, error) {
	row := q.db.QueryRowContext(ctx, createEval, commitSha)
	var i Evaluation
	err := row.Scan(&i.ID, &i.CommitSha, &i.Timestamp)
	return i, err
}

const getDerivation = `-- name: GetDerivation :one
SELECT id, drv, system FROM derivation
WHERE drv = ? LIMIT 1
`

func (q *Queries) GetDerivation(ctx context.Context, drv string) (Derivation, error) {
	row := q.db.QueryRowContext(ctx, getDerivation, drv)
	var i Derivation
	err := row.Scan(&i.ID, &i.Drv, &i.System)
	return i, err
}

const getDerivationAttr = `-- name: GetDerivationAttr :one
SELECT id, attr, derivation_id FROM derivationattr WHERE derivation_id = ? AND attr = ? LIMIT 1
`

type GetDerivationAttrParams struct {
	DerivationID int64
	Attr         string
}

func (q *Queries) GetDerivationAttr(ctx context.Context, arg GetDerivationAttrParams) (Derivationattr, error) {
	row := q.db.QueryRowContext(ctx, getDerivationAttr, arg.DerivationID, arg.Attr)
	var i Derivationattr
	err := row.Scan(&i.ID, &i.Attr, &i.DerivationID)
	return i, err
}

const getDerivationOutputs = `-- name: GetDerivationOutputs :many
SELECT id, output, store_path, derivation_id FROM derivationoutput WHERE store_path = ?
`

func (q *Queries) GetDerivationOutputs(ctx context.Context, storePath string) ([]Derivationoutput, error) {
	rows, err := q.db.QueryContext(ctx, getDerivationOutputs, storePath)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Derivationoutput
	for rows.Next() {
		var i Derivationoutput
		if err := rows.Scan(
			&i.ID,
			&i.Output,
			&i.StorePath,
			&i.DerivationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDerivationOutputsByDerivationID = `-- name: GetDerivationOutputsByDerivationID :many
SELECT id, output, store_path, derivation_id FROM derivationoutput WHERE derivation_id = ?
`

func (q *Queries) GetDerivationOutputsByDerivationID(ctx context.Context, derivationID int64) ([]Derivationoutput, error) {
	rows, err := q.db.QueryContext(ctx, getDerivationOutputsByDerivationID, derivationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Derivationoutput
	for rows.Next() {
		var i Derivationoutput
		if err := rows.Scan(
			&i.ID,
			&i.Output,
			&i.StorePath,
			&i.DerivationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEval = `-- name: GetEval :one
SELECT id, commit_sha, timestamp FROM evaluation
WHERE commit_sha = ? LIMIT 1
`

func (q *Queries) GetEval(ctx context.Context, commitSha string) (Evaluation, error) {
	row := q.db.QueryRowContext(ctx, getEval, commitSha)
	var i Evaluation
	err := row.Scan(&i.ID, &i.CommitSha, &i.Timestamp)
	return i, err
}
