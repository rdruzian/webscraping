package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	playwright "github.com/playwright-community/playwright-go"
)

const enem = "https://www.gov.br/inep/pt-br/areas-de-atuacao/avaliacao-e-exames-educacionais/enem/provas-e-gabaritos"

// const fcmscsp = "https://vestibular.brasilescola.uol.com.br/downloads/faculdade-ciencias-medicas-santa-casa-sao-paulo.htm"
const downloadTest = "https://download.inep.gov.br/enem/provas_e_gabaritos"

func main() {
	resp, err := http.Get(enem)
	if err != nil {
		fmt.Printf("error to search page %d", err)
	}
	defer resp.Body.Close()

	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	enemPath := "C:\\Users\\renat\\Downloads\\enem"
	err = os.Mkdir(enemPath, 0777)
	assertErrorToNilf("could not create directory: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto(enem)
	assertErrorToNilf("could not goto: %w", err)
	assertErrorToNilf("could not set content: %w", page.SetContent(`<a href data-id="2023">2023</a>`))

	test := fmt.Sprintf(`<a class="external-link" href="%s/%s_PV_impresso_D1_CD1.pdf" target="_blank" title="" data-tippreview-enabled="false" data-tippreview-image="" data-tippreview-title="">Prova</a>`, downloadTest, strconv.Itoa(2023))
	fmt.Sprintf("Download link: %s", test)
	assertErrorToNilf("could not set content: %w", page.SetContent(test))

	var download playwright.Download
	err = download.SaveAs(enemPath)
	assertErrorToNilf("could change directory: %w", err)

	download, err = page.ExpectDownload(func() error {
		return page.Locator(`text=Prova`).Click()
	})
	assertErrorToNilf("could not download: %w", err)

	pathDownload, err := download.Path()
	assertErrorToNilf("could not get directory: %w", err)

	fmt.Println("Directory before set a specific diretory", pathDownload)
	fmt.Println("URL: ", download.URL())
	assertErrorToNilf("could not get directory: %w", err)

	assertErrorToNilf("could not copy file: %w", err)
	assertErrorToNilf("could not close browser: %w", browser.Close())
	assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
