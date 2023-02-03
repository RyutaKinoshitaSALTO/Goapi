### 実行
```
% make serve
//取得
% curl -i localhost:8080/todos/
//追加
% curl -i -X POST -H "Content-Type: application/json" -d '{"title":"add", "content":"追加"}' localhost:8080/todos/
//変更
% % curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"delete", "content":"変更"}' localhost:8080/todos/4
//削除
% curl -i -X DELETE localhost:8080/todos/4
```

###テスト実行
```
% cd controller
% go test
```
