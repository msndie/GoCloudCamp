# GoCloudCamp
Тестовое задание для поступления в GoCloudCamp
<!--
# 1. Вопросы для разогрева

- Опишите самую интересную задачу в программировании, которую вам приходилось решать?

  Когда я писал проект <a href="https://github.com/msndie/cub3d">cub3d</a> (Raycast engine), мне захотелось сделать двери, но не просто стену которая исчезнет если нажать определенную клавишу, а такие же как в Wolfenstein3D, это пожалуй было самой интересной задачей, а так как я далеко не математик мне пришлось с ней повозиться.

- Расскажите о своем самом большом факапе? Что вы предприняли для решения проблемы?

  Во время работы над <a href="https://github.com/msndie/minishell">minishell</a> (реализация оболочки терминала, референс bash), а точнее уже на заключающей стадии проекта, оказалось что вся внутренняя логика запуска процессов не соответствует логике реализации heredoc, пришлось реализовать ее через костыль, иначе пришлось бы переписывать половину кода, а времени на это не оставалось. Работает отлично, но внутренне все же не так как bash.

- Каковы ваши ожидания от участия в буткемпе?

  Я не писал на Go, пишу на C, C++, Java, для меня это в первую очередь опыт, как в плане первой работы в IT, так и в изучении языка. Работа с новыми инструментами и прокачка навыков обращения со старыми.
  
# 2. Distributed config
-->
<details>
<summary>protobuf</summary>

```proto
syntax = "proto3";
import "google/protobuf/empty.proto";

message Property {
  string key = 1;
  string value = 2;
}

message Config {
  string service = 1;
  repeated Property data = 2;
}

message Configs {
  repeated Config configs = 1;
}

message ConfigNameRequest {
  string service = 1;
}

service ConfigService {
  rpc addConfig(Config) returns (Config);
  rpc getConfig(ConfigNameRequest) returns (Config);
  rpc getAllVersionsOfConfig(ConfigNameRequest) returns (Configs);
  rpc getAllConfigs(google.protobuf.Empty) returns (Configs);
  rpc updateConfig(Config) returns (Config);
  rpc deleteConfig(ConfigNameRequest) returns (Config);
  rpc useConfig(ConfigNameRequest) returns (stream Config);
  rpc stopConfigUseForAll(ConfigNameRequest) returns (google.protobuf.Empty);
}
```
</details>
<details>
<summary>Java</summary>

### Сервис

Все конфиги сохраняются в mongodb, верисионирование я реализовал следующим образом:

- eсли полученные данные отлючаются от текущего представления конфига в бд, я создаю "entity" которое отличается от дефолтного класса только переменной времени создания и добавляю его в коллекцию.

Когда мне нужно достать конкретный конфиг я сортирую всю коллекцию по времени создания в обратном порядке и достаю первый документ. Я добавил метод для получения всех версий одного конфига, они также отсортированы. Надеюсь это хоть немного похоже на то как вы это представляли.

Я реализовал метод useConfig, он возвращает stream, подписчиков на однин конфиг может быть много, каждого подписчика я прослушиваю на отмену/отключение. При изменении конфига я отправляю новый конфиг с помощью этих сохраненных потоков. Если подпичиков не останется, то конфиг может быть удален.

Также добавил метод для принудительного отключения всех подписчиков для конкретного конфига.

Сам сервис покрыл тестами не затрагивая слои бизнес логики и репозитория, в тестах они закрыты заглушками, покрытие 95%.

### Тестовый клиент

В тестовом клиенте создается или обновляется до дефолтного состояния конфиг для "Test app". После происходит подписка на данный конфиг, клиент начинает ждать 2,5 сек и с интервалами в 0,5с писать в консоль сообщение ожидания, которое было полученно из конфига, по истечении этого времени клиент вызывает метод updateConfig, из стрима получет новые данные и обновляет сообщения согласно этим данным. После этого клиент вызывает метод stopConfigUseForAll, поток закрывается. Клиент все свои действия логирует в консоль.

## Запуск

Добавил docker файлы для сервиса и тестового клиента, также для удобства добавил мейкфайл с помощью которого можно все это запустить. Сервис работает на порту 9090.

### Make

```
make all
or
make
```
Запустит в докере mongodb и сам сервис. Компиляция также происходит в докере. mongodb запустится на нестандартном порту - 27018.

```
make down
```
Команда остановит запущенные контейнеры. (Если у вас есть другие запущенные контейнеры лучше остановите их руками)

```
make client
```
Запустит в докере описанного выше клиента.

### Manual
```
mvn -f Service/pom.xml clean package
mvn -f Client/pom.xml clean package
java -jar Service/target/ConfigurationService-1.0-jar-with-dependencies.jar
java -jar Client/target/Client-1.0-jar-with-dependencies.jar
```
### Сервис доступен по адресу localhost:9090

</details>

<details>
<summary>Go</summary>
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ConfigService.proto
</details>