# Como deve funcionar?

## Estilos de cerveja (beerStyle)

- [ ] Criar CRUD de estilos de cerveja e temperatura
    - [ ] Listar
    - [ ] Cadastrar
    - [ ] Atualizar
    - [ ] Deletar

- [ ] Pegar um beerStyle
    - Busca estilo de cerveja pela temperatura
    - Busca playlist para esse estilo de cerveja
        - Busca no spotify playlist que contenha nome desse estilo no nome
    **Input:** { "temperature": -7 } int
    **Output:** { "beerStyle": string, "playlist": { "name": string, "tracks": Tracks[]{ "name": string, "artist": string, "link": string } } }

    **Regras de negócio**
    - Todos estilos de cerveja tem temperatura mínima e máxima
    - Cálculo para selecionar beerStyle ideal:
        - Média das temperaturas mais próxima do input
    - Retornar sempre 1 beerStyle
        **Restrições:**
        - Em caso de mais de 1 resultado:
            - Ordenar por ordem alfabética crescente
            - Retornar primeira
        - Em caso de não haver playlist que contenha o nome do beerstyle
            - Retornar o estilo com playlist vazia

