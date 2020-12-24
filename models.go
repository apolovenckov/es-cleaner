package main

// Indices struct. Для декодинга JSON формата и извлечения списка индексов и даты создания.
type Indices []struct {
	Index        string `json:"index"`
	CreationDate string `json:"creation.date"`
}

// DeleteContainerNameLog struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму.
type DeleteContainerNameLog struct {
	Query QueryContainerNameLog `json:"query"`
}

// QueryContainerNameLog struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму.
type QueryContainerNameLog struct {
	Term TermContainerNameLog `json:"term"`
}

// TermContainerNameLog struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму.
type TermContainerNameLog struct {
	KubContName string `json:"kubernetes.container.name"`
}

// DeleteMessageLog struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму и сообщению.
type DeleteMessageLog struct {
	Query Query `json:"query"`
}

// Query struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму и сообщению.
type Query struct {
	Bool Bool `json:"bool"`
}

// Bool struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму и сообщению.
type Bool struct {
	Must   Must   `json:"must"`
	Filter Filter `json:"filter"`
}

// Must struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму и сообщению.
type Must struct {
	Match Match `json:"match"`
}

// Match struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму и сообщению.
type Match struct {
	MessageLog string `json:"message"`
}

// Filter struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму и сообщению.
type Filter struct {
	Term Term `json:"term"`
}

// Term struct. Для енкодинга в JSON и отправки в Elastic на удаление логов по контейнер нейму и сообщению.
type Term struct {
	ContainerName string `json:"kubernetes.container.name"`
}

// Task struct. Для декодинга JSON формата и извлечения таска на удаление.
type Task struct {
	Task string `json:"task"`
}

// TaskStatus struct. Для декодинга JSON формата и извлечения статуса таска на удаление.
type TaskStatus struct {
	Status bool `json:"completed"`
}

// DeleteComplitedTask struct. Для декодинга JSON формата и извлечения статуса таска на удаление.
type DeleteComplitedTask struct {
	Result string `json:"result"`
}

// ParamsToDeleteStruct struct. Для декодинга JSON формата и извлечения параметров из ENV на удаление.
type ParamsToDeleteStruct []struct {
	ContainerName string `json:"ContainerName"`
	Lifetime      int    `json:"Lifetime"`
	Message       string `json:"Message"`
}
