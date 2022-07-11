package service

import (
	"fmt"
	"github.com/LuxAeterna-git/my-parcer/pkg/model"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"sourcegraph.com/sourcegraph/go-selenium"
	"strings"
)

type Repository interface {
	Store(product model.Good)
	FindAll() []model.Good
	DeleteByID(id int)
}
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ParseGoods() error {

	// Use selenium to make our request work, usual Client.Get() doesn't work here cuz of defense of this site
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:4444/wd/hub"); err != nil {
		return err
	}
	defer webDriver.Quit()

	err = webDriver.Get("https://www.ozon.ru/category/moloko-9283/")
	if err != nil {
		return err
	}

	page, err := webDriver.PageSource()
	if err != nil {
		return err
	}
	// Got this page, now going to start parsing some data

	doc, err := htmlquery.Parse(strings.NewReader(page))
	if err != nil {
		return err
	}
	list, _ := htmlquery.QueryAll(doc, "//div")
	cards := make([]*html.Node, 0)
	for _, element := range list {
		atr := element.Attr
		if len(atr) != 0 {
			if element.Attr[0].Val == "uj0 j1u" {
				cards = append(cards, element)
			}
		}

	}
	err = s.parseCards(cards)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) parseCards(cards []*html.Node) error {
	var counter int
	for _, element := range cards {
		counter++
		var good model.Good
		good.Name = searchName(element)
		good.Url = searchURL(element)
		good.UrlImg = searchIMG(element)
		good.Price = searchPrice(element)
		s.repo.Store(good)
	}
	fmt.Println("Cards parsed: ", counter)
	return nil
}
func searchName(node *html.Node) string {
	list, _ := htmlquery.QueryAll(node, "//span")
	for _, element := range list {
		atr := element.Attr
		if len(atr) != 0 {
			if element.Attr[0].Val == "d9m m9d dn0 n1d tsBodyL s4j" {
				return element.FirstChild.FirstChild.Data
			}
		}

	}
	return ""
}

func searchURL(node *html.Node) string {
	list, _ := htmlquery.Query(node, "//a@href")

	//for _, element := range list {
	//	res := htmlquery.SelectAttr(element, "href")
	//	return "ozon.ru/" + res
	//}
	res := "www.ozon.ru/" + htmlquery.SelectAttr(list, "href")

	return res
}

func searchIMG(node *html.Node) string {
	img, _ := htmlquery.Query(node, "//img@src")

	res := htmlquery.SelectAttr(img, "src")

	return res
}

func searchPrice(node *html.Node) string {

	list, _ := htmlquery.QueryAll(node, "//span")
	for _, element := range list {
		atr := element.Attr
		if len(atr) != 0 {
			if element.Attr[0].Val == "ui-q5 ui-q9 ui-r1" {
				return element.FirstChild.Data
			}
		}

	}
	return ""
}
