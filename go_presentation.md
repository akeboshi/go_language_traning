<!-- $theme: gaia -->

# Ruby ではなくGoLangを使う理由
# RJ Isao Aruga

---
# 型 
* Ruby
  - 動的型付け
  - 型は書かない！Rubyの信念(matz談)
    - https://twitter.com/yukihiro_matz/status/773871448435720192

* Go
  - 静的型付け言語
  - 型推論はある

---
# 並行処理
* ruby
  * thread size 2mb
* Go
  * goroutine
    * stack size 10KB 軽量！ 
  * CSP(communicating sequential processes)
  * channel

---
# コーディングスタイル
* Ruby
  - rubocopが一般的
    - 設定する項目が多すぎる
    - 設定値のデファクトがない？
    - 初期値の設定が厳しすぎてハゲる
  - RubyMineがやってくれたりもする

* Go
  - go fomatter さいつよ。
    - go format に従うのが標準的
    - どのプロジェクト見ても同じなので見やすい。すごいつよい。

---
# 標準ライブラリ (HTTP Server)

---
# 標準ライブラリ (Parser)
* Ruby (JSON, XML, YAML 標準)
  簡単！
  ```
  require 'json'
  JSON.parse('{"foo":"bar"}')
  ```
* Go (JSON 標準)
  ```
  
  ```

---
# Frame Work
* Ruby
  - Ruby on Rails

* Go
  - Echo, Gin , etc...

---
# Test
test
coverage
example


---
# 実行環境
* Ruby
  - 環境ごとにバージョンにあわせたRubyをインストールしなければいけない

* Go
  - Cross compiler
  - シングルバイナリで動作する

---
# 開発環境 (Editor)
* Ruby
  - Ruby Mine (有料: 基本はこれかなぁ)
  - Vim (あれこれ設定すると便利。最近使ってないのでわからんが、補完機能が微妙だった気がする。)
    - vimに慣れてない(Emacs派のことを言っているわけではない。Emacsでも開発できるのかもしれないけど知らない。)人からの敷居が高い
* Go
  - Atom
  - Vim (何回も書くが、Emacsは知らん！！)
  - IntelliJ?
  - Eclipse? (Eclipse使うならAtomでいい気がする。)
---
# 開発環境
* Ruby
  - rbenv で複数バージョン管理しながら使う

* Go
  https://github.com/moovweb/gvm#features
gvm pkgset create --local とすると、実行したディレクトリを gvm の package set にできます。
その後は、そのディレクトリ以下（サブディレクトリでもよい）で gvm pkgset use --local とすれば gvm がそのディレクトリを $GOPATH に設定してくれます。
違うバージョンを使う場合は gvm use go1.3.3 && gvm pkgset create --local として最初に local package set を作成すれば、次回からは gvm use go1.3.3 && gvm pkgset use --local とすれば切り替えられます。
---
# 実行速度

---
# Reflection
---
# Shall we Go?
---
---
---
