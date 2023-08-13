package dataloader

// import (
// 	"reflect"
// 	"testing"
// )

// func Test_test(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			test()
// 		})
// 	}
// }

// func Test_outFormat(t *testing.T) {
// 	type args struct {
// 		in string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{"SP", args{"SP"}, "SP"},
// 		{"td", args{"td"}, "td"},
// 		{"Space", args{"{{space}}"}, " "},
// 		{"Eq", args{"{{eq}}"}, "="},
// 		{"Gt", args{"{{gt}}"}, ">"},
// 		{"Null", args{"null"}, ""},
// 		{"{{Null}}", args{"{{null}}"}, ""},
// 		{"amp", args{"{{amp}}"}, "&"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := outFormat(tt.args.in); got != tt.want {
// 				t.Errorf("outFormat() = [%v], want [%v]", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_setSearch(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{"Empty/Empty", args{"", ""}, ""},
// 		{"Empty/Value", args{"", "cat"}, "cat"},
// 		{"Value/Empty", args{"prop", ""}, "prop"},
// 		{"Value/Value", args{"prop", "cat"}, "prop$cat"},
// 		{"Value/Value", args{"PrOp", "cAt"}, "PrOp$cAt"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := setSearch(tt.args.property, tt.args.category); got != tt.want {
// 				t.Errorf("setSearch() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_get(t *testing.T) {
// 	type args struct {
// 		in   string
// 		kind string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{"Empty-Empty", args{"", ""}, ""},
// 		{"Empty-Value", args{"", "cat"}, ""},
// 		{"Value-Empty", args{"prop", ""}, "prop"},
// 		{"ok", args{"ok", ""}, "OK"},
// 		{"ok-alt", args{"ok", "ALT"}, "OK alt"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := get(tt.args.in, tt.args.kind); got != tt.want {
// 				t.Errorf("get() = %v, want %v", got, tt.want)
// 			}
// 			t.Logf("get(%v,%v) = %v, want %v", tt.args.in, tt.args.kind, get(tt.args.in, tt.args.kind), tt.want)
// 		})
// 	}
// }

// func Test_getInt(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    int
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"Empty/Empty", args{"", ""}, 0, true},
// 		{"Empty/Value", args{"", "cat"}, 0, true},
// 		{"Value/Empty", args{"prop", ""}, 0, true},
// 		{"Value/Value", args{"int", ""}, 123, false},
// 		{"Value/Value", args{"int", "alt"}, 0, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := getInt(tt.args.property, tt.args.category)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getInt() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("getInt() = %v, want %v", got, tt.want)
// 			}
// 			t.Logf("get(%v,%v) = %v, want %v", tt.args.property, tt.args.category, get(tt.args.property, tt.args.category), tt.want)

// 		})
// 	}
// }

// func Test_getBool(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    bool
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"Empty/Empty", args{"", ""}, false, true},
// 		{"Empty/Value", args{"", "cat"}, false, true},
// 		{"Value/Empty", args{"prop", ""}, false, true},
// 		{"Value/Value", args{"bool", ""}, true, false},
// 		{"Value/Value", args{"bool", "alt"}, false, false},
// 		{"Value/Value", args{"fail", ""}, false, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := getBool(tt.args.property, tt.args.category)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getBool() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("getBool() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// // The function `Test_getFloat` is a test function that tests the `getFloat` function.
// func Test_getFloat(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    float64
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"Empty/Empty", args{"", ""}, 0, true},
// 		{"Empty/Value", args{"", "cat"}, 0, true},
// 		{"Value/Empty", args{"prop", ""}, 0, true},
// 		{"Value/Value", args{"float", ""}, 123.456, false},
// 		{"Value/Value", args{"float", "alt"}, 0.0, false},
// 		{"Value/Value", args{"fail", ""}, 0, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := getFloat(tt.args.property, tt.args.category)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getFloat() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("getFloat() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_getList(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"EmptyEmpty", args{"", ""}, nil, true},
// 		{"EmptyValue", args{"", "cat"}, nil, true},
// 		{"ValueEmpty", args{"prop", ""}, nil, true},
// 		{"ValueValue", args{"list", ""}, []string{"a", "b", "c"}, false},
// 		{"ValueValueAlt", args{"list", "alt"}, []string{"d", "e", "f"}, false},
// 		{"ValueEmptyFail", args{"fail", ""}, nil, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := getList(tt.args.property, tt.args.category)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getList() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("getList() = %v, want %v", got, tt.want)
// 			}
// 			t.Logf("getList([%v],[%v]) = [%v], want [%v] err [%e]", tt.args.property, tt.args.category, got, tt.want, err)
// 		})
// 	}
// }

// func Test_getListInt(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []int
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"EmptyEmpty", args{"", ""}, nil, true},
// 		{"EmptyValue", args{"", "cat"}, nil, true},
// 		{"ValueEmpty", args{"prop", ""}, nil, true},
// 		{"ValueValue", args{"listint", ""}, []int{1, 2, 3}, false},
// 		{"ValueValueAlt", args{"listint", "alt"}, []int{4, 5, 6}, false},
// 		{"ValueEmptyFail", args{"fail", ""}, nil, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := getListInt(tt.args.property, tt.args.category)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getListInt() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("getListInt() = %v, want %v", got, tt.want)
// 			}
// 			t.Logf("getList([%v],[%v]) = [%v], want [%v] err [%e]", tt.args.property, tt.args.category, got, tt.want, err)
// 		})
// 	}
// }

// func Test_getListFloat(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []float64
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"EmptyEmpty", args{"", ""}, nil, true},
// 		{"EmptyValue", args{"", "cat"}, nil, true},
// 		{"ValueEmpty", args{"prop", ""}, nil, true},
// 		{"ValueValue", args{"listfloat", ""}, []float64{1.1, 2.2, 3.3}, false},
// 		{"ValueValueAlt", args{"listfloat", "alt"}, []float64{4.4, 5.5, 6.6}, false},
// 		{"ValueEmptyFail", args{"fail", ""}, nil, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := getListFloat(tt.args.property, tt.args.category)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getListFloat() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("getListFloat() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_getMap(t *testing.T) {
// 	type args struct {
// 		property string
// 		category string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    map[string]string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"EmptyEmpty", args{"", ""}, nil, true},
// 		{"EmptyValue", args{"", "cat"}, nil, true},
// 		{"ValueEmpty", args{"prop", ""}, nil, true},
// 		{"ValueValue", args{"map", ""}, map[string]string{"a": "1", "b": "2", "c": "3"}, false},
// 		{"ValueValueAlt", args{"map", "alt"}, map[string]string{"d": "4", "e": "5", "f": "6"}, false},
// 		{"ValueEmptyFail", args{"fail", ""}, nil, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := getMap(tt.args.property, tt.args.category)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getMap() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("getMap() = %v, want %v", got, tt.want)
// 			}
// 			t.Logf("getMap([%v],[%v]) = [%v], want [%v] err [%e]", tt.args.property, tt.args.category, got, tt.want, err)
// 		})
// 	}
// }

// func Test_setVerbose(t *testing.T) {
// 	type args struct {
// 		v bool
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 		{"True", args{true}},
// 		{"False", args{false}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			setVerbose(tt.args.v)
// 		})
// 	}
// }
