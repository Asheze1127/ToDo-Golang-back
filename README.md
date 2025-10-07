# 🧭 Go バックエンド API 設計書

## 🧑‍💻 ユーザー & 認証系

| メソッド | パス | 概要 | 認証 | 備考 |
|----------|------|------|------|------|
| `POST`   | `/signup`          | ユーザー登録 | ❌ | 新規ユーザーを作成する |
| `POST`   | `/login`           | ログインしてJWTトークンを発行 | ❌ | アクセストークンとリフレッシュトークンを返す |
| `GET`    | `/me`              | 自分のプロフィール取得 | ✅ | SSRや初期化用 |
| `PUT`    | `/me`              | 自分のプロフィール更新 | ✅ | パスワード・ユーザー名変更など |
| `GET`    | `/users/{userId}`  | 特定ユーザーのプロフィール取得 | 🔓 or ✅ | 公開プロフィール表示用 |

---

## 📝 タスク系

| メソッド | パス | 概要 | 認証 | 備考 |
|----------|------|------|------|------|
| `GET`    | `/tasks`                  | 自分のタスク一覧を取得 | ✅ | TL・マイページ用 |
| `POST`   | `/tasks`                  | タスクを新規作成 | ✅ |  |
| `GET`    | `/tasks/{taskId}`         | 特定のタスクを取得 | ✅ | 編集・詳細表示用 |
| `PUT`    | `/tasks/{taskId}`         | 特定のタスクを更新 | ✅ |  |
| `DELETE` | `/tasks/{taskId}`         | 特定のタスクを削除 | ✅ |  |
| `GET`    | `/users/{userId}/tasks`   | 特定ユーザーのタスク一覧を取得 | 🔓 or ✅ | プロフィールページのタイムライン用 |

---

## 📊 統計系

| メソッド | パス | 概要 | 認証 | 備考 |
|----------|------|------|------|------|
| `GET`    | `/statistics`             | 自分の統計を取得 | ✅ | 完了率・件数など |
| `GET`    | `/statistics/{userId}`    | 特定ユーザーの統計を取得 | 🔓 or ✅ | プロフィール表示・管理画面用 |

---

## 🌿 ルーティング構成（chi）

```go
r.Post("/signup", signUpHandler)
r.Post("/login", loginHandler)

// 認証が必要なルート
r.Group(func(r chi.Router) {
    r.Use(AuthMiddleware)

    // 自分のプロフィール
    r.Get("/me", meHandler)
    r.Put("/me", updateMeHandler)

    // 自分のタスク
    r.Route("/tasks", func(r chi.Router) {
        r.Get("/", getMyTasksHandler)
        r.Post("/", createTaskHandler)
        r.Get("/{taskID}", getTaskHandler)
        r.Put("/{taskID}", updateTaskHandler)
        r.Delete("/{taskID}", deleteTaskHandler)
    })

    // 自分の統計
    r.Get("/statistics", getMyStatisticsHandler)
})

// 公開プロフィール
r.Route("/users", func(r chi.Router) {
    r.Get("/{userID}", getUserProfileHandler)
    r.Get("/{userID}/tasks", getUserTasksHandler)
    r.Get("/{userID}/statistics", getUserStatisticsHandler)
})

## 🔐 JWTミドルウェアの使い方

```go
import "myapp-backend/internal/middleware"

// ...
r.With(middleware.JWTAuth).Get("/me", meHandler)
```

`Authorization: Bearer <token>` ヘッダー付きのリクエストだけが通過し、ミドルウェア内部で検証済みの `userID` をコンテキストに詰めています。ハンドラー側では `middleware.UserIDFromContext(r.Context())` で取り出せます。

"""
myapp-backend/
├── cmd/
│   └── server/
│       └── main.go         ← サーバーのエントリーポイント
├── internal/
│   ├── handlers/           ← 各エンドポイントの処理
│   │   ├── auth.go
│   │   ├── tasks.go
│   │   ├── statistics.go
│   │   └── users.go
│   ├── models/             ← DBの構造体定義（GORM）
│   │   ├── user.go
│   │   ├── task.go
│   │   └── statistic.go
│   ├── database/           ← DB接続周り
│   │   └── db.go
│   └── middleware/         ← JWT認証など
│       └── auth.go
├── go.mod
└── go.sum
"""

## 学習メモ

Handler（プレゼン層）        ← 外からのリクエストを受ける
       ↓
Service / Usecase層        ← ビジネスロジック（何をするか）
       ↓
Infrastructure層           ← DBやAPIなど外部と通信する
