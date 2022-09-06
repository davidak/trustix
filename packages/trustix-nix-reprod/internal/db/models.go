// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: GPL-3.0-only

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Derivation struct {
	ID     int64
	Drv    string
	System string
}

type Derivationattr struct {
	ID           int64
	Attr         string
	DerivationID int64
}

type Derivationoutput struct {
	ID           int64
	Output       string
	StorePath    string
	DerivationID int64
}

type Derivationoutputresult struct {
	ID         int64
	OutputHash string
	StorePath  string
	LogID      int64
}

type Derivationrefdirect struct {
	ID         int64
	DrvID      int64
	ReferrerID int64
}

type Derivationrefrecursive struct {
	ID         int64
	DrvID      int64
	ReferrerID int64
}

type Evaluation struct {
	ID        int64
	CommitSha string
	Timestamp time.Time
}

type Log struct {
	ID       int64
	LogID    string
	TreeSize int64
}
