# Banco de dados:

- Para este projeto estamos utilizando um banco de dados postgresql, e para facilitar a criação e configuração é só rodar `docker-compose up` e o banco será criado na sua máquina local.

# Executar o projeto:

- Para executar projetos em Go, é necessário ter o pacote do `GO` instalado no seu ambiente.
- Seguir a documentação da instalação de acordo com o seu `SO`, package disponível em `https://go.dev/dl/`.
- Depois de instalado, rodar `go run main.go` e em seguida o projeto já devera estar rodando na porta `3333`
  <img src="./doc/evidence.png" alt="Evidência dos endpoints sendo executado"/>

## Insomnia

- Se você possui o insomnia instalado, faça a importação da collection dentro de `doc/insomnia_collection.json`.
  <img src="./doc/insomnia_evidence.png" alt="Evidência do Insomnia sendo executado"/>

<!-- ## Executar com Docker

- Para gerar uma imagem docker local, rodar `docker build -t api-mvc .`
- Para executar a imagem docker `docker run --rm -it api-mvc` -->
