# AtCoder for Go

## プロジェクトの作り方

```bash
mkdir [project]
cd [project]
go mod init [project]
go run main.go
go build -o [project].exe
```

## ソース管理

### 初期解答時

- プロジェクトディレクトリを作成
- 初期化
- `main.go`を書く
- 正解ならコミット・プッシュ

## レビュー時

- レビュー用のブランチを作成
- Copilotにレビューを依頼
- 結果をソースコードに取込
- コミット・プッシュ
- プルリクエストを作成
- リクエストをマージ（コメントにレビュー結果）
- レビュー用のブランチを削除

```bash
git checkout -b [branch_name]
git push -u origin [branch_name]
```
