package main

import (
	"reflect"
	"testing"
)

func Test_detectItemTypeError(t *testing.T) {
	type args struct {
		rucksack string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{name: "rucksack 1", args: args{rucksack: "vJrwpWtwJgWrhcsFMMfFFhFp"}, want: 'p'},
		{name: "rucksack 2", args: args{rucksack: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}, want: 'L'},
		{name: "rucksack 3", args: args{rucksack: "PmmdzqPrVvPwwTWBwg"}, want: 'P'},
		{name: "rucksack 4", args: args{rucksack: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"}, want: 'v'},
		{name: "rucksack 5", args: args{rucksack: "ttgJtRGJQctTZtZT"}, want: 't'},
		{name: "rucksack 6", args: args{rucksack: "CrZsJsPPZsGzwwsLwLmpwMDw"}, want: 's'},
		{name: "correct format rucksack", args: args{rucksack: "CrZsJsPPZsGzwwaLwLmpwMDw"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectItemTypeError(tt.args.rucksack); got != tt.want {
				t.Errorf("detectItemTypeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convItemTypeToPriority(t *testing.T) {
	type args struct {
		itemType byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Lowercase", args: args{itemType: 'p'}, want: 16, wantErr: false},
		{name: "Uppercase", args: args{itemType: 'L'}, want: 38, wantErr: false},
		{name: "Not Item Type", args: args{itemType: '6'}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convItemTypeToPriority(tt.args.itemType)
			if (err != nil) != tt.wantErr {
				t.Errorf("convItemTypeToPriority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convItemTypeToPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfPriorities(t *testing.T) {
	type args struct {
		itemTypes []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Case 1", args: args{itemTypes: []byte{'p', 'L', 'P', 'v', 't', 's'}}, want: 157},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfPriorities(tt.args.itemTypes); got != tt.want {
				t.Errorf("sumOfPriorities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findErrorInRucksacks(t *testing.T) {
	type args struct {
		rucksacks []string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "Without correct format rucksack",
			args: args{
				rucksacks: []string{
					"vJrwpWtwJgWrhcsFMMfFFhFp",
					"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
					"PmmdzqPrVvPwwTWBwg",
					"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
					"ttgJtRGJQctTZtZT",
					"CrZsJsPPZsGzwwsLwLmpwMDw",
					"CrZsJsPPZsGzwwaLwLmpwMDw",
				},
			},
			want: []byte{'p', 'L', 'P', 'v', 't', 's'},
		},
		{
			name: "With correct format rucksack",
			args: args{
				rucksacks: []string{
					"vJrwpWtwJgWrhcsFMMfFFhFp",
					"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
					"PmmdzqPrVvPwwTWBwg",
					"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
					"ttgJtRGJQctTZtZT",
					"CrZsJsPPZsGzwwsLwLmpwMDw",
					"CrZsJsPPZsGzwwaLwLmpwMDw",
				},
			},
			want: []byte{'p', 'L', 'P', 'v', 't', 's'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorInRucksacks(tt.args.rucksacks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorInRucksacks() = %v, want %v", got, tt.want)
			}
		})
	}
}
