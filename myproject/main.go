package main

import (
	"encoding/json"
	"fmt"
	"github.com/tommitoan/bazica"
	"time"
)

func main() {
	// 시간대 설정
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	
	// 날짜 설정 (1985년 5월 15일 오전 10시)
	now := time.Date(1985, time.Month(5), 15, 10, 0, 0, 0, loc)
	gender := 1 // 성별: 0 = 여성, 1 = 남성

	// 입력된 날짜, 시간 출력 (디버깅용)
	fmt.Println("입력된 날짜와 시간:", now)
	fmt.Println("지역:", loc)
	fmt.Println("성별:", gender)

	// 사주 계산
	chart, err := bazica.GetBaziChart(now, loc, gender)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 결과를 JSON으로 변환
	jsonData, _ := json.Marshal(chart)
	fmt.Println("사주 결과:", string(jsonData))
}
