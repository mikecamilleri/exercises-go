# Autocomplete

This is a response to a programming challenge for a job interview. See CHALLENGE.md.

## Starting the Web Service

Assuming you have go installed, the easiest way to start the web service is to `cd` to the `autocomplete` directory and execute `go run .`

```
mike@Darwin-D autocomplete % pwd                                     
/Users/mike/Development/go/src/github.com/mikecamilleri/exercises-go/autocomplete
mike@Darwin-D autocomplete % go run .                                
2020/06/28 18:15:03 Hello World!
2020/06/28 18:15:03 extracting words from file and building database in memory ...
2020/06/28 18:15:03 listening on port 9000 ...
```

## How To Make a Request

By default the service runs on port `9000`. To request autocompletions, a `GET` request may be sent to `http://localhost:9000/autocomplete` with a query for `term` (i.e. `http://localhost:9000/autocomplete?term=th`). See below for example cURL requests.



## Example Results

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=th"

{"Completions":["the","that","this","thou","thy","thee","they","then","their","them","than","there","these","th","think","thus","though","therefore","those","thine","that's","there's","three","thought","thing"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=fr"

{"Completions":["from","france","friends","friend","french","free","friar","fresh","freely","francis","frown","frame","friendship","friendly","fruit","frederick","freedom","fright","froth","front","fran","frenchman","frowns","frail","fray"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=pi"

{"Completions":["pity","pistol","pisanio","piece","picture","pieces","pitch","pinch","pitiful","pierce","pit","piteous","pindarus","pin","pine","pitied","pick'd","pilgrimage","pillow","pick","pisa","pipe","pigeons","piercing","pilgrim"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=sh"

{"Completions":["she","should","show","shame","shalt","shakespeare","shallow","she's","shepherd","shylock","shows","shake","short","shape","shadow","shouldst","sharp","shut","shore","show'd","ship","shed","shortly","shot","shine"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=wu"

{"Completions":["wul"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=ar"

{"Completions":["are","art","arms","arm","armado","army","ariel","archbishop","argument","arviragus","arthur","armour","arm'd","arise","armed","arrest","articles","arriv'd","article","arrant","array","arras","archive","arrows","arts"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=il"

{"Completions":["ill","illinois","illyria","il","ills","ill-favour'd","ilion","illustrious","ill-favouredly","ill-beseeming","illusion","illusions","illustrate","illegitimate","ill-temper'd","ill-boding","ill-dispos'd","ild","ilium","ilbow","ils","illume","illumineth","illuminate","illumin'd"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=ne"

{"Completions":["never","news","near","new","ne'er","need","next","needs","neck","nerissa","nestor","neighbour","necessity","nephew","newly","needful","nest","neglect","neighbours","nearer","ned","necessary","negligence","nell","nearest"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=se"

{"Completions":["see","set","service","servant","sent","seen","send","seek","sea","serve","seem","sebastian","self","sense","seems","servants","senator","several","seven","seal","secret","senators","seat","seeming","seest"]}
```

```
mike@Darwin-D ~ % curl "http://localhost:9000/autocomplete?term=pl"

{"Completions":["place","please","play","pleasure","plain","pluck","plague","plantagenet","pleas'd","plot","plead","play'd","places","pleasures","players","plays","pleasant","pluck'd","plebeians","plant","pleasing","playing","plainly","pledge","pleases"]}
```