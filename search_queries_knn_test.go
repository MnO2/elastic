package elastic

import (
	"encoding/json"
	"testing"
)

func TestKnnQuery(t *testing.T) {
	q := NewKnnQuery(2).Vector(1.2, 2.3, 3.4, 4.5).QueryName("my_vector")
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"knn":{"my_vector":{"k":2,"vector":[1.2,2.3,3.4,4.5]}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

