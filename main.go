package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	
	playwright "github.com/playwright-community/playwright-go"
)

const downloadTestLink = "https://download.inep.gov.br/enem/provas_e_gabaritos"

// const downloadTestUntil2020 = "https://download.inep.gov.br/educacao_basica/enem/provas"
// const downloadAnswerKeyUntil2020 = "https://download.inep.gov.br/educacao_basica/enem/gabaritos"
const enemPath = ".\\enem"

func main() {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	
	browser, err := pw.Chromium.Launch()
	assertErrorToNilf("could not launch Chromium: %w", err)
	
	page, err := browser.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	
	checkDirectory(enemPath)
	
	for i := 2020; i < 2025; i++ {
		selectYear := fmt.Sprintf(`<a href data-id=%s>%s</a>`, i, i)
		linkTestDayOne := fmt.Sprintf("%s/%s_PV_impresso_D1_CD1.pdf", downloadTestLink, strconv.Itoa(i))
		linkTestDayTwo := fmt.Sprintf("%s/%s_PV_impresso_D2_CD6.pdf", downloadTestLink, strconv.Itoa(i))
		linkAnswerDayOne := fmt.Sprintf("%s/%s_GB_impresso_D1_CD1.pdf", downloadTestLink, strconv.Itoa(i))
		linkAnswerDayTwo := fmt.Sprintf("%s/%s_GB_impresso_D2_CD6.pdf", downloadTestLink, strconv.Itoa(i))
		
		downloadTest(i, linkTestDayOne, enemPath, selectYear, page)
		downloadTest(i, linkTestDayTwo, enemPath, selectYear, page)
		downloadTestAnswerKey(i, linkAnswerDayOne, enemPath, selectYear, page)
		downloadTestAnswerKey(i, linkAnswerDayTwo, enemPath, selectYear, page)
	}
	assertErrorToNilf("could not close browser: %w", browser.Close())
	assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}

func downloadTest(year int, downloadTest, path, selectYear string, page playwright.Page) {
	assertErrorToNilf("could not set content: %w", page.SetContent(selectYear))
	fmt.Println("link:", downloadTest)
	test := fmt.Sprintf(`<a class="external-link" href=%s target="_blank" title="" data-tippreview-enabled="false" data-tippreview-image="" data-tippreview-title="">Prova</a>`, downloadTest)
	assertErrorToNilf("could not set content: %w", page.SetContent(test))
	download, err := page.ExpectDownload(func() error {
		return page.Locator(`text=Prova`).Click()
	})
	assertErrorToNilf("could not download: %w", err)
	// Salvar o arquivo no diretório especificado
	filename := fmt.Sprintf("enem_%s_dia_um.pdf", strconv.Itoa(year))
	err = download.SaveAs(fmt.Sprintf("%s/%s", path, filename))
	assertErrorToNilf("could not save file: %w", err)
}

func downloadTestAnswerKey(year int, downloadTest, path, selectYear string, page playwright.Page) {
	assertErrorToNilf("could not set content: %w", page.SetContent(selectYear))
	fmt.Println("link:", downloadTest)
	test := fmt.Sprintf(`<a class="external-link" href=%s target="_blank" title="" data-tippreview-enabled="false" data-tippreview-image="" data-tippreview-title="">Gabarito</a>`, downloadTest)
	assertErrorToNilf("could not set content: %w", page.SetContent(test))
	download, err := page.ExpectDownload(func() error {
		return page.Locator(`text=Gabarito`).Click()
	})
	assertErrorToNilf("could not download: %w", err)
	// Salvar o arquivo no diretório especificado
	filename := fmt.Sprintf("enem_%s_gabarito_dia_um.pdf", strconv.Itoa(year))
	err = download.SaveAs(fmt.Sprintf("%s/%s", path, filename))
	assertErrorToNilf("could not save file: %w", err)
}

func checkDirectory(path string) {
	if _, err := os.Open(path); os.IsNotExist(err) {
		fmt.Println("The directory named", path, "does not exist")
		fmt.Println("Creating directory", path)
		
		err = os.Mkdir(path, 0777)
		assertErrorToNilf("could not create directory: %w", err)
	} else {
		fmt.Println("The directory namend", path, "exists")
	}
}

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}
