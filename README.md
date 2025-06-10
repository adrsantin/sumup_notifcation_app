# Sumup Notification App

Aplicação designada para envio de notificações de pagamentos realizados pelos usuários. Desenvolvida como teste técnico por Adriano Tetsuaki Ogawa Santin

## Fluxo do sistema

- Existe um endpoint http: POST /payments/notification definido na camada de api. Esse endpoint recebe um body contendo um "user_id" do tipo inteiro e um "amount" de ponto decimal representando o valor do pagamento efetuado (1)
- A api envia esses dados para a camada de serviço (2) que busca, através da camada do repositório os dados do usuário no banco de dados (3) e (4). Esses dados incluem os tipos de notificação que o usuário vai receber.
- Essas informações são então enviadas, através de um Produtor, para um tópico do kafka (5) e (6).
- Dentro da aplicação temos também um worker, que roda em uma goroutine separada da principal, que consome as mensagens que chegam no tópico do kafka (7).
- Ao consumir a mensagem, os dados são enviados à camada de serviço (8) que separa, através do tipo de notificação, como vai fazer o envio da mesma. Isso é feito através de um design pattern de estratégia, em que existe um único método SendNotification que é chamado no serviço, mas que dentro dele, existem múltiplas implementações do método dependendo do tipo de notificação recebido.

### Definições

- Assumimos que cada usuário possui cadastro com dados e uma seleção de tipos de notificação de sua preferência. Essas informações ficam salvas em um banco de dados mysql.
- Não foi desenvolvido o ponto em que a aplicação por fim faz o envio das notificações aos usuários, mas essa aplicação foca em todo o design e arquitetura até esse ponto final, assumindo que ao final, seria feita uma chamada simples a algum sistema externo para efetuar o envio em si.
- Com o design pattern de estratégia, é bem fácil de implementar novos meios de notificação, bastando a criação de um novo método dentro de `internal/business/notifications/notifications.go`, e de adicionar o nome da nova notificação como válida no banco de dados e dentro de `internal/entities/notifications.go`.
- Foi utilizado também um sistema assíncrono para o envio das notificações usando o kafka. Isso foi feito para manter rápida a resposta ao sistema de pagamentos que solicita o envio das notificações, e ainda garantir o envio das mesmas de acordo com as configurações do próprio kafka.
- Usando o docker-compose, conseguimos subir localmente toda a infraestrutura que a aplicação utilizaria (mysql e kafka). Para o banco de dados, também subimos com um script para popular o banco com alguns dados de teste iniciais. Porém, em um ambiente de produção, não precisariamos subir essa infra, já que provavelmente o banco e os tópicos do kafka deveriam ser administrados externamente.


### Para rodar localmente

- É necessário Docker e docker compose
- Então, basta rodar o comando `docker compose up`
- A aplicação deve rodar em `localhost:8080`

## Pontos a melhorar AKA o que eu faria com mais tempo
- Testes unitários. Foram adicionados testes unitários para vários métodos, porém não em tudo. Colocaria mais testes unitários até 100% de cobertura de código.
- Observabilidade. Eu melhoraria como os logs estão até o momento, vendo que pontos merecem mais atenção. Também implementaria o lançamento de métricas para um datadog ou similar, para verificar melhor a saúde do sistema.
- Implementação do kafka. O kafka é algo bem mais complexo e poderoso do que eu implementei aqui, e infelizmente eu não conheço tão a fundo, então é algo que acho que merece mais estudo para poder melhorar. 
- Configurações do sistema. As implementações das ferramentas externas (o mysql e kafka) foram feitas de forma bem simples e com configurações padrão. Seria bom implementar um sistema para ler os dados de configuração de algum outro lugar externo, como variáveis de ambiente, ou algum sistema de vault (principalmente para dados mais sensíveis)
