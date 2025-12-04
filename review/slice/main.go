package main



func main(){



	// 슬라이스 리터럴
	notesliteral := []string{"one","two","three","four","five","six","seven"}
	for i := range notesliteral{
		print(notesliteral[i])
	}
	print("\n")



	// 슬라이스
	// 슬라이스는 배열과 다르게 크기가 고정되어 있지 않다.
	// 동적으로 크기가 변할 수 있다.
	// 슬라이스는 배열의 일부를 '참조'하는 자료구조이다.

	// var notes []string // 배열과 달리 크기를 지정하지 않음!
	// notes = make([]string, 7)
	//일곱 개의 문자열을 담을 수 있는 슬라이스 생성
	notes := make([]string, 7)
	notes[0] = "zero"
	notes[1] = "one"
	notes[2] = "two"
	print(notes[0])
	print(notes[1])
	print(notes[2])
	print(notes)
	
} 