第４回メモ
=========
## 継承が難しすぎる問題
```
Rectangle
  setWidth(int width) {
    this.width = width
  }
  setHeight()
```

```
Square
  setWidth(int width){
    this.width = width
    this.height = width
  }
  setHeight(int height){
    setWidth(height)
  }
```
とすると
```
a = new Rectangle()
a.setWidth(100)
a.setHeight(200)
a.getWidth() * a.getHeight() # 20000
b = new Square)()
b.setWidth(100)
b.setHeight(200)
b.getWidth() * b.getHeight() # 40000
```
振る舞いが異なる。

継承は・・・難しいンゴ・・・

## オブジェクト指向言語
* Java, Ruby, JavaScript 1995
* Design Pattern 1995
  - 継承よりコンポジションをうたっている

## インターフェース
* structual型付け
  - Golang
* nominal型付け
  - Java

## Javaの契約
* 使う側への契約
* サブクラスを作るための契約 (スーパークラスにするのであれば)

## ポリモフィズム (多態性)
* 同じメソッドを呼べることと書いてあることが多い
* 一つのオブジェクトは複数の型を持つ (by JPL)
