package parser

import (
	"errors"
	"fmt"
	"io"
	"os"

	//"log"
	"net/http"
	"time"

	"regexp"
	"strings"
)

func (p *Parser) Parse(login, password, path, fileName string, mothly bool) (string, error) {

	data := strings.NewReader(fmt.Sprintf("%s%s%s%s", "type=logon&action=execute&redirect_params=&login=", login, "&password=", password))
	req, err := http.NewRequest("POST", p.Client.BasePath+"/?type=logon", data)
	if err != nil {

		return "", err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "http://external.roszdravnadzor.ru")
	req.Header.Set("Referer", "http://external.roszdravnadzor.ru/?type=logon")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")

	resp, err := p.Client.Do(req)
	if err != nil || resp == nil {
		//	log.Fatal(err)

		return "", err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		//	log.Fatal(err)

		return "", err
	}

	coockie := resp.Header.Get("Set-Cookie")

	reg := regexp.MustCompile(`sid_EXTERNAL=(\d+)`)
	if len(reg.FindStringSubmatch(coockie)) == 0 {

		return "", errors.New("wrong login or password")
	}

	sidValue := reg.FindStringSubmatch(coockie)
	if len(sidValue) < 2 {
		return "", errors.New("empty sid value")
	}
	count, err := p.reqWithFilter(sidValue[1], path, fileName, mothly)
	if err != nil {
		return "", err
	}

	return count, nil
}

func (p *Parser) reqWithFilter(sid, path, fileName string, monthly bool) (string, error) {
	now := time.Now()
	var startTime time.Time = now.AddDate(0, 0, -7)
	var publishedFrom, publishedTo string

	if monthly {
		startTime = now.AddDate(0, -1, 0)
	}
	year, month, day := startTime.Date()
	publishedFrom = fmt.Sprintf("%d.%d.%d", day, month, year)

	cyear, cmonth, cday := now.Date()
	publishedTo = fmt.Sprintf("%d.%d.%d", cday, cmonth, cyear)

	req, err := http.NewRequest("GET", p.Client.BasePath, nil)

	query := req.URL.Query()

	query.Add("q_dt_published_from", publishedFrom)
	query.Add("q_dt_published_to", publishedTo)
	query.Add("q_kind", "0")
	query.Add("q_pregnancy", "0")
	req.URL.RawQuery = query.Encode()

	if err != nil {
		//	log.Fatal(err)
		fmt.Println("1")
		return "0", err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", fmt.Sprintf("%s%s", "sid_EXTERNAL=", sid))
	req.Header.Set("Referer", "http://external.roszdravnadzor.ru/?type=logon")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	resp, err := p.Client.Do(req)
	if err != nil {
		fmt.Println("2")
		//	log.Fatal(err)
		return "0", err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("3")
		//	log.Fatal(err)
		return "0", err
	}

	regex := regexp.MustCompile(`\d+\s-\s\d+\sиз\s(\d+)`)
	matches := regex.FindStringSubmatch(string(bodyText))
	if len(matches) > 0 {
		if err := p.downloadXLS(sid, path, fileName, publishedTo, publishedFrom); err != nil {
			return matches[1], nil
		}
		return matches[1], nil
	} else {
		return "0", nil
	}

}

func (p *Parser) downloadXLS(sid, path, fileName, publishTo, publishFrom string) error {

	req, err := http.NewRequest("GET", p.Client.BasePath+"/?", nil)

	query := req.URL.Query()

	query.Add("q_dt_published_from", publishFrom)
	query.Add("q_dt_published_to", publishTo)
	query.Add("type", "phv")
	query.Add("xls", "t1")
	query.Add("portion", "15")
	req.URL.RawQuery = query.Encode()

	if err != nil {
		//	log.Fatal(err)
		fmt.Println("1")
		return err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", fmt.Sprintf("%s%s", "sid_EXTERNAL=", sid))
	req.Header.Set("Referer", "http://external.roszdravnadzor.ru/?type=logon")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	resp, err := p.Client.Do(req)
	if err != nil {
		return err
		// 	}
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(fmt.Sprintf("%s/%s", path, fileName))
	if err != nil {
		return nil
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {

		return nil
	}

	//fmt.Println(resp)
	if err != nil {
		fmt.Println("2")
		//	log.Fatal(err)
		return err
	}

	defer resp.Body.Close()
	//bodyText, err := io.ReadAll(resp.Body)

	//return string(bodyText), nil
	fmt.Println("Download done")
	return nil
}
