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
│   └── main.go                      # アプリのエントリーポイント
└── internal/                        # 同じモジュール内でしか使えないという意味
    ├── domain/                      # ドメイン層（エンティティ、IF）
    │   ├── question.go              # 質問構造体・ビジネスルール
    │   ├── question_repository.go   # 質問リポジトリのinterface
    │   ├── user.go                  # ユーザー構造体・ビジネスルール
    |   └── user_repository.go       # 「ユーザー情報をどこから／どう取得・保存するか」はここで抽象
    ├── usecase/                     # ユースケース層（ビジネスロジック）
    │   ├── question_interactor.go   # 質問ユースケース
    │   └── user_interactor.go       # ユーザーユースケース
    ├── infrastructure/              # インフラ層（DBなど外部実装）
    │   ├── supabase_question_repository.go # 質問DB実装
    │   └── supabase_user_repository.go     # ユーザーDB実装
    └── interface/                   # プレゼンテーション層 外部I/O層（Gin, HTTPなど）
        ├── handler/                 # HTTPルーティングとリクエスト処理
        │   ├── question_handler.go  # 質問APIハンドラ
        │   └── auth_handler.go      # 認証APIハンドラ
        └── middleware/
            └── auth_middleware.go   #
