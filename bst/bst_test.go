package bst

import (
	"reflect"
	"testing"
)

var demo = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	},
	Right: &TreeNode{Val: 4},
}

func TestInitBST(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "empty tree",
			args: args{
				values: []int{3, 2, 1, 4},
			},
			want: demo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitBST(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "empty tree",
			args: args{
				val: 1,
			},
			want: &TreeNode{Val: 1},
		},
		{
			name: "tree exists",
			args: args{
				root: &TreeNode{Val: 2},
				val:  1,
			},
			want: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
		},
		{
			name: "repeat",
			args: args{
				root: &TreeNode{Val: 2},
				val:  2,
			},
			want: &TreeNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				root: InitBST([]int{1, 5, 7, 9, 2, 8, 8, 10}),
			},
			want: []int{1, 2, 5, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExist(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exists",
			args: args{
				root: InitBST([]int{1, 5, 7, 9, 2, 8, 8, 10}),
				val:  5,
			},
			want: true,
		},
		{
			name: "not exists",
			args: args{
				root: InitBST([]int{1, 5, 7, 9, 2, 8, 8, 10}),
				val:  6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.args.root, tt.args.val); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "delete exists - case1",
			args: args{
				root: demo,
				val:  1,
			},
			want: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		},
		{
			name: "delete exists - case2",
			args: args{
				root: demo,
				val:  2,
			},
			want: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}},
		},
		{
			name: "delete exists - case3",
			args: args{
				root: demo,
				val:  3,
			},
			want: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
		},
		{
			name: "delete not exists",
			args: args{
				root: demo,
				val:  9,
			},
			want: demo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is BST",
			args: args{
				root: demo,
			},
			want: true,
		},
		{
			name: "null is BST",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "is not BST",
			args: args{
				root: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBST(tt.args.root); got != tt.want {
				t.Errorf("IsBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
