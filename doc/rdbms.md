### 企業テーブル (companies)
| カラム名            | データ型     | 制約     | 説明     |
| ------------------- | ------------ | -------- | -------- |
| company_id          | varchar(255) | PK       | 企業ID   |
| name                | varchar(255) | NOT NULL | 法人名   |
| representative_name | varchar(255) | NOT NULL | 代表者名 |
| phone_number        | varchar(20)  | NOT NULL | 電話番号 |
| postal_code         | varchar(10)  | NOT NULL | 郵便番号 |
| address             | varchar(255) | NOT NULL | 住所     |
| created_at          | datetime     |          | 作成日時 |
| updated_at          | datetime     |          | 更新日時 |
| deleted_at          | datetime     |          | 削除日時 |

### ユーザーテーブル (users)
| カラム名   | データ型     | 制約     | 説明                         |
| ---------- | ------------ | -------- | ---------------------------- |
| user_id    | varchar(255) | PK       | ユーザーID                   |
| company_id | varchar(255) | FK       | 企業ID                       |
| name       | varchar(255) | NOT NULL | 氏名                         |
| email      | varchar(255) | NOT NULL | メールアドレス               |
| password   | varchar(255) | NOT NULL | パスワード（ハッシュ化済み） |
| created_at | datetime     |          | 作成日時                     |
| updated_at | datetime     |          | 更新日時                     |
| deleted_at | datetime     |          | 削除日時                     |

### 取引先テーブル (clients)
| カラム名   | データ型     | 制約 | 説明                       |
| ---------- | ------------ | ---- | -------------------------- |
| company_id | varchar(255) | FK   | クライアントを持つ企業ID   |
| client_id  | varchar(255) | FK   | クライアント先となる企業ID |

### 取引先銀行口座テーブル (company_banks)
| カラム名        | データ型     | 制約     | 説明                     |
| --------------- | ------------ | -------- | ------------------------ |
| company_bank_id | varchar(255) | PK       | 口座ID                   |
| company_id      | varchar(255) | FK       | 企業ID                   |
| bank_code       | varchar(255) | NOT NULL | 銀行コード               |
| branch_code     | varchar(255) | NOT NULL | 支店コード               |
| bank_name       | varchar(255) | NOT NULL | 銀行名                   |
| branch_name     | varchar(255) | NOT NULL | 支店名                   |
| account_type    | varchar(255) | NOT NULL | 口座の種類（普通or当座） |
| account_number  | varchar(20)  | NOT NULL | 口座番号                 |
| account_holder  | varchar(255) | NOT NULL | 口座名義人               |
| created_at      | datetime     |          | 作成日時                 |
| updated_at      | datetime     |          | 更新日時                 |
| deleted_at      | datetime     |          | 削除日時                 |

### 請求書データテーブル (invoices)
| カラム名       | データ型       | 制約     | 説明                                            |
| -------------- | -------------- | -------- | ----------------------------------------------- |
| invoice_id     | varchar(255)   | PK       | 請求書ID                                        |
| company_id     | varchar(255)   | FK       | 企業ID                                          |
| client_id      | varchar(255)   | FK       | 取引先ID                                        |
| payment_amount | decimal(10, 2) | NOT NULL | 支払金額                                        |
| fee            | decimal(10, 2) | NOT NULL | 手数料                                          |
| fee_rate       | decimal(5, 2)  | NOT NULL | 手数料率                                        |
| tax            | decimal(10, 2) | NOT NULL | 消費税                                          |
| tax_rate       | decimal(5, 2)  | NOT NULL | 消費税率                                        |
| total_amount   | decimal(10, 2) | NOT NULL | 請求金額                                        |
| status         | varchar(20)    | NOT NULL | ステータス (未処理、処理中、支払い済み、エラー) |
| due_at         | datetime       | NOT NULL | 支払予定日時                                    |
| created_at     | datetime       |          | 作成日時(発行日時)                              |
| updated_at     | datetime       |          | 更新日時                                        |
| deleted_at     | datetime       |          | 削除日時                                        |