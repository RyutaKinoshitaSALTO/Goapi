### 実行
```
% make serve
//取得
% curl -i localhost:8080/todos/
//追加
% curl -i -X POST -H "Content-Type: application/json" -d '{"title":"add", "content":"add"}' localhost:8080/todos/
//変更
% % curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"delete", "content":"change"}' localhost:8080/todos/4
//削除
% curl -i -X DELETE localhost:8080/todos/4
```

### テスト実行
```
% cd controller
% go test
```

今回実行の際にmakeコマンドで実行しており、下記URLサイトを参考にセットアップを行った。
https://zenn.dev/genki86web/articles/6e61c167fbe926
