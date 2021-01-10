package main

import (
	"fmt"
	"math/rand"
	"time"
)

//　　構造体を定義する
type Player struct {
	name  string
	point byte
	hand  byte
}

//  後ろに書いた要素の数を自動でカウントしてくれている
func show_hand(hand byte) string {
	hn := [...]string{"グー", "チョキ", "パー"}
	return hn[hand-1] //謎
}

func judge(p1 *Player, p2 *Player) {
	h1, h2 := p1.hand, p2.hand
	fmt.Printf("%s　対　%s で　", show_hand(h2), show_hand(h1))
	var win *Player //ポインタから、インスタンスを取り出す「操作」　// ポインタはメモリアドレス
	if h1 == h2 {
		fmt.Println("引き分けです。")
		return
	} else if (h2-h1)%3 == 2 { //%は余り
		win = p2
	} else {
		win = p1
	}
	fmt.Printf("%sの勝ちです。\n", win.name) //\n"改行無し
	win.point++                        //変数の値を１増やす

}

func game(p1 *Player, p2 *Player, n int) {
	fmt.Printf("*** %d回戦 ***\n", n)

	p1.hand = 0
	p2.hand = 0

	for p1.hand < 1 || 3 < p1.hand {
		fmt.Printf("%sの手を入力して下さい（1:グー, 2:チョキ, 3:パー）：", p1.name)
		fmt.Scanf("%d", &p1.hand) //整数、基数10
	}
	p2.hand = byte(rand.Intn(3) + 1) //Intn 満たすものをランダムに選んで返す

	judge(p1, p2)
}

func final_winner(p1 *Player, p2 *Player) {
	po1, po2 := p1.point, p2.point
	fmt.Println("\n*** 最終結果 ***")
	fmt.Printf("%d 対 %d で　", po2, po1)
	var win *Player
	if po1 == po2 {
		fmt.Println("引き分けです。")
		return
	} else if po1 > po2 {
		win = p1
	} else {
		win = p2
	}
	fmt.Printf("%sの勝ちです。\n", win.name)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	man := &Player{point: 0}
	cpu := &Player{name: "Computer", point: 0}

	fmt.Print("あなたの名前を入力して下さい：")
	fmt.Scanf("%s", &man.name) //Scanは標準入力から読み込んだテキストをスキャンしてスペースで区切られた値として、順に引数に格納
	fmt.Printf("%s　対　%s　：じゃんけん開始\n\n", cpu.name, man.name)

	for i := 1; i <= 5; i++ {
		game(man, cpu, i)
	}

	final_winner(man, cpu)

}
