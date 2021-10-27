package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)
func pikaqiu(blood1 int)(blood int) {
	var skill []int
	skill = make([]int, 3)
	skill[1] = 20
	skill[2] = 30
	m := rand.Intn(2)
	var c int
	c = skill[m]
	blood = blood1 - skill[m]
	if blood <= 0{
		blood = 0
	}
	fmt.Printf("战斗开始：您的可达鸭受到精灵的伤害为%v\n"+"剩余血量:%v\n", c, blood)
	return blood
}
func people(blood1 int) int  {
	var (
		blood,c int
	)
	blood = 50
	fmt.Print("请选择你要可达鸭使用的技能\n" + "1.我最呆：造成对对方总生命值%80的伤害\n"+"2.萌系攻击：造成30伤害\n")
	fmt.Scan(&c)
	if c==1 {
		blood = blood1 - blood1*4/5
	}else if c==2 {
		blood = blood1 - 30
	}
	if blood <= 0 {
		blood = 0
	}
	fmt.Printf("使用技能%v",c)
	fmt.Printf("精灵剩余血量为：%v\n",blood)
	return blood
}

func main() {
	fmt.Println(time.Now().UnixNano())
	var k int
	var name string
	fmt.Print("请输入您的玩家姓名")
	fmt.Scan(&name)
	fmt.Printf("好的，神奇宝贝大师%s，您现在只有一只可达鸭，请选择您的对战对象\n", name)
	fmt.Println("1.皮卡丘\n" + "2.杰尼龟\n" + "3.妙蛙种子")
	fmt.Printf("请注意血量低于总生命值的20%%时即可开始捕捉\n")
	fmt.Scan(&k)
	var (
		a, b int
	)
	switch k {
	case 1:
		fmt.Print("您选择和皮卡丘对战\n")
		a = 100
		b = 50
		for true {
			if a <= 0 {
				a = 0
				fmt.Print("您失败了")
				break
			} else if b <= 0 {
				b = 0
				fmt.Print("您成功了")
				break
			}
			if b <= 10 {
				fmt.Print("请选择是否开始进行捕捉\n","1.捕捉\n2.不捕捉")
				i := 0
				var f int
				fmt.Scan(&i)
				switch i {
				case 1 :
				f =	rand.Intn(100)
					if f >20{
						fmt.Print("恭喜您捕捉成功\n")
						os.Exit(0)
					} else if f <20{
						fmt.Print("很遗憾捕捉失败，战斗继续\n")
					}
				case 2:
					fmt.Print("战斗继续\n")
				}
			}
			a = pikaqiu(a)
			b = people(b)
		}

	case 2:
		fmt.Print("您选择和杰尼龟对战\n")
		a = 100
		b = 50
		for true {
			if a <= 0 {
				a = 0
				fmt.Print("您已死亡")
				break
			} else if b <= 0 {
				b = 0
				fmt.Print("精灵死亡")
				break
			}
			if b <= 10 {
				fmt.Print("请选择是否开始进行捕捉\n","1.捕捉\n2.不捕捉")
				i := 0
				var f int
				fmt.Scan(&i)
				switch i {
				case 1 :
					f =	rand.Intn(100)
					if f >20{
						fmt.Print("恭喜您捕捉成功\n")
						os.Exit(0)
					} else if f <20{
						fmt.Print("很遗憾捕捉失败，战斗继续\n")
					}
				case 2:
					fmt.Print("战斗继续\n")
				}
			}
			a = pikaqiu(a)
			b = people(b)
		}
	case 3:
		fmt.Print("您选择和妙蛙种子对战\n")
		a = 100
		b = 50
		for true {
			if a <= 0 {
				a = 0
				fmt.Print("您失败了")
				break
			} else if b <= 0 {
				b = 0
				fmt.Print("您成功了")
				break
			}
			if b <= 10 {
				fmt.Print("请选择是否开始进行捕捉\n","1.捕捉\n2.不捕捉")
				i := 0
				var f int
				fmt.Scan(&i)
				switch i {
				case 1 :
					f =	rand.Intn(100)
					if f >20{
						fmt.Print("恭喜您捕捉成功\n")
						os.Exit(0)
					} else if f <20{
						fmt.Print("很遗憾捕捉失败，战斗继续\n")
					}
				case 2:
					fmt.Print("战斗继续\n")
				}
			}
			a = pikaqiu(a)
			b = people(b)
		}
	}
}