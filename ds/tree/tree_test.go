package tree_test

import (
	"testing"

	"github.com/eng618/go-eng/ds/tree"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tr := tree.New()
	if tr.Root != nil {
		t.Errorf("New() should create an empty tree, got Root = %v", tr.Root)
	}
}

func TestNewWithRoot(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value int
	}{
		{name: "zero value", value: 0},
		{name: "positive value", value: 10},
		{name: "negative value", value: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := tree.NewWithRoot(tt.value)
			if tr.Root == nil {
				t.Errorf("NewWithRoot(%d) should create a tree with root, got nil", tt.value)
			}
			if tr.Root.Value != tt.value {
				t.Errorf("NewWithRoot(%d) root value = %d, want %d", tt.value, tr.Root.Value, tt.value)
			}
		})
	}
}

func TestTree_Insert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		values []int
	}{
		{name: "single value", values: []int{5}},
		{name: "multiple values", values: []int{5, 3, 7, 1, 9}},
		{name: "ascending order", values: []int{1, 2, 3, 4, 5}},
		{name: "descending order", values: []int{5, 4, 3, 2, 1}},
		{name: "duplicates", values: []int{5, 3, 5, 7, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := tree.New()
			for _, v := range tt.values {
				tr.Insert(v)
			}
			if tr.Root == nil {
				t.Errorf("Insert() should create root node")
			}
		})
	}
}

func TestTree_Search(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		values     []int
		searchFor  int
		wantFound  bool
	}{
		{name: "empty tree", values: []int{}, searchFor: 5, wantFound: false},
		{name: "find root", values: []int{5}, searchFor: 5, wantFound: true},
		{name: "find left child", values: []int{5, 3, 7}, searchFor: 3, wantFound: true},
		{name: "find right child", values: []int{5, 3, 7}, searchFor: 7, wantFound: true},
		{name: "value not found", values: []int{5, 3, 7}, searchFor: 10, wantFound: false},
		{name: "find in larger tree", values: []int{10, 5, 15, 3, 7, 12, 20}, searchFor: 12, wantFound: true},
		{name: "negative value found", values: []int{10, -5, 15}, searchFor: -5, wantFound: true},
		{name: "negative value not found", values: []int{10, 5, 15}, searchFor: -5, wantFound: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := tree.New()
			for _, v := range tt.values {
				tr.Insert(v)
			}
			got := tr.Search(tt.searchFor)
			if got != tt.wantFound {
				t.Errorf("Search(%d) = %v, want %v", tt.searchFor, got, tt.wantFound)
			}
		})
	}
}

func TestTree_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		values        []int
		deleteValue   int
		wantErr       bool
		searchAfter   int
		wantFoundAfter bool
	}{
		{
			name:          "delete from empty tree",
			values:        []int{},
			deleteValue:   5,
			wantErr:       true,
			searchAfter:   5,
			wantFoundAfter: false,
		},
		{
			name:          "delete root with no children",
			values:        []int{5},
			deleteValue:   5,
			wantErr:       false,
			searchAfter:   5,
			wantFoundAfter: false,
		},
		{
			name:          "delete leaf node",
			values:        []int{5, 3, 7},
			deleteValue:   3,
			wantErr:       false,
			searchAfter:   3,
			wantFoundAfter: false,
		},
		{
			name:          "delete node with one child",
			values:        []int{5, 3, 7, 1},
			deleteValue:   3,
			wantErr:       false,
			searchAfter:   3,
			wantFoundAfter: false,
		},
		{
			name:          "delete node with two children",
			values:        []int{10, 5, 15, 3, 7, 12, 20},
			deleteValue:   5,
			wantErr:       false,
			searchAfter:   5,
			wantFoundAfter: false,
		},
		{
			name:          "delete non-existent value",
			values:        []int{5, 3, 7},
			deleteValue:   10,
			wantErr:       true,
			searchAfter:   5,
			wantFoundAfter: true,
		},
		{
			name:          "delete root with two children",
			values:        []int{10, 5, 15, 3, 7, 12, 20},
			deleteValue:   10,
			wantErr:       false,
			searchAfter:   10,
			wantFoundAfter: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := tree.New()
			for _, v := range tt.values {
				tr.Insert(v)
			}
			err := tr.Delete(tt.deleteValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete(%d) error = %v, wantErr %v", tt.deleteValue, err, tt.wantErr)
			}
			
			// Verify the value is deleted
			got := tr.Search(tt.searchAfter)
			if got != tt.wantFoundAfter {
				t.Errorf("After Delete(%d), Search(%d) = %v, want %v", tt.deleteValue, tt.searchAfter, got, tt.wantFoundAfter)
			}
		})
	}
}

// TestTree_DeleteAndInsert tests the combination of operations
func TestTree_DeleteAndInsert(t *testing.T) {
	t.Parallel()

	tr := tree.New()
	
	// Build a tree
	values := []int{10, 5, 15, 3, 7, 12, 20}
	for _, v := range values {
		tr.Insert(v)
	}
	
	// Delete a node
	if err := tr.Delete(5); err != nil {
		t.Errorf("Delete(5) unexpected error: %v", err)
	}
	
	// Verify deletion
	if tr.Search(5) {
		t.Errorf("Search(5) should return false after deletion")
	}
	
	// Verify other nodes still exist
	for _, v := range []int{10, 15, 3, 7, 12, 20} {
		if !tr.Search(v) {
			t.Errorf("Search(%d) should return true", v)
		}
	}
	
	// Re-insert the deleted value
	tr.Insert(5)
	
	// Verify re-insertion
	if !tr.Search(5) {
		t.Errorf("Search(5) should return true after re-insertion")
	}
}

func BenchmarkTree_Insert(b *testing.B) {
	tr := tree.New()
	for i := 0; i < b.N; i++ {
		tr.Insert(i)
	}
}

func BenchmarkTree_Search(b *testing.B) {
	tr := tree.New()
	for i := 0; i < 1000; i++ {
		tr.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Search(i % 1000)
	}
}

func BenchmarkTree_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tr := tree.New()
		for j := 0; j < 1000; j++ {
			tr.Insert(j)
		}
		b.StartTimer()
		_ = tr.Delete(i % 1000)
	}
}
