package main

import (
	"testing"
	"tst2/tst"
)

func TestTST(t *testing.T) {
	testData := []struct {
		key string
		val string
	}{
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
	}

	trie := &tst.Trie{}
	for _, td := range testData {
		trie.Put(td.key, td.val)
		if trie.Get(td.key) == nil {
			t.Errorf("get nil after inserting (%s, %s)\n", td.key, td.val)
		}
	}
}

func BenchmarkTraversal(b *testing.B) {
	testData := []struct {
		key string
		val string
	}{
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Tex-Mex Tilapia", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Honey Sesame Chicken", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cajun-Spiced Pulled Pork", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Cherry Balsamic Pork Chops", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
		{"Hot Honey Barbecue Chicken Legs", ""},
	}
	trie := &tst.Trie{}
	for _, td := range testData {
		trie.Put(td.key, td.val)
	}

	for i := 0; i < b.N; i++ {
		trie.Traverse()
	}
}
