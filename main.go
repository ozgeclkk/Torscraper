package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func main() {

	inputFile := "targets.yaml"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("[-] Hata: targets.yaml bulunamadı.")
		return
	}
	defer file.Close()

	logFile, _ := os.Create("scan_report.log")
	defer logFile.Close()

	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		fmt.Println("[-] Hata: Tor servisine bağlanılamadı.")
		return
	}

	client := &http.Client{
		Transport: &http.Transport{Dial: dialer.Dial},
		Timeout:   time.Second * 45,
	}

	os.MkdirAll("outputs", os.ModePerm)
	scanner := bufio.NewScanner(file)

	fmt.Println("[+] Tarama başladı. Loglar scan_report.log dosyasına yazılıyor...")

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}

		status := "SUCCESS"
		resp, err := client.Get(url)

		if err != nil {
			status = "TIMEOUT/FAILED"
			fmt.Printf("[ERR] %s -> %s\n", url, status)
		} else {

			body, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()

			safeName := strings.ReplaceAll(url, "http://", "")
			replacer := strings.NewReplacer(
				"/", "_",
				"?", "_",
				"=", "_",
				"&", "_",
				":", "_",
				"*", "_",
				"\"", "_",
				"<", "_",
				">", "_",
				"|", "_",
			)
			safeName = replacer.Replace(safeName)
			fileName := "outputs/" + safeName + ".html"

			ioutil.WriteFile(fileName, body, 0644)
			fmt.Printf("[INFO] %s -> %s\n", url, status)
		}

		logEntry := fmt.Sprintf("[%s] URL: %s - Status: %s\n", time.Now().Format("15:04:05"), url, status)
		logFile.WriteString(logEntry)
	}

	fmt.Println("[+] İşlem tamamlandı. 'outputs' klasörünü ve 'scan_report.log' dosyasını kontrol edin.")
}
