# ShittakaKeijiban_Back
- 知ったか掲示板バックエンド





## フォルダ再構成

/project-root
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── README.md
├── cmd/                                  # アプリケーションのエントリーポイント（mainパッケージ）
│   └── main.go                           # サーバー起動、DI（依存関係注入）、ルータ登録
└── internal/                             # 内部パッケージ（外部公開しない）
    ├── domain/                           # ドメイン層（ビジネスルールの中心）
    │   ├── entity/                       # エンティティ：ビジネスモデルと基本ルール
    │   │   ├── user.go                   # ユーザー情報（ニックネーム、認証関連）
    │   │   ├── question.go               # 問題情報（タイトル、本文、ジャンルなど）
    │   │   ├── choice.go                 # 選択肢情報（テキスト、正解フラグ）
    │   │   ├── answer.go                 # 回答履歴（誰がどの選択肢を選んだか）
    │   │   ├── genre.go                  # ジャンル情報
    │   │   ├── comment.go                # コメント情報（問題や回答へのコメント）
    │   │   └── tag.go                    # タグ情報（タグ名）
    │   └── repository/                   # 永続化の抽象化（DBアクセスのIF）
    │       ├── user_repository.go        # ユーザー用リポジトリIF（認証、プロフィール）
    │       ├── question_repository.go    # 問題用リポジトリIF
    │       ├── choice_repository.go      # 選択肢用リポジトリIF
    │       ├── answer_repository.go      # 回答用リポジトリIF
    │       ├── comment_repository.go     # コメント用リポジトリIF
    │       └── tag_repository.go         # タグ用リポジトリIF
    ├── usecase/                           # ユースケース層（アプリケーションロジック）
    │   ├── user_usecase.go               # ユーザー関連の操作（登録、マイページ）
    │   ├── question_usecase.go           # 問題投稿、閲覧、編集、削除
    │   ├── ranking_usecase.go            # 人気問題ランキング取得
    │   ├── comment_usecase.go            # コメント投稿、取得
    │   └── search_usecase.go             # 検索（タグ、タイトル、ジャンル）
    ├── infrastructure/                   # インフラ層（外部サービスやDB）
    │   ├── database/                     # DB接続やリポジトリ実装
    │   │   ├── supabase_client.go        # Supabaseクライアント設定（APIキー、URL）
    │   │   ├── user_repository.go        # ユーザーリポジトリの実装（Supabase Auth）
    │   │   ├── question_repository.go    # 問題リポジトリの実装
    │   │   ├── comment_repository.go     # コメントリポジトリの実装
    │   │   └── tag_repository.go         # タグリポジトリの実装
    │   └── config/                       # 設定管理
    │       └── config.go                 # 環境変数読み込み、設定情報の取得
    └── interface/                        # プレゼンテーション層（外部I/O）
        ├── handler/                      # HTTPハンドラ（Ginコントローラ）
        │   ├── auth_handler.go           # 認証API（/auth/register, /auth/login）
        │   ├── user_handler.go           # ユーザーAPI（マイページ、プロフィール）
        │   ├── question_handler.go       # 問題API（投稿、一覧、詳細）
        │   ├── ranking_handler.go        # ランキングAPI
        │   ├── comment_handler.go        # コメントAPI
        │   └── search_handler.go         # 検索API
        ├── middleware/                   # HTTPミドルウェア
        │   ├── auth_middleware.go        # JWT認証ミドルウェア
        │   ├── cors_middleware.go        # CORS対応
        │   └── logger_middleware.go      # アクセスログ出力
        ├── dto/                          # データ転送オブジェクト（APIリクエスト/レスポンス）
        │   ├── user_dto.go               # ユーザー用DTO
        │   ├── question_dto.go           # 問題用DTO
        │   └── comment_dto.go            # コメント用DTO
        └── router/                       # ルーティング設定
            └── router.go                 # エンドポイントとハンドラの紐付け
