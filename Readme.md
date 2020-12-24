### es-cleaner

Данный сервис предназначен для удаления ненужных логов контейнеров **kubernetes** по полю **kubernetes.container.name** через API ElasticSearch.

Сервис принимает из окружения (ENV) следующие переменные:

***ELASTIC_HOST*** - hostname или ip мастер ноды ElasticSearch,

***ELASTIC_PORT*** - port, на котором котором работает ElasticSearch,

***ELASTIC_INDICES_K8S*** - index pattern, по которому будет выполняться поиск индексов (сервис поддерживает только формат логов **kubernetes**, т.к. поиск и удаление документов из индекса выполняется по полю **kubernetes.container.name**),

***ELASTIC_API_FORMAT*** - формат, для работы с API (сервис поддерживает только **JSON** формат),

***ELASTIC_INDEX_LIST*** - переменная для извлечения из API параметров индекса (Сервис поддерживает только **index** - возвращает имя индекса, и **creation.date** - дата создания индекса. Данные переменные работают только в связке и должны быть заданы через запятую. Необходимо для сортировки индексов.),

***ELASTIC_URL*** - URL к API ElasticSearch, если не задан, то URL формируется автоматически из переменных **http://ELASTIC_HOST:ELASTIC_PORT/**,

***PARAMS_TO_DELETE*** - переменная которая принимает в себя имя контейнера (**ContainerName**), количество дней после которого можно удалять сообщения из индекса (**Lifetime**), и паттерн сообщения, которое будет удалено (**Message**). Данная переменная формируется в формате ***JSON***, в виде строки с экранированием, например:
```json
"[{\"ContainerName\":\"resource\",\"Lifetime\":7,\"Message\":\"*\"},{\"ContainerName\":\"cursus-bank\",\"Lifetime\":3,\"Message\":\"keklol\"},{\"ContainerName\":\"cursus-bank\",\"Lifetime\":14,\"Message\":\"lolkek\"}]"
```
Для корректного формирование используйте шаблон JSON формата:
```json
[
	{
		"ContainerName": "resource",
		"Lifetime": 7,
		"Message": "*"
	},
	{
		"ContainerName": "cursus-bank",
		"Lifetime": 3,
		"Message": "keklol"
	},
	{
		"ContainerName": "cursus-bank",
		"Lifetime": 14,
		"Message": "lolkek"
	}
]
```
Для перевода в строку и формирования енва можно воспользоваться онлайн сервисом - https://tools.knowledgewalls.com/jsontostring.

Сервис разбирает ***PARAMS_TO_DELETE***, исходя из поля **Lifetime** и списка индексов, которые он получает с помощью индекс паттерна (***ELASTIC_INDICES_K8S***) формирует список индексов из которых необходимо удалить либо сообщения (**Message**).  Если в **Message** задана **"\*"** будут удалены все логи указанного контейнер нейма (**ContainerName**). В поле **Message** **НЕЛЬЗЯ** использовать **"\*"** кроме как для обознаяения **ВСЕХ** сообщений.