# Pacotes

Pacotes
 ```
 go get -u github.com/go-chi/chi
 go get -u github.com/jinzhu/gorm
 go get -u github.com/dgrijalva/jwt-go
 go get -u github.com/gin-gonic/gin
 go get -u github.com/denisenkom/go-mssqldb
 ```

# Desafio
O servidor de Quake 3 Arena gera um arquivo de LOG chamado **games.log**.

Nele fica registrado todas as informações dos jogos, como quando um jogo começa, quando termina, as mortes, atualizações de profile, itens, etc.

O desafio consiste em construir um parser capaz de ler o arquivo de LOG e expor uma API para consulta de informações dos jogos.

O parser deve ser capaz de ler o arquivo, agrupar as informações de cada jogo e coletar as informações referentes as mortes e pontuações de cada jogador em determinada partida.

Já a API deve ser capaz de expor um método onde o usuário possa ver a pontuação e mortes de todos os jogos.


# Exemplo
    21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
  
*O player "Isgalamido" morreu pois estava ferido e caiu de uma altura que o matou.*

    2:22 Kill: 3 2 10: Isgalamido killed Dono da Bola by MOD_RAILGUN
  
*O player "Isgalamido" matou o player "Dono da Bola" usando a arma Railgun.*
  

### Observações
- Quando o `<world>` mata o player ele perde -1 kill.
- `<world>` não é um player e não deve aparecer na lista de players e nem no dicionário de kills.
- `total_kills` são os kills dos games, isso inclui mortes do `<world>`.


