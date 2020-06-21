# Go Language + Docker + Echo プラクティス

Go言語についてググる際に `Go` を検索するよりも `Golang` で検索した方がGo言語に近い情報が表示されます。そのため、以下 Go言語 を Golangに略します。
これから連絡することです。

- [x] staticファイルの設置
- [x] ビューファイルおよびビューへパラメーターの渡し方
- [x] パッケージのprivateとpublicメソッド
- [ ] MySQLと接続
- [ ] ORMを実装
- [ ] 自動コンパイルを実装
- [ ] .gitignoreの定義
- [ ] go.modとgo.sumの役割

### Dockerで Golang を始める

```bash
docker-compose up

# コンパイル＆goアプリケーションを起動
docker-compose exec app go run server.go

# ブラウザを起動
open http://localhost:8080
```

### Golang について理解したこと

* 定義した変数を使わないと unused declared "variable" とコンパイルエラーが出ます。
* GOPATHとGOROOTを正しく設定しなくてもコンパイルができます。
* フォルダ名を正しく定義しないとコンパイルエラー(Cannot find packageエラー)が起きます。ただし、ファイル名はコンパイルと関係ありません。
* 関数の最初文字が大文字か小文字でpublicとprivate関数で定義します。

### Echo について理解したこと

* Golangを完全理解した前提のドキュメント
* HTMLテンプレートはGolangデフォルのhtml/templateを利用しています。HTMLテンプレートの説明や使い方はgolang公式サイトのdocを参考した方が良いです。
* 最新バージョンのGolangに追い付いていないことです。最新のGolangに追い付く必要もなさそうです。


