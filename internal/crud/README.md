# CRUDの比較
## Ent
- CRUD実行時にuser.SetName、user.SetAgeを使って明確にユーザーの値を入れることができる
## Gorm
- CURD実行時にテーブル作成の際の型がそのまま使える
- CreateやFind関数の引数がinterface{}で実装されているため型の安全性が低い
## Beego
## bun
