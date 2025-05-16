package route

import (
	"Todo/models"
	"github.com/Aurivena/answer"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary      Создать задачу
// @Description  Создание новой задачи для текущей сессии
// @Tags         Задачи
// @Accept       json
// @Produce      json
// @Param        // @Param X-Session-ID header string true "Идентификатор сессии (указывается в cookie, но описан как header для Swagger)" string true "Идентификатор сессии пользователя"
// @Param        input body models.TodoInput true "Данные задачи"
// @Success      200 {object} models.TodoOutput "Задача создана"
// @Failure      401 {object} string "Не авторизирован"
// @Failure      400 {object} string "Некорректные данные"
// @Failure      500 {object} string "Внутренняя ошибка сервера"
// @Router       /todo [post]
func (r *Route) Create(c *gin.Context) {
	var input models.TodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		answer.SendError(c, "Неверный формат входных данных", answer.BadRequest)
		return
	}
	session, err := c.Cookie("X-Session-ID")
	if err != nil {
		answer.SendError(c, "Сессия отсутствует", answer.Unauthorized)
		return
	}
	out, code := r.action.Create(&input, session)
	if code != answer.OK {
		answer.SendError(c, "Ошибка при создании задачи", code)
		return
	}
	answer.SendResponseSuccess(c, out, code)
}

// @Summary      Получить список задач
// @Description  Получение всех задач для текущей сессии
// @Tags         Задачи
// @Produce      json
// @Param        // @Param X-Session-ID header string true "Идентификатор сессии (указывается в cookie, но описан как header для Swagger)" string true "Идентификатор сессии пользователя"
// @Success      200 {object} []models.TodoOutput "Список задач получен"
// @Failure      401 {object} string "Не авторизирован"
// @Failure      500 {object} string "Внутренняя ошибка сервера"
// @Router       /todo [get]
func (r *Route) Get(c *gin.Context) {
	session, err := c.Cookie("X-Session-ID")
	if err != nil {
		answer.SendError(c, "Сессия отсутствует", answer.Unauthorized)
		return
	}
	out, code := r.action.Get(session)
	if code != answer.OK {
		answer.SendError(c, "Ошибка при получении задач", code)
		return
	}
	answer.SendResponseSuccess(c, out, code)
}

// @Summary      Удалить задачу
// @Description  Удаляет задачу по ID
// @Tags         Задачи
// @Param        id path int true "ID задачи"
// @Param        // @Param X-Session-ID header string true "Идентификатор сессии (указывается в cookie, но описан как header для Swagger)" string true "Идентификатор сессии пользователя"
// @Success      204 {object} string "NoContent"
// @Failure      400 {object} string "Некорректные данные"
// @Failure      500 {object} string "Внутренняя ошибка сервера"
// @Router       /todo/:id [delete]
func (r *Route) Delete(c *gin.Context) {
	session, err := c.Cookie("X-Session-ID")
	if err != nil {
		answer.SendError(c, "Сессия отсутствует", answer.Unauthorized)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		answer.SendError(c, "Неверный ID", answer.BadRequest)
		return
	}
	code := r.action.Delete(id, session)
	if code != answer.OK {
		answer.SendError(c, "Ошибка при удалении задачи", code)
		return
	}
	answer.SendResponseSuccess(c, answer.NoContent, code)
}

// @Summary      Обновить задачу
// @Description  Полное обновление задачи по ID
// @Tags         Задачи
// @Accept       json
// @Param        id path int true "ID задачи"
// @Param        // @Param X-Session-ID header string true "Идентификатор сессии (указывается в cookie, но описан как header для Swagger)" string true "Идентификатор сессии пользователя"
// @Param        input body models.TodoInput true "Новые данные задачи"
// @Success      204 {object} string "NoContent"
// @Failure      400 {object} string "Некорректные данные"
// @Failure      500 {object} string "Внутренняя ошибка сервера"
// @Router       /todo/:id [put]
func (r *Route) Update(c *gin.Context) {
	session, err := c.Cookie("X-Session-ID")
	if err != nil {
		answer.SendError(c, "Сессия отсутствует", answer.Unauthorized)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		answer.SendError(c, "Неверный ID", answer.BadRequest)
		return
	}
	var input models.TodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		answer.SendError(c, "Неверный формат входных данных", answer.BadRequest)
		return
	}
	code := r.action.Update(&input, id, session)
	if code != answer.OK {
		answer.SendError(c, "Ошибка при обновлении задачи", code)
		return
	}
	answer.SendResponseSuccess(c, answer.NoContent, code)
}
