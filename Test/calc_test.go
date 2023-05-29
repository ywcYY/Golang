package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}

}

type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}

}

func TestMul(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
}

// calc_test.go
/*
func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(3, 3) != 6 {
			t.Fatal("fail")
			t.Log()
			t.Helper()

		}
		//之前的例子测试失败时使用 t.Error/t.Errorf，这个例子中使用 t.Fatal/t.Fatalf，区别在于前者遇错不停，还会继续执行其他的测试用例，后者遇错即停。
	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})

}*/
func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func TestConn(t *testing.T) {
	setup()
	defer teardown()
	request := httptest.NewRequest("GET", "http://baidu.com", nil)
	recorder := httptest.NewRecorder()
	helloHandler(recorder, request)
	all, _ := io.ReadAll(recorder.Result().Body)
	if string(all) != "hello world" {
		t.Fatal("expected hello world, but got\", string(bytes)")
	} else {
		fmt.Println("Connection success")
	}

}
