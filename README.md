# shorclink

[![GitHub License](https://img.shields.io/github/license/ayane0857/shorclink)](https://opensource.org/licenses/MIT)
[![GitHub top language](https://img.shields.io/github/languages/top/ayane0857/shorclink)](https://go.dev/)

`shorclink`とは省略リンクプロジェクトです。go で開発されています。

## 環境変数の設定

`shorclink`の動作は`.env`ファイルによって制御されます。

### ルートパスのリダイレクト先の設定

ルートパス (`/`) にアクセスした際のリダイレクト先 URL を設定します。

```.env
ORIGINAL_LINK = https://example.com/
```

---

### API 関連の設定

**API のモード設定**

API の動作モードを指定します。

- `invalid`: API を無効化します
- `required`: API の利用に API キーを必須にします
- `free`: API キーなしで利用を許可します

```.env
API = required
```

**API トークン**

`API = required` の場合に必要となるトークンを設定します。

```.env
API_TOKEN = your_api_token
```

---

### ショートコードの設定

生成されるショートコードに関する設定です。

**ショートコードのモード設定**

ショートコードの動作モードを指定します。

- `required`: ショートコードを指定しなければなりません。
- `auto`: ショートコードが指定されていない場合は自動生成します。
- `generate`: ショートコードを自動生成します。

```.env
SHORT_CODE = auto
```

**ショートコードの長さ**

`API = required`以外の場合に必要となるショートコードの長さを設定します。

```.env
SHORT_CODE_LENGTH = 6
```

---

### データベース設定 (PostgreSQL)

接続する PostgreSQL データベースの情報を設定します。

```.env
POSTGRES_USER=your_db_user
POSTGRES_PASSWORD=your_db_password
POSTGRES_DB=your_db_name
POSTGRES_HOST=localhost
```
