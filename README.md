# go-redmine-discord-notificator

**完全に個人用に作っているため、雑に書いています**

- 送信元からはSMTPサーバーとして振る舞い、DiscordのDMへメールの題名と本文を送るプロキシ的なやつ
- メールのToのlocalpartをDisplayNameと考え、MemberListに書かれたDisplayNameと対応するDiscordIDに対してDMを送る
- **想定としてはRedmineの個人宛通知メールをDiscordに送るためのもの**

## 事前準備
- MemberListにメアドの@前(≒DisplayName),DiscordのIDの順番で入れたCSVを作る
  例:
  ```csv
  DisplayName,DiscordID
  hogeTaro,1234567890
  ```

## 動かし方

Linux or macOSを想定しています。

```
git clone https://github.com/mikuta0407/go-redmine-discord-notificator.git
cd go-redmine-discord-notificator
go build
```

```console
DISCORD_TOKEN="ReplaceToYourDiscordBotToken" ./go-redmine-discord-notificator
```

- `-l 0.0.0.0:25`とかすれば無防備に25番で口を開ける。
- 何もしなければ`127.0.0.1:1025`で開く

## 参考(というかここの改変)

- [emersion/go-smtp: 📤 An SMTP client & server library written in Go](https://github.com/emersion/go-smtp)
  - [go-smtp/cmd/smtp-debug-server/main.go](https://github.com/emersion/go-smtp/blob/master/cmd/smtp-debug-server/main.go)

## LICENSE

MIT