Output
 ```
 Numero de Rounds:  21
Jogo :  1
Numero de kills no jogo:  0
Jogadores :  [{Isgalamido}]
Score dos Jogadores:  [{Isgalamido 0}]
Jogo :  2
Numero de kills no jogo:  11
Jogadores :  [{Isgalamido} {Dono da Bola} {Mocinha}]
Score dos Jogadores:  [{Isgalamido 3} {Dono da Bola 0} {Mocinha 0}]
Jogo :  3
Numero de kills no jogo:  4
Jogadores :  [{Dono da Bola} {Mocinha} {Isgalamido} {Zeh}]
Score dos Jogadores:  [{Dono da Bola 0} {Mocinha 0} {Isgalamido 1} {Zeh 0}]
Jogo :  4
Numero de kills no jogo:  105
Jogadores :  [{Dono da Bola} {Isgalamido} {Zeh} {Assasinu Credi}]
Score dos Jogadores:  [{Zeh 22} {Assasinu Credi 16} {Dono da Bola 20} {Isgalamido 27}]
Jogo :  5
Numero de kills no jogo:  14
Jogadores :  [{Dono da Bola} {Isgalamido} {Zeh} {Assasinu Credi}]
Score dos Jogadores:  [{Dono da Bola 0} {Isgalamido 2} {Zeh 2} {Assasinu Credi 5}]
Jogo :  6
Numero de kills no jogo:  29
Jogadores :  [{Fasano Again} {Oootsimo} {Zeh} {Dono da Bola} {UnnamedPlayer} {Maluquinho} {Isgalamido} {Assasinu Credi} {Mal}]
Score dos Jogadores:  [{Dono da Bola 2} {Oootsimo 9} {Isgalamido 4} {Zeh 8} {UnnamedPlayer 0} {Maluquinho 1} {Assasinu Credi 1} {Mal 0} {Fasano Again 0}]
Jogo :  7
Numero de kills no jogo:  130
Jogadores :  [{Chessus!} {Chessus} {Oootsimo} {Isgalamido} {Zeh} {Dono da Bola} {Mal} {Assasinu Credi}]
Score dos Jogadores:  [{Chessus! 0} {Chessus 0} {Oootsimo 24} {Isgalamido 20} {Zeh 14} {Dono da Bola 14} {Mal 9} {Assasinu Credi 22}]
Jogo :  8
Numero de kills no jogo:  89
Jogadores :  [{Oootsimo} {Isgalamido} {Zeh} {Dono da Bola} {Mal} {Assasinu Credi}]
Score dos Jogadores:  [{Oootsimo 17} {Isgalamido 24} {Zeh 15} {Dono da Bola 5} {Mal 1} {Assasinu Credi 12}]
Jogo :  9
Numero de kills no jogo:  67
Jogadores :  [{Dono da Bola} {Mal} {Assasinu Credi} {Chessus!} {Chessus} {Oootsimo} {Isgalamido} {Zeh}]
Score dos Jogadores:  [{Zeh 15} {Dono da Bola 3} {Mal 7} {Assasinu Credi 11} {Chessus! 0} {Chessus 9} {Oootsimo 9} {Isgalamido 2}]
Jogo :  10
Numero de kills no jogo:  60
Jogadores :  [{Dono da Bola} {Zeh} {Chessus} {Mal} {Assasinu Credi} {Isgalamido} {Oootsimo}]
Score dos Jogadores:  [{Oootsimo 1} {Dono da Bola 5} {Zeh 9} {Chessus 6} {Mal 6} {Assasinu Credi 5} {Isgalamido 10}]
Jogo :  11
Numero de kills no jogo:  20
Jogadores :  [{Isgalamido} {Zeh} {Oootsimo} {Chessus} {Assasinu Credi} {UnnamedPlayer} {Mal} {Dono da Bola}]
Score dos Jogadores:  [{Chessus 0} {Assasinu Credi 0} {UnnamedPlayer 0} {Mal 0} {Dono da Bola 1} {Isgalamido 7} {Zeh 0} {Oootsimo 4}]
Jogo :  12
Numero de kills no jogo:  160
Jogadores :  [{Oootsimo} {Chessus} {Assasinu Credi} {Mal} {Isgalamido} {Dono da Bola} {Zeh}]
Score dos Jogadores:  [{Assasinu Credi 23} {Mal 8} {Isgalamido 26} {Dono da Bola 11} {Zeh 14} {Oootsimo 22} {Chessus 17}]
Jogo :  13
Numero de kills no jogo:  6
Jogadores :  [{Isgalamido} {Dono da Bola} {Zeh} {Oootsimo} {Chessus} {Assasinu Credi} {Mal}]
Score dos Jogadores:  [{Assasinu Credi 0} {Mal 0} {Isgalamido 0} {Dono da Bola 0} {Zeh 2} {Oootsimo 2} {Chessus 0}]
Jogo :  14
Numero de kills no jogo:  122
Jogadores :  [{Mal} {Isgalamido} {Dono da Bola} {Zeh} {Oootsimo} {Chessus} {Assasinu Credi}]
Score dos Jogadores:  [{Oootsimo 12} {Chessus 10} {Assasinu Credi 12} {Mal 6} {Isgalamido 25} {Dono da Bola 9} {Zeh 12}]
Jogo :  15
Numero de kills no jogo:  3
Jogadores :  [{Oootsimo} {Zeh} {Assasinu Credi} {Dono da Bola} {Fasano Again} {Isgalamido}]
Score dos Jogadores:  [{Zeh 0} {Assasinu Credi 0} {Dono da Bola 0} {Fasano Again 0} {Isgalamido 0} {Oootsimo 0}]
Jogo :  16
Numero de kills no jogo:  0
Jogadores :  [{Oootsimo} {Isgalamido} {Assasinu Credi} {Zeh} {Dono da Bola}]
Score dos Jogadores:  [{Dono da Bola 0} {Oootsimo 0} {Isgalamido 0} {Assasinu Credi 0} {Zeh 0}]
Jogo :  17
Numero de kills no jogo:  13
Jogadores :  [{Dono da Bola} {Oootsimo} {Isgalamido} {Assasinu Credi} {Zeh} {UnnamedPlayer} {Mal}]
Score dos Jogadores:  [{Dono da Bola 0} {Oootsimo 2} {Isgalamido 1} {Assasinu Credi 0} {Zeh 1} {UnnamedPlayer 0} {Mal 0}]
Jogo :  18
Numero de kills no jogo:  7
Jogadores :  [{Isgalamido} {Assasinu Credi} {Zeh} {Mal} {Dono da Bola} {Oootsimo}]
Score dos Jogadores:  [{Zeh 2} {Mal 0} {Dono da Bola 0} {Oootsimo 0} {Isgalamido 1} {Assasinu Credi 2}]
Jogo :  19
Numero de kills no jogo:  95
Jogadores :  [{Isgalamido} {Oootsimo} {Dono da Bola} {Assasinu Credi} {Zeh} {Mal}]
Score dos Jogadores:  [{Zeh 21} {Mal 8} {Isgalamido 15} {Oootsimo 11} {Dono da Bola 15} {Assasinu Credi 12}]
Jogo :  20
Numero de kills no jogo:  3
Jogadores :  [{Isgalamido} {Oootsimo} {Dono da Bola} {Assasinu Credi} {Zeh} {Mal}]
Score dos Jogadores:  [{Dono da Bola 2} {Assasinu Credi 0} {Zeh 0} {Mal 0} {Isgalamido 0} {Oootsimo 1}]
Jogo :  21
Numero de kills no jogo:  134
Jogadores :  [{Zeh} {Mal} {Isgalamido} {Oootsimo} {Dono da Bola} {Assasinu Credi}]
Score dos Jogadores:  [{Mal 12} {Isgalamido 19} {Oootsimo 25} {Dono da Bola 18} {Assasinu Credi 22} {Zeh 21}]
```


