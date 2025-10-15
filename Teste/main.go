package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://download.inep.gov.br/enem/provas_e_gabaritos"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	//entries, err := page.Locator(".athing").All()
	//if err != nil {
	//	log.Fatalf("could not get entries: %v", err)
	//}
	selectYear := fmt.Sprintf(`<a href data-id=%s>%s</a>`, strconv.Itoa(2023), strconv.Itoa(2023))
	assertErrorToNilf("could not set content: %w", page.SetContent(selectYear))
	//for i, entry := range entries {
	title, err := page.Locator("a > class > href").TextContent()
	if err != nil {
		log.Fatalf("could not get text content: %v", err)
	}
	fmt.Printf("%s\n", title)
	//}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
