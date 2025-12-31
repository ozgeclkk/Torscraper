Tor Onion Scraper (CTI Tool)

Bu proje, Siber Tehdit İstihbaratı (CTI) süreçlerinde veri toplama (Collection) aşamasını otomatize etmek için geliştirilmiş bir araçtır. Go (Golang) dili kullanılarak hazırlanan bu uygulama, toplu bir hedef listesindeki .onion uzantılı adresleri Tor ağı üzerinden anonim olarak tarar ve içeriklerini yerel olarak saklar.
🚀 Temel Özellikler

    Anonimlik: Tüm trafik SOCKS5 (Tor) üzerinden yönlendirilir, gerçek IP adresi gizlenir.

    Otomasyon: Yüzlerce linki tek tek gezmek yerine tek komutla toplu tarama yapar.

    Hata Toleransı: Erişilemeyen (dead) siteler programı durdurmaz; hata loglanır ve sıradaki URL'e geçilir.

    Veri Seti Oluşturma: Başarılı taramalar outputs/ klasörüne HTML formatında kaydedilir.

🛠️ Kurulum ve Gereksinimler

Programın çalışması için bilgisayarınızda Tor Browser'ın açık olması veya arka planda bir Tor Servisi'nin çalışıyor olması gerekmektedir.

    Bağımlılıkları Yükleyin:
    Bash

    go mod tidy

    Hedef Listesini Düzenleyin: targets.yaml dosyasına taramak istediğiniz .onion adreslerini her satıra bir tane gelecek şekilde ekleyin.

💻 Kullanım

Uygulamayı çalıştırmak için terminale şu komutu yazın:
Bash

go run main.go

Eğer uygulamayı derlenmiş bir dosya (binary) olarak kullanmak isterseniz:
Bash

go build -o cti_scraper.exe main.go
./cti_scraper.exe

📂 Çıktılar

    outputs/: Tarama sonucunda elde edilen sitelerin HTML kaynak kodları.

    scan_report.log: Taramaların durum raporu (Success/Timeout).
