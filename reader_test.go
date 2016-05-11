// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptyy

import (
	"strings"
	"testing"

	"github.com/chai2010/ptyy/static"
)

func TestLoadJsonFile(t *testing.T) {
	db, err := LoadJsonFile("./static/hospital_list.20160508.json")
	if err != nil {
		t.Fatal(err)
	}

	info := db["信阳博士医院"]
	if len(info.WebSite) != 1 || info.WebSite[0] != "http://www.bbrmyy.com" {
		t.Fatalf("bad info: %#v", info)
	}
	if len(info.Tel) != 1 || info.Tel[0] != "03763222555" {
		t.Fatalf("bad info: %#v", info)
	}
}

func TestReadJsonFrom(t *testing.T) {
	strData := static.Files["hospital_list.20160508.json"]
	db, err := ReadJsonFrom(strings.NewReader(strData))
	if err != nil {
		t.Fatal(err)
	}

	info := db["信阳博士医院"]
	if len(info.WebSite) != 1 || info.WebSite[0] != "http://www.bbrmyy.com" {
		t.Fatalf("bad info: %#v", info)
	}
	if len(info.Tel) != 1 || info.Tel[0] != "03763222555" {
		t.Fatalf("bad info: %#v", info)
	}
}
