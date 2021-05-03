package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adlio/trello"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const PERSONAL_BOARD_TODAY = "PERSONAL_BOARD_TODAY"
const PERSONAL_BOARD_TOMORROW = "PERSONAL_BOARD_TOMORROW"
const PERSONAL_BOARD_THIS_WEEK = "PERSONAL_BOARD_THIS_WEEK"
const PERSONAL_BOARD_WAITING = "PERSONAL_BOARD_WAITING"
const PERSONAL_BOARD_INBOX = "PERSONAL_BOARD_INBOX"
const PERSONAL_BOARD_DONE = "PERSONAL_BOARD_DONE"

const WORK_BOARD_TODAY = "WORK_BOARD_TODAY"
const WORK_BOARD_TOMORROW = "WORK_BOARD_TOMORROW"
const WORK_BOARD_THIS_WEEK = "WORK_BOARD_THIS_WEEK"
const WORK_BOARD_WAITING = "WORK_BOARD_WAITING"
const WORK_BOARD_INBOX = "WORK_BOARD_INBOX"
const WORK_BOARD_DONE = "WORK_BOARD_DONE"

var trelloClient *trello.Client

type FetchListLength func() float64

func MakeFetchListFunc(listEnvName string) FetchListLength {
	listId := os.Getenv(listEnvName)
	return func() float64 {
		list, err := trelloClient.GetList(listId, trello.Defaults())
		if err != nil {
			log.Fatalln("Error in retrieving list with id:", listId, "err:", err)
		}
		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			log.Fatalln("Error in retrieving cards from list.", "err:", err)
		}
		return float64(len(cards))
	}
}

func main() {

	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	token := os.Getenv("TOKEN")

	trelloClient = trello.NewClient(apiKey, token)
	retryClient := retryablehttp.NewClient()
	trelloClient.Client = retryClient.StandardClient()

	configureExportedTrelloMetrics()

	log.Println("Starting Trello productivity monitoring exporter...")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func configureExportedTrelloMetrics() {
	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "personal_board_today_count",
		Help: "The count of cards in the today list of the Productivity System (Personal) board",
	}, MakeFetchListFunc(PERSONAL_BOARD_TODAY))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "personal_board_tomorrow_count",
		Help: "The count of cards in the tomorrow list of the Productivity System (Personal) board",
	}, MakeFetchListFunc(PERSONAL_BOARD_TOMORROW))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "personal_board_this_week_count",
		Help: "The count of cards in the this week list of the Productivity System (Personal) board",
	}, MakeFetchListFunc(PERSONAL_BOARD_THIS_WEEK))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "personal_board_waiting_count",
		Help: "The count of cards in the waiting list of the Productivity System (Personal) board",
	}, MakeFetchListFunc(PERSONAL_BOARD_WAITING))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "personal_board_inbox_count",
		Help: "The count of cards in the inbox list of the Productivity System (Personal) board",
	}, MakeFetchListFunc(PERSONAL_BOARD_INBOX))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "personal_board_done_count",
		Help: "The count of cards in the done list of the Productivity System (Personal) board",
	}, MakeFetchListFunc(PERSONAL_BOARD_DONE))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "work_board_today_count",
		Help: "The count of cards in the today list of the Productivity System (Work) board",
	}, MakeFetchListFunc(WORK_BOARD_TODAY))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "work_board_tomorrow_count",
		Help: "The count of cards in the tomorrow list of the Productivity System (Work) board",
	}, MakeFetchListFunc(WORK_BOARD_TOMORROW))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "work_board_this_week_count",
		Help: "The count of cards in the this week list of the Productivity System (Work) board",
	}, MakeFetchListFunc(WORK_BOARD_THIS_WEEK))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "work_board_waiting_count",
		Help: "The count of cards in the waiting list of the Productivity System (Work) board",
	}, MakeFetchListFunc(WORK_BOARD_WAITING))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "work_board_inbox_count",
		Help: "The count of cards in the inbox list of the Productivity System (Work) board",
	}, MakeFetchListFunc(WORK_BOARD_INBOX))

	var _ = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "work_board_done_count",
		Help: "The count of cards in the done list of the Productivity System (Work) board",
	}, MakeFetchListFunc(WORK_BOARD_DONE))

}
