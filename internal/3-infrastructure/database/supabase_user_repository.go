/*
### コードの目的について 🚀

このGoコードは、**Supabase** (オープンソースのFirebase代替サービス) と連携するユーザー管理のための**リポジトリ**を定義しています。主な役割は以下の通りです。

* **`infrastructure` パッケージ**: このパッケージは、`domain`層で定義されたインターフェースの実装を含みます。データベースや外部APIなど、外部サービスとの連携を担当します。
* **`SupabaseUserRepository` 構造体**: この構造体は、ユーザーリポジトリの具体的な実装です。`*supabase.Client`を保持しており、Supabaseプロジェクトと通信するための主要なツールとなります。
* **`newSupabaseUserRepository()` 関数**: これは**コンストラクタ関数**です。新しい`SupabaseUserRepository`を作成し、設定するのが役割です。重要なのは、**環境変数**からSupabaseのURLとサービスキーを取得して`supabase.Client`を初期化している点です。これにより、機密情報がコードベースに直接含まれることを防ぐ、良いプラクティスが採用されています。
* **`register()` メソッド**: このメソッドは、新しいユーザーを登録するための主要なロジックを実装しています。`domain.User`オブジェクト（おそらく`Email`と`Password`フィールドを持つ）を受け取り、`client.Auth.SignUp()`メソッドを使用してSupabaseの認証システムに新しいユーザーアカウントを作成します。

要するに、このコードは**関心の分離**を明確にし、アプリケーションのビジネスロジック（`domain`層で定義）が、基盤となるデータベースや認証システム（この場合はSupabase）の詳細を知ることなく、ユーザーデータとやり取りできるようにしています。
*/

package infrastructure

import (
	"ShittakaKeijiban_Back/internal/domain"
	"context"
	"os"

	"github.com/supabase-community/supabase-go"
)

// SupabaseUserRepository は Supabase と連携するユーザーリポジトリの実装です。
type SupabaseUserRepository struct {
	client *supabase.Client
}

// NewSupabaseUserRepository は SupabaseUserRepository のコンストラクタです。
func NewSupabaseUserRepository() (*SupabaseUserRepository, error) {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_KEY"), nil)
	if err != nil {
		return nil, err
	}
	return &SupabaseUserRepository{client: client}, nil
}

// Register は Supabase で新規ユーザー登録を行います。
func (r *SupabaseUserRepository) Register(user *domain.User) error {
	_, err := r.client.Auth.SignUp(context.Background(), supabase.AuthSignUpParams{
		Email:    user.Email,
		Password: user.Password,
	})
	return err
}
