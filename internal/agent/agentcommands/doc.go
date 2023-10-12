// Package agentcommands реализует обработку пользовательских комманд, переданных от контроллера агента.
// Основной тип пакета AgentCommands взаимодействует с внутренним состоянием agentstate.State(где хранятся данные
// прив. записей тек.пользователя) и grpc агентом, подключенным к соотв. серверу с прив.данными пользователя.
//
// Состояние хранится только активной сессии авторизов.(!) пользователя(после получения информации с сервера).
// Т.е. чтобы получить данные нужно: зарегистр-ся\авториз-ся и потом либо запросить все записи с сервера либо добавить
// новые на клиенте.
// Без авторизации agentcommands будет возвращать ошибку на методах работы с данными.
// При смене пользователя(вызове методов регистр-ии\автор-ии после успешной авторизации) стейт будет полностью очищен
// от данных предыдущего пользователя.
// Факт авторизации хранится в переменной isAuthorized.
// Выйти из авториз. можно только перезапуском программы.
//
// agentcommands может добавлять прив. данные пользователя только(!) при наличии соединения с сервером, иначе
// возвращается ошибка. (например, при попытке создания текст. записи)
//
// go файлы пакета, для удобства, разделены по классам описываемых методов: auth - методы авторизации, create - методы
// создания польз.данных и т.п.
//
// Ввод и вывод, как и в контроллере агента реализован через reader/writer объекты, прежде всего для удобства тест-ия.
package agentcommands
