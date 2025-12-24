package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/proxy" 
)

func main() {
	file, err := os.Open("targets.yaml")
	if err != nil {
		log.Fatalf("targets.yaml bulunamadı!")
	}
	defer file.Close()

	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Tor bağlantısı kurulamadı!")
	}

	transport := &http.Transport{Dial: dialer.Dial}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 120,
	}

	os.MkdirAll("outputs", 0755)
	logFile, _ := os.Create("scan_report.log")
	defer logFile.Close()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer("socks5://127.0.0.1:9150"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	scanner := bufio.NewScanner(file)
	fmt.Println("TARAMA BAŞLATILDI")

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" || strings.HasPrefix(url, "#") {
			continue
		}

		fmt.Printf("[INFO] Scanning: %s -> ", url)

		_, err := client.Get(url)
		if err != nil {
			fmt.Println("TIMEOUT")
			logFile.WriteString(fmt.Sprintf("[ERR] Scanning: %s -> TIMEOUT\n", url))
			continue
		}

		var buf []byte
		taskCtx, _ := chromedp.NewContext(allocCtx)
		taskCtx, cancelTask := context.WithTimeout(taskCtx, 100*time.Second)
		
		err = chromedp.Run(taskCtx,
			chromedp.Navigate(url),
			chromedp.Sleep(5*time.Second),
			chromedp.FullScreenshot(&buf, 100),
		)
		
		if err == nil && len(buf) > 0 {
			fileName := strings.NewReplacer("http://", "", "https://", "", "/", "_", ".", "_").Replace(url)
			_ = os.WriteFile("outputs/"+fileName+".png", buf, 0644)
			fmt.Println("SUCCESS (Captured)")
			logFile.WriteString(fmt.Sprintf("[INFO] Scanning: %s -> SUCCESS\n", url))
		} else {
			fmt.Println("SUCCESS (SS Failed)")
			logFile.WriteString(fmt.Sprintf("[INFO] Scanning: %s -> SUCCESS (But SS Error)\n", url))
		}
		cancelTask()
	}
	fmt.Println("İŞLEM TAMAM")
}