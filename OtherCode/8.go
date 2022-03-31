/* 三门问题
参赛者的面前有三扇关闭着的门，其中一扇的后面是天使，选中后天使会达成你的一个愿望，而另
外两扇门后面则是恶魔，选中就会死亡。
当你选定了一扇门，但未去开启它的时候，上帝会开启剩下两扇门中的一扇，露出其中一只恶魔。（上
帝是全能的，必会打开恶魔门）随后上帝会问你要不要更换选择，选另一扇仍然关着的门。
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	changeDoor, unchangeDoor := 0, 0
	door := []int{0, 1, 2}
	for i := 0; i < 10000; i++ {
		angleDoor := rand.Intn(3)
		selectDoor := rand.Intn(3)
		for j := 0; j < len(door); j++ {
			if door[j] != angleDoor && door[j] != selectDoor {
				door = append(door[0:j], door[j+1:]...)
			}
		}
		if angleDoor == selectDoor {
			unchangeDoor++
		} else {
			changeDoor++

		}

	}
	fmt.Println("unchangeDoor:", float64(unchangeDoor)/10000.0)
	fmt.Println("changeDoor:", float64(changeDoor)/10000.0)

}
