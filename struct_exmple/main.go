package main

import (
	"fmt"
)

// 1. typeで型定義
// 先頭大文字でPublicフィールド。外部Packageから参照可能
// 先頭小文字でPrivateフィールド。外部Packageから参照不可能
type T struct {
	PublicField  int
	privateField int
}

// 2. 構造体は値型（いわゆるプリミティブってやつ？？？）
// 関数の引数として呼び出す時にコピーが行われる。
// もとの構造体を変更する場合はポインタをやりとりする。
type T1 struct {
	Field1 int
	Field2 int
}

// コピーした値を変更しているだけなのでもとの構造体は変化しない
func func1(t T1) {
	t.Field1 = 2 * t.Field1
	t.Field2 = 2 * t.Field2
}

// ポインタを使ってもとの構造体を変更する
func func2(t *T1) {
	t.Field1 = 2 * t.Field1
	t.Field2 = 2 * t.Field2
}

// 3. method
// typeで定義した型にmethodの形で紐付ける。クラス構文は存在しない
// フィールドと同じように先頭大文字でPublic、小文字でPrivate
type T3 struct {
	Field1 int
	Field2 int
}

func (v T3) Method1() {
	fmt.Println(v.Field1)
}

func (v T3) method2() {
	fmt.Println(v.Field2)
}

// var i interface{}

// 4. Interface
// 何もメソッドを要求しない空のインターフェースを作成した場合、
// 全ての変数がこのインターフェースを満たしていると言える
func describe(i interface{}) {
	fmt.Println(i)
}

type MyInterface interface {
	MethodA(x, y int) string
}

type MyStruct1 struct {
	Name string
}

type MyStruct2 struct {
	Name string
}

// Inerface MyStruct1
func (s MyStruct1) MethodA(x, y int) string {
	sum := x + y
	return fmt.Sprintf("MyS1 Name: %+v sum: %v\n", s.Name, sum)
}

// Inerface MyStruct2
func (s MyStruct2) MethodA(x, y int) string {
	sum := x + y
	return fmt.Sprintf("MyS2 Name: %+v sum: %v\n", s.Name, sum)
}

func main() {
	t := T{
		PublicField:  0,
		privateField: 0,
	}
	fmt.Printf("%+v\n", t)

	t1 := T1{Field1: 100, Field2: 100}
	fmt.Printf("%+v\n", t1)

	func1(t1)
	fmt.Printf("%+v\n", t1)

	func2(&t1)
	fmt.Printf("%+v\n", t1)

	t3 := T3{Field1: 100, Field2: 100}
	t3.Method1()
	t3.method2()

	x := 100
	describe(x)

	y := "100"
	describe(y)

	myStruct1 := MyStruct1{"hello"}
	myStruct2 := MyStruct2{"world"}

	fmt.Print(myStruct1.MethodA(100, 50))
	fmt.Print(myStruct2.MethodA(100, 50))
}
