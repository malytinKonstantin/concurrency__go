package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/malytinKonstantin/concurrency__go/internal/compute/parser"
	"github.com/malytinKonstantin/concurrency__go/internal/storage/engine"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		logger.Info("Получен сигнал завершения")
		cancel()
	}()

	cmdParser := parser.NewParser()
	eng := engine.NewEngine()

	runCLI(ctx, cmdParser, eng, logger)
}

func runCLI(ctx context.Context, cmdParser parser.Parser, eng engine.Engine, logger *zap.Logger) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			logger.Info("Завершение работы...")
			return
		default:
			handleCommand(scanner.Text(), cmdParser, eng, logger)
			fmt.Print("> ")
		}
	}

	if err := scanner.Err(); err != nil {
		logger.Error("Ошибка сканера", zap.Error(err))
	}
}

func handleCommand(input string, cmdParser parser.Parser, eng engine.Engine, logger *zap.Logger) {
	if strings.TrimSpace(input) == "" {
		return
	}

	cmd, err := cmdParser.Parse(input)
	if err != nil {
		logger.Error("Ошибка парсинга команды", zap.Error(err))
		return
	}

	switch cmd.Type {
	case parser.SET:
		eng.Set(cmd.Args[0], cmd.Args[1])
		logger.Info("Значение установлено", zap.String("key", cmd.Args[0]))
	case parser.GET:
		if value, exists := eng.Get(cmd.Args[0]); exists {
			fmt.Printf("Value: %s\n", value)
			logger.Info("Значение получено", zap.String("key", cmd.Args[0]))
		} else {
			logger.Info("Ключ не найден", zap.String("key", cmd.Args[0]))
		}
	case parser.DEL:
		if deleted := eng.Del(cmd.Args[0]); deleted {
			logger.Info("Ключ удалён", zap.String("key", cmd.Args[0]))
		} else {
			logger.Info("Ключ не найден для удаления", zap.String("key", cmd.Args[0]))
		}
	default:
		logger.Error("Неизвестная команда", zap.String("command", string(cmd.Type)))
	}
}
