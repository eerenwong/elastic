// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSirenJoinQuery(t *testing.T) {
	indices := []string{"profile_doc"}
	on := []string{"id.keyword", "user_id.keyword"}
	query := NewMatchQuery("message", "this is a test")
	q := NewSirenJoinQuery(indices, on).Query(query)
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(src)
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"join":{"indices":["profile_doc"],"on":["id.keyword","user_id.keyword"],"request":{"query":{"match":{"message":{"query":"this is a test"}}}}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
