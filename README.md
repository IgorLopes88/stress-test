# Go Expert

Desafio **Sistema de Stress test** do curso **Pós Go Expert**.

**Objetivo:** Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

### Execução da **aplicação**
Para executar a aplicação execute o comando:
```
docker run igorlopes88/stress-test --url=http://google.com --requests=100 --concurrency=10
```

O resultado deverá ser esse:

```
-- RUN STRESS TEST -->
Url: http://google.com
Requests: 100
Concurrency: 10

Test finished!

TEST RESULT
Test duration: 10.23s
Successful Requests: 100

Status  Total
200     100
```

Pronto!


### Correções de Bugs
1. Inclusão do mutex para corrigir erro `fatal error: concurrent map writes`;
