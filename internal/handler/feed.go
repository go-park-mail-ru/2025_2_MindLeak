package handler

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func FeedHandler(w http.ResponseWriter, r *http.Request, sessions *repository.InMemorySession, articles *repository.InMemoryArticle) {
	session, err := sessions.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	cookies.SetCookie(w, session.SessionId)

	mockArticles, err := articles.GetAllArticles()
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(mockArticles) == 0 {
		authorId := uuid.New()

		_, _ = articles.CreateArticle(authorId,
			"ИИ в 2025: Как нейросети меняют бизнес-процессы",
			"Искусственный интеллект в 2025 году стал неотъемлемой частью бизнеса...")

		_, _ = articles.CreateArticle(authorId,
			"Как российский стартап привлёк $10M на рынке SaaS",
			"Российский стартап CloudPeak разработал SaaS-платформу...")

		_, _ = articles.CreateArticle(authorId,
			"Тренды контент-маркетинга: Что работает в 2025 году",
			"Контент-маркетинг в 2025 году переживает новый виток...")

		_, _ = articles.CreateArticle(authorId,
			"Почему 80% стартапов терпят неудачу в первый год",
			"Запуск стартапа — это всегда риск...")

		_, _ = articles.CreateArticle(authorId,
			"Как мы увеличили конверсию на 30% с помощью UX",
			"Компания BrightPath переработала интерфейс...")

		_, _ = articles.CreateArticle(authorId,
			"Экспериментальный сверхдлинный заголовок статьи, в котором мы попробуем уместить сразу и суть, и интригу, и даже немного юмора, чтобы проверить, как фронтенд справится с рендерингом текста, выходящего далеко за пределы обычной длины заголовков в реальных публикациях, ведь иногда бывают ситуации, когда авторы злоупотребляют символами, а пользователи всё равно должны видеть аккуратно оформленный результат",
			`Это тестовое содержимое статьи, которое специально сделано очень длинным, чтобы проверить работу фронтенда с большими объёмами текста. Представьте, что здесь находится целый лонгрид: мы начинаем с введения, потом плавно переходим к основным идеям, приводим десятки примеров, вставляем длинные цитаты и оформляем структурированные абзацы.  

На протяжении всего текста мы проверяем перенос строк, отступы, работу с большими блоками текста. Например:  

- Пункт первый: фронтенд должен уметь обрезать или корректно отображать длинный текст в превью.  
- Пункт второй: необходимо убедиться, что карточка поста не «разъезжается» при слишком большом контенте.  
- Пункт третий: нужно протестировать скролл или сокращение текста.  

Далее идёт ещё больше текста, чтобы имитировать статью на несколько страниц.  
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus varius, justo sit amet varius bibendum, orci felis interdum dui, non tempus justo eros nec velit. Nulla facilisi. Donec at semper massa, sed bibendum nulla. Mauris eget neque nec nunc facilisis porttitor. Cras id mauris lorem.  

В итоге эта статья должна показать, что и бэкенд корректно отдаёт большие строки, и фронтенд правильно их рендерит, не ломая сетку, не обрезая важную информацию и не портя внешний вид.  
`)

		mockArticles, _ = articles.GetAllArticles()
	}

	if err := json.Write(w, http.StatusOK, mockArticles); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

}
