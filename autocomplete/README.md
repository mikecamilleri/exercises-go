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

## Design and References

Throughout my development of this web service, the [offical Go docs](https://golang.org/doc/) and [this intro book](https://www.golang-book.com/books/intro) were used as references.

This web service consists of three main parts:

### 1. The Datastore

The words and their counts are stored in memory in a "trie" data structure. I thought that a tree data structure would be a relatively straightforward and very fast way to store the words and values, so I did a quick Google search which turned up the [Wikipedia page for trie](https://en.wikipedia.org/wiki/Trie) which I used as reference. I decided to mark the end of words (and store their count) in the tree using a `wordCount` field on the node.  

### 2. The Word Extractor

The word extractor is run as a goroutine and emits words on a channel. Words are extracted from a plaintext file by first "scanning" the file for words (whitespace delimited), each word is then "cleaned" and finally sent out on the channel. I had to make some decisions regarding what a "word" is, which are described in the comments. With more work, the extraction and cleaning could likely be made more efficient.

### 3. The HTTP API

It gets the job done and is implemented directly in `main.go`.

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
