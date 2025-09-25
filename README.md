沼ったので別リポジトリでやり直し



# ShittakaKeijiban_Back
- 知ったか掲示板バックエンド

## フォルダ構成
/your_project_root
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   └── main.go                  # アプリのエントリーポイント
└── internal/                    # 同じモジュール内でしか使えないという意味
    ├── domain/                  # ドメイン層（エンティティ、IF）
    │   ├── question.go          # 構造体・ビジネスルール
    │   └── question_repository.go # リポジトリのinterface
    ├── usecase/                 # ユースケース層（ビジネスロジック）
    │   └── question_interactor.go
    ├── infrastructure/         # インフラ層（DBなど外部実装）
    │   └── supabase_question_repository.go
    └── interface/              # プレゼンテーション層 外部I/O層（Gin, HTTPなど）
        ├── handler/            # HTTPルーティングとリクエスト処理
        │   └── question_handler.go
        └── middleware/
            └── auth_middleware.go