# Retorno

```javascript
{"result":{"Games":[{"game":1,"TotalKill":0,"Players":[{"Name":"Isgalamido"}],"Kills":[{"Player":"Isgalamido","Score":0}]},{"game":2,"TotalKill":11,"Players":[{"Name":"Isgalamido"},{"Name":"Dono da Bola"},{"Name":"Mocinha"}],"Kills":[{"Player":"Dono da Bola","Score":0},{"Player":"Mocinha","Score":0},{"Player":"Isgalamido","Score":3}]},{"game":3,"TotalKill":4,"Players":[{"Name":"Mocinha"},{"Name":"Isgalamido"},{"Name":"Zeh"},{"Name":"Dono da Bola"}],"Kills":[{"Player":"Dono da Bola","Score":0},{"Player":"Mocinha","Score":0},{"Player":"Isgalamido","Score":1},{"Player":"Zeh","Score":0}]},{"game":4,"TotalKill":105,"Players":[{"Name":"Dono da Bola"},{"Name":"Isgalamido"},{"Name":"Zeh"},{"Name":"Assasinu Credi"}],"Kills":[{"Player":"Dono da Bola","Score":20},{"Player":"Isgalamido","Score":27},{"Player":"Zeh","Score":22},{"Player":"Assasinu Credi","Score":16}]},{"game":5,"TotalKill":14,"Players":[{"Name":"Dono da Bola"},{"Name":"Isgalamido"},{"Name":"Zeh"},{"Name":"Assasinu Credi"}],"Kills":[{"Player":"Dono da Bola","Score":0},{"Player":"Isgalamido","Score":2},{"Player":"Zeh","Score":2},{"Player":"Assasinu Credi","Score":5}]},{"game":6,"TotalKill":29,"Players":[{"Name":"UnnamedPlayer"},{"Name":"Assasinu Credi"},{"Name":"Fasano Again"},{"Name":"Isgalamido"},{"Name":"Dono da Bola"},{"Name":"Mal"},{"Name":"Oootsimo"},{"Name":"Zeh"},{"Name":"Maluquinho"}],"Kills":[{"Player":"UnnamedPlayer","Score":0},{"Player":"Maluquinho","Score":1},{"Player":"Mal","Score":0},{"Player":"Isgalamido","Score":4},{"Player":"Zeh","Score":8},{"Player":"Dono da Bola","Score":2},{"Player":"Assasinu Credi","Score":1},{"Player":"Fasano Again","Score":0},{"Player":"Oootsimo","Score":9}]},{"game":7,"TotalKill":130,"Players":[{"Name":"Chessus"},{"Name":"Oootsimo"},{"Name":"Isgalamido"},{"Name":"Zeh"},{"Name":"Dono da Bola"},{"Name":"Mal"},{"Name":"Assasinu Credi"},{"Name":"Chessus!"}],"Kills":[{"Player":"Dono da Bola","Score":14},{"Player":"Mal","Score":9},{"Player":"Assasinu Credi","Score":22},{"Player":"Chessus!","Score":0},{"Player":"Chessus","Score":0},{"Player":"Oootsimo","Score":24},{"Player":"Isgalamido","Score":20},{"Player":"Zeh","Score":14}]},{"game":8,"TotalKill":89,"Players":[{"Name":"Isgalamido"},{"Name":"Zeh"},{"Name":"Dono da Bola"},{"Name":"Mal"},{"Name":"Assasinu Credi"},{"Name":"Oootsimo"}],"Kills":[{"Player":"Mal","Score":1},{"Player":"Assasinu Credi","Score":12},{"Player":"Oootsimo","Score":17},{"Player":"Isgalamido","Score":24},{"Player":"Zeh","Score":15},{"Player":"Dono da Bola","Score":5}]},{"game":9,"TotalKill":67,"Players":[{"Name":"Chessus"},{"Name":"Oootsimo"},{"Name":"Isgalamido"},{"Name":"Zeh"},{"Name":"Dono da Bola"},{"Name":"Mal"},{"Name":"Assasinu Credi"},{"Name":"Chessus!"}],"Kills":[{"Player":"Dono da Bola","Score":3},{"Player":"Mal","Score":7},{"Player":"Assasinu Credi","Score":11},{"Player":"Chessus!","Score":0},{"Player":"Chessus","Score":9},{"Player":"Oootsimo","Score":9},{"Player":"Isgalamido","Score":2},{"Player":"Zeh","Score":15}]},{"game":10,"TotalKill":60,"Players":[{"Name":"Zeh"},{"Name":"Chessus"},{"Name":"Mal"},{"Name":"Assasinu Credi"},{"Name":"Isgalamido"},{"Name":"Oootsimo"},{"Name":"Dono da Bola"}],"Kills":[{"Player":"Zeh","Score":9},{"Player":"Chessus","Score":6},{"Player":"Mal","Score":6},{"Player":"Assasinu Credi","Score":5},{"Player":"Isgalamido","Score":10},{"Player":"Oootsimo","Score":1},{"Player":"Dono da Bola","Score":5}]},{"game":11,"TotalKill":20,"Players":[{"Name":"UnnamedPlayer"},{"Name":"Mal"},{"Name":"Dono da Bola"},{"Name":"Isgalamido"},{"Name":"Zeh"},{"Name":"Oootsimo"},{"Name":"Chessus"},{"Name":"Assasinu Credi"}],"Kills":[{"Player":"Assasinu Credi","Score":0},{"Player":"UnnamedPlayer","Score":0},{"Player":"Mal","Score":0},{"Player":"Dono da Bola","Score":1},{"Player":"Isgalamido","Score":7},{"Player":"Zeh","Score":0},{"Player":"Oootsimo","Score":4},{"Player":"Chessus","Score":0}]},{"game":12,"TotalKill":160,"Players":[{"Name":"Dono da Bola"},{"Name":"Zeh"},{"Name":"Oootsimo"},{"Name":"Chessus"},{"Name":"Assasinu Credi"},{"Name":"Mal"},{"Name":"Isgalamido"}],"Kills":[{"Player":"Isgalamido","Score":26},{"Player":"Dono da Bola","Score":11},{"Player":"Zeh","Score":14},{"Player":"Oootsimo","Score":22},{"Player":"Chessus","Score":17},{"Player":"Assasinu Credi","Score":23},{"Player":"Mal","Score":8}]},{"game":13,"TotalKill":6,"Players":[{"Name":"Mal"},{"Name":"Isgalamido"},{"Name":"Dono da Bola"},{"Name":"Zeh"},{"Name":"Oootsimo"},{"Name":"Chessus"},{"Name":"Assasinu Credi"}],"Kills":[{"Player":"Chessus","Score":0},{"Player":"Assasinu Credi","Score":0},{"Player":"Mal","Score":0},{"Player":"Isgalamido","Score":0},{"Player":"Dono da Bola","Score":0},{"Player":"Zeh","Score":2},{"Player":"Oootsimo","Score":2}]},{"game":14,"TotalKill":122,"Players":[{"Name":"Zeh"},{"Name":"Oootsimo"},{"Name":"Chessus"},{"Name":"Assasinu Credi"},{"Name":"Mal"},{"Name":"Isgalamido"},{"Name":"Dono da Bola"}],"Kills":[{"Player":"Chessus","Score":10},{"Player":"Assasinu Credi","Score":12},{"Player":"Mal","Score":6},{"Player":"Isgalamido","Score":25},{"Player":"Dono da Bola","Score":9},{"Player":"Zeh","Score":12},{"Player":"Oootsimo","Score":12}]},{"game":15,"TotalKill":3,"Players":[{"Name":"Dono da Bola"},{"Name":"Fasano Again"},{"Name":"Isgalamido"},{"Name":"Oootsimo"},{"Name":"Zeh"},{"Name":"Assasinu Credi"}],"Kills":[{"Player":"Oootsimo","Score":0},{"Player":"Zeh","Score":0},{"Player":"Assasinu Credi","Score":0},{"Player":"Dono da Bola","Score":0},{"Player":"Fasano Again","Score":0},{"Player":"Isgalamido","Score":0}]},{"game":16,"TotalKill":0,"Players":[{"Name":"Oootsimo"},{"Name":"Isgalamido"},{"Name":"Assasinu Credi"},{"Name":"Zeh"},{"Name":"Dono da Bola"}],"Kills":[{"Player":"Zeh","Score":0},{"Player":"Dono da Bola","Score":0},{"Player":"Oootsimo","Score":0},{"Player":"Isgalamido","Score":0},{"Player":"Assasinu Credi","Score":0}]},{"game":17,"TotalKill":13,"Players":[{"Name":"Dono da Bola"},{"Name":"Oootsimo"},{"Name":"Isgalamido"},{"Name":"Assasinu Credi"},{"Name":"Zeh"},{"Name":"UnnamedPlayer"},{"Name":"Mal"}],"Kills":[{"Player":"Dono da Bola","Score":0},{"Player":"Oootsimo","Score":2},{"Player":"Isgalamido","Score":1},{"Player":"Assasinu Credi","Score":0},{"Player":"Zeh","Score":1},{"Player":"UnnamedPlayer","Score":0},{"Player":"Mal","Score":0}]},{"game":18,"TotalKill":7,"Players":[{"Name":"Zeh"},{"Name":"Mal"},{"Name":"Dono da Bola"},{"Name":"Oootsimo"},{"Name":"Isgalamido"},{"Name":"Assasinu Credi"}],"Kills":[{"Player":"Assasinu Credi","Score":2},{"Player":"Zeh","Score":2},{"Player":"Mal","Score":0},{"Player":"Dono da Bola","Score":0},{"Player":"Oootsimo","Score":0},{"Player":"Isgalamido","Score":1}]},{"game":19,"TotalKill":95,"Players":[{"Name":"Dono da Bola"},{"Name":"Assasinu Credi"},{"Name":"Zeh"},{"Name":"Mal"},{"Name":"Isgalamido"},{"Name":"Oootsimo"}],"Kills":[{"Player":"Mal","Score":8},{"Player":"Isgalamido","Score":15},{"Player":"Oootsimo","Score":11},{"Player":"Dono da Bola","Score":15},{"Player":"Assasinu Credi","Score":12},{"Player":"Zeh","Score":21}]},{"game":20,"TotalKill":3,"Players":[{"Name":"Oootsimo"},{"Name":"Dono da Bola"},{"Name":"Assasinu Credi"},{"Name":"Zeh"},{"Name":"Mal"},{"Name":"Isgalamido"}],"Kills":[{"Player":"Isgalamido","Score":0},{"Player":"Oootsimo","Score":1},{"Player":"Dono da Bola","Score":2},{"Player":"Assasinu Credi","Score":0},{"Player":"Zeh","Score":0},{"Player":"Mal","Score":0}]},{"game":21,"TotalKill":134,"Players":[{"Name":"Oootsimo"},{"Name":"Dono da Bola"},{"Name":"Assasinu Credi"},{"Name":"Zeh"},{"Name":"Mal"},{"Name":"Isgalamido"}],"Kills":[{"Player":"Dono da Bola","Score":18},{"Player":"Assasinu Credi","Score":22},{"Player":"Zeh","Score":21},{"Player":"Mal","Score":12},{"Player":"Isgalamido","Score":19},{"Player":"Oootsimo","Score":25}]}]}}
```