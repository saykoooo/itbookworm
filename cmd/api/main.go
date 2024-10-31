package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

// Объявим строковую константу, которая содержит номер версии приложения.
// Позже мы будем генерировать версию автоматически во время сборки, а пока
// просто сохраним жёстко заданную глобальную константу.
const version = "1.0.0"

// Определим структуру конфигурации, которая будет содержать все параметры
// конфигурации для нашего приложения. На данный момент параметрами будут порт,
// который должен прослушивать сервер, и название производственной среды 
// (разработка, производство и т.д.). Эти параметры будут считываться из флагов командной строки.
type config struct {
	port	int
	env		string
}

// Определим структуру приложения, которая будет содержать зависимости для
// обработчиков HTTP, вспомогательных функций и middleware. На данный момент
// она содержит копию структуры конфигурации и логгер. По мере развития проекта
// она будет включать в себя гораздо больше.
type application struct {
	config	config
	logger	*log.Logger
}

func main() {
	// Объявляем экземпляр структуры config.
	var cfg config

	// Записываем значения флагов командной строки port и env в структуру
  // конфигурации. По умолчанию мы используем номер порта 8000 и среду development.
	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Инициализируем новый логгер, который записывает сообщения в стандартный вывод
  // с указанием текущей даты и времени.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Объявляем экземпляр структуры приложения, которая содержит структуру
  // конфигурации и логгер.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Объявляем новый мультиплексор и добавляем маршрут `/v1/healthcheck`,
	// который будет перенаправлять запросы в метод `healhcheckHandler`
	// (мы создадим его чуть позже).
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthCheck", app.healhcheckHandler)

	// Объявляем HTTP-сервер с настройками тайм-аута, который прослушивает порт,
	// указанный в структуре конфигурации, и использует созданный выше мультиплексор.
	srv := &http.Server{
		Addr: 				fmt.Sprintf(":%d", cfg.port),
		Handler: 			mux,
		IdleTimeout: 	time.Minute,
		ReadTimeout: 	10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Запускаем HTTP-сервер
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
