// STEP03: パッケージを分けてみよう

// mainパッケージの定義
package main

// greetingパッケージのインポート
import "github.com/gohandson/testing-ja/greeting"

// main関数から実行される
func main() {
	// greeting.Do関数を呼び出す
	greeting.Do()
}
