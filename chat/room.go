package main

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func (r *room) run() {
	// goroutineでバックグラウンドで実行される無限ループ
	// アプリケーション内の他の処理をブロックすることはない

	// join,leave,forwardのチャネルを監視して
	// いずれかにメッセージが届くとcase節が実行される
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
