package controllers

import (
	"gotrading/app/models"
	"gotrading/bitflyer"
	"gotrading/config"
	"log"
)

func StreamIngestionData() {
	// bitflyer.Tickerを扱うチャンネルを作る
	var tickerChannl = make(chan bitflyer.Ticker)
	// apiを扱うハンドル
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	// goroutineでたくさん??リアルタイムにTickerをとってくる
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	// とってきたTickerを各TickerChannelごとに表示
	// 他のファンクションをブロックしないようにgoroutine化する
	go func() {
		for ticker := range tickerChannl {
			log.Printf("action=StreamIngestionData, %v", ticker)
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated == true && duration == config.Config.TradeDuration {
					// TODO
				}
			}
		}
	}()
}
