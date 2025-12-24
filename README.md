# Tor CTI Web Scraper (Golang)

Tor ağı üzerindeki gizli servisleri (**.onion**) anonim biçimde tarayarak **görsel kanıt (screenshot)** toplayan ve süreç boyunca **detaylı log** üreten CTI (Cyber Threat Intelligence) otomasyon aracıdır.  
Bu proje, **Siber Vatan** programı kapsamında CTI süreçlerini pratikte otomatize etmek amacıyla geliştirilmiştir.


## Özellikler

- **Anonimlik (SOCKS5 Zorunlu Yönlendirme)**  
  Tüm trafik yerel SOCKS5 proxy üzerinden (**127.0.0.1:9150**) zorunlu olarak yönlendirilir. Bu sayede IP sızıntısı riski minimize edilir.

- **Otomatik Görsel Kanıt (Full Page Screenshot)**  
  Erişilebilen her onion hedefi için **tam sayfa ekran görüntüsü** otomatik alınır.

- **Hata Toleransı**  
  Kapanmış/erişilemeyen (dead) onion sitelerinde takılmadan taramaya devam eder.

- **Detaylı Loglama**  
  Tüm süreç **scan_report.log** dosyasına anlık olarak yazılır.


##  Gereksinimler

- Bilgisayarınızda **Tor Browser** yüklü olmalı ve **arka planda çalışıyor** olmalı.
- Tor’un yerel SOCKS5 proxy portu olan **9150** dinleniyor olmalı.
---

