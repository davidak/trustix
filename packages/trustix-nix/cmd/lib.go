// Copyright (C) 2021 Tweag IO
// Copyright © 2020-2022 The Trustix Authors
//
// SPDX-License-Identifier: GPL-3.0-only

package cmd

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/nix-community/trustix/packages/trustix-nix/schema"
	"github.com/nix-community/trustix/packages/trustix-proto/api"
	log "github.com/sirupsen/logrus"
)

var NixB32Encoding = base32.NewEncoding("0123456789abcdfghijklmnpqrsvwxyz")

func createKVPair(storePath string) (*api.KeyValuePair, error) {

	if storePath == "" {
		return nil, fmt.Errorf("Empty input store path")
	}

	tmpDir, err := ioutil.TempDir("", "nix-trustix")
	if err != nil {
		return nil, err
	}
	err = os.RemoveAll(tmpDir)
	if err != nil {
		return nil, err
	}

	var storeHash []byte
	{
		storeHashStr := strings.Split(filepath.Base(storePath), "-")[0]
		storeHash, err = NixB32Encoding.DecodeString(storeHashStr)
		if err != nil {
			log.Fatal(err)
		}
		if len(storeHash) == 0 {
			log.Fatal("Empty decoded store path hash")
		}
	}

	var narinfo *schema.NarInfo
	{
		out, err := exec.Command("nix", "path-info", "--json", storePath).Output()
		if err != nil {
			return nil, err
		}

		var narinfos []*schema.NarInfo
		err = json.Unmarshal(out, &narinfos)
		if err != nil {
			log.Fatalf("Could not get path info: %v", err)
		}

		if len(narinfos) != 1 {
			log.Fatalf("Unexpected number of narinfos returned: %d", len(narinfos))
		}

		narinfo = narinfos[0]

		sort.Strings(narinfo.References)
	}

	log.WithFields(log.Fields{
		"storePath": storePath,
	}).Debug("Submitting mapping")

	narinfoBytes, err := json.Marshal(narinfo)
	if err != nil {
		log.Fatalf("Could not marshal narinfo: %v", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return &api.KeyValuePair{
		Key:   storeHash,
		Value: narinfoBytes,
	}, nil
}
