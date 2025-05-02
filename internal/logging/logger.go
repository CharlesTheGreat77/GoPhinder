package logging

import (
	"fmt"
	"strings"
	"time"

	"gophinder/utils"
)

// function to help output the logs just a tad bit cleaner
func LogBlock(title string, fields map[string]string) {
	const ( // prob a cleaner way to do this
		reset = "\033[0m"
		bold  = "\033[1m"
		green = "\033[32m"
		cyan  = "\033[36m"
	)

	// timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	header := fmt.Sprintf("┌──[%s%s%s]%s[%s%s%s]",
		bold, title, reset, strings.Repeat("─", 4), cyan, timestamp, reset)

	// should print any map[string]string being thrown
	visibleHeader := utils.StripANSI(header)
	totalWidth := len(visibleHeader)

	staticPart := fmt.Sprintf("┌──[%s%s%s]", bold, title, reset)
	staticVisible := utils.StripANSI(staticPart)
	timestampPart := fmt.Sprintf("[%s%s%s]", cyan, timestamp, reset)
	timestampVisible := utils.StripANSI(timestampPart)

	dashCount := totalWidth - len(staticVisible) - len(timestampVisible)
	header = fmt.Sprintf("%s%s%s", staticPart, strings.Repeat("─", dashCount), timestampPart)
	footer := "└" + strings.Repeat("─", totalWidth-1)

	fmt.Println()
	fmt.Println(header)
	for key, value := range fields {
		// only concern is this fixed len for clean output (12)
		fmt.Printf("│ %s[+] %-12s%s : %s\n", green, key, reset, value)
	}
	fmt.Println(footer)
}
