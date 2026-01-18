package main

import (
	"fmt"

	. "sentiment-analyzer/types"
	. "sentiment-analyzer/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func AppInstance() *fiber.App {
	app := fiber.New()
	return app
}

var app = AppInstance()
var result AnalysisResponse

func main() {
	// Rota analyze-feed
	// Analisa um feed de mensagens e retorna métricas de sentimento
	app.Post("/analyze-feed", func(c *fiber.Ctx) error {
		var msg MessageRequest

		// Parse e validação da requisição
		if err := c.BodyParser(&msg); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error ao parsear o corpo da requisição",
				"code":  fmt.Sprint(fiber.StatusBadRequest),
			})
		}

		if err := validate.Struct(&msg); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "O corpo da requisição possui dados inválidos ou incorretos",
				"code":  fmt.Sprint(fiber.StatusBadRequest),
			})
		}

		// Inicia o processamento da mensagem com SentimentAnalyzer
		// if len(msg.Messages) < 100 {
		// 	_, result = SentimentAnalyzer(msg)
		// }

		// Para MUITOS dados (> 1000 mensagens) - usar pool com mais workers
		// if len(msg.Messages) > 1000 {
		// Usar 2x o número de CPUs para I/O intensivo
		_, result = SentimentAnalyzer(msg)
		// }

		for _, msgTime := range msg.Messages {
			if IsTimeWindow(msgTime.Timestamp, msg.TimeWindowMinutes) {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": "Valor de janela temporal não suportado na versão atual",
					"code":  "UNSUPPORTED_TIME_WINDOW",
				})
			}
		}

		// if WindowTimeCalculator(msg.Messages) {
		// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		// 		"error": "Não foi possível processar as mensagens fornecidas",
		// 		"code":  "UNSUPPORTED_TIME_WINDOW",
		// 	})
		// }

		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Listen(":8000")
}
