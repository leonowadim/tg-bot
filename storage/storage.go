package storage

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"telegramBot/lib/e"
)

type Storage interface { // МЕТОДЫ ИНТЕРФЕЙСА ДЛЯ РАБОТЫ СО СТРАНИЦМИ, ОТПРАВЛЕННЫМИ ПОЛЬЗОВАТЕЛЕМ
	Save(ctx context.Context, p *Page) error
	PickRandom(ctx context.Context, userName string) (*Page, error)
	Remove(ctx context.Context, p *Page) error
	IsExists(ctx context.Context, p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("no saved pages")

type Page struct { // СТРАНИЦА, КОТОРУЮ ОТПРАВЛЯЕМ БОТУ ДЛЯ СОХРАНЕНИЯ
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) { // ХЭШИРОВАНИЕ ПО URL И ИМЕНИ ПОЛЬЗОВАТЕЛЯ
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can not calculate hash", err)
	}
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can not calculate hash", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
