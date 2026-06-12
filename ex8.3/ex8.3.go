package main

import "fmt"

const Pig int = 0 //그냥 가독성 좋고 디버깅 잘할려고 이렇게씀. 예시인듯?
const Cow int = 1
const Chicken = 2
const ( //그래서 위에거 대신 이렇게 쓰기도함, iota는 1씩증가
	Red   int = iota //1
	Blue  int = iota //2
	Green int = iota //3
)
const ( //이렇게 하기도함
	C1 int = iota + 1 //1 = 0 + 1
	C2                //2 = 1 + 1
	C3                //3 = 2 + 1
)
const (
	BitFlag1 uint = 1 << iota //1 = 1 << 1
	BitFlag2                  // 2 = 1 << 1
	BitFlag3                  // 4 = 1 << 2
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}

}
func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}
