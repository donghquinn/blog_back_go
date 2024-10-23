package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/donghquinn/blog_back_go/libraries/network"
	"github.com/donghquinn/blog_back_go/utils"
)

func main() {
	network.SetConfigs()
	// 첫 번째 로그 파일 생성
	utils.RotateLogFile()

	// 로그 회전을 스케줄링
	go utils.ScheduleLogRotation()

	network.DatabaseConnect()

	server := network.OpenServer()

	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	log.Printf("[DEBUG] App Host %s", configs.GlobalConfig.AppHost)
	log.Printf("[START] Server Listening On: %s", configs.GlobalConfig.AppPort)
	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

	// 종료 신호를 받을 채널 생성
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 서버를 고루틴으로 실행
	go func() {
		log.Println("Server Start")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Start Error: %v", err)
		}
	}()

	// 종료 신호 대기
	<-quit
	log.Println("Received Shut Down Signal")

	// 셧다운 컨텍스트 설정 (예: 5초 제한)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 서버 그레이스풀 셧다운
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed Graceful Shutdown: %v", err)
	}

	log.Println("Server Has been Shutdown Gracefully")
}
