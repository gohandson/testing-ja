# テストを書こう

## ハンズオンのやりかた

`skeleton`ディレクトリ以下に問題があり、8個のステップに分けられています。
STEP01からSTEP08までステップごとに進めていくことで、Goのパッケージ分けやテストの方法が学べます。

各ステップに、READMEが用意されていますので、まずは`README`を読みます。
`README`には、そのステップを理解するための解説が書かれています。

`README`を読んだら、ソースコードを開き`TODO`コメントが書かれている箇所をコメントに従って修正して行きます。
`TODO`コメントをすべて修正し終わったら、`README`に書かれた実行例に従ってプログラムをコンパイルして実行します。

途中でわからなくなった場合は、`solution`ディレクトリ以下に解答例を用意していますので、そちらをご覧ください。

`macOS`の動作結果をもとに解説しています。
`Windows`の方は、パスの区切り文字やコマンド等を適宜読み替えてください。

## 目次

* [STEP01: Goに触れる](./skeleton/step01)
* [STEP02: 標準パッケージを使ってみよう](./skeleton/step02)
* [STEP03: パッケージを分けよう](./skeleton/step03)
* [STEP04: 外部パッケージを使ってみよう](./skeleton/step04)
* [STEP05: テストを書いてみよう](./skeleton/step05)
* [STEP06: テストヘルパーを作って見よう](./skeleton/step06)
* [STEP07: テストのパッケージを分けよう](./skeleton/step07)
* [STEP08: テーブル駆動テストを行おう](./skeleton/step08)

## ソースコードの取得

```
$ go env GOPATH
$ cd ↑のディレクトリに移動
$ mkdir -p src/github.com/gohandson/
$ cd src/github.com/gohandson
$ git clone https://github.com/gohandson/testing-ja
$ cd testing-ja
```

## ライセンス

<a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja">
	<img width="200" src="by-nc.eu.png">
</a>
