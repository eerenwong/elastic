// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

// SirenJoinQuery perform join index with query.
//
// For details, see
// https://docs.siren.solutions/5.6.8-10.0.0-rc.1/federate/
type SirenJoinQuery struct {
	indices []string
	on      []string
	query   Query
}

// NewSirenJoinQuery creates and initializes a new SirenJoinQuery.
func NewSirenJoinQuery(indices []string, on []string) *SirenJoinQuery {
	return &SirenJoinQuery{indices: indices, on: on}
}

// Query sets the query to use with SirenJoinQuery.
func (s *SirenJoinQuery) Query(query Query) *SirenJoinQuery {
	s.query = query
	return s
}

// Source returns JSON for the query.
func (s *SirenJoinQuery) Source() (interface{}, error) {
	// {"join":{"indices":["profile_doc"],"on":["id.keyword","user_id.keyword"],"request":{"query":{"match":{"message":{"query":"this is a test"}}}}}}
	source := make(map[string]interface{})
	tq := make(map[string]interface{})
	rq := make(map[string]interface{})

	source["join"] = tq

	tq["indices"] = s.indices
	tq["on"] = s.on
	tq["request"] = rq
	if s.query != nil {
		src, err := s.query.Source()
		if err != nil {
			return nil, err
		}
		rq["query"] = src
	}
	return source, nil
}
