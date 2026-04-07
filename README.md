# AtCoder for Go

AtCoderの過去問題をGo言語で解答するためのリポジトリです。

## プロジェクトの作り方

- 解答の作成は`workspace`ディレクトリで行います。
- `main.go`に解答を作成します。
    - すでにソースが記入済みの場合はクリアしてください。
- テストがOKなら提出します。
- `mv.py`スクリプトを使って提出済みのソースをコピーします。
- 提出済みのソースを修正する場合は`cp.py`スクリプトを使って`main.go`にコピーします。

### mv.py

```bash
python mv.py [category] [sequence] [level]
```

例）`python mv.py ABC 142 A` → `main.go`を`/ABC/101-200/141-150/ABC-142-A.go`にコピー

### cp.py

```bash
python cp.py [category] [sequence] [level]
```

例）`python mv.py ABC 142 A` → `/ABC/101-200/141-150/ABC-142-A.go`を`main.go`にコピー

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

ブランチを作成して移動
変更分をリモートブランチにプッシュ

```bash
git checkout -b [branch_name]
git push -u origin [branch_name]
```

Web側でプルリクエストを作成してマージ
ローカルブランチを削除

```bash
git checkout main
git branch -d [branch_name]
```
