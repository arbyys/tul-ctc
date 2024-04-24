# Aplikace Benzínová stanice

- Golang
- Docker

# Origo
- Cars arrive at the gas station and wait in the queue for free station
- Total number of cars and their arrival time is configurable
- There are 4 types of stations: gas, diesel, LPG, electric
- Count of stations and their serve time is configurable as interval (e.g. 2–5s) and can be different for each type
- Each station can serve only one car at a time, serving time is chosen randomly from station's interval
- After the car is served, it goes to the cash register
- Count of cash registers and their handle time is configurable
- After the car is handled (random time from register handle time range) by the cash register, it leaves the station
- Program collects statistics about the time spent in the queue, time spent at the station and time spent at the cash register for every car
- Program prints the aggregate statistics at the end of the simulation

## ⚠ Implementation details

* All time values in `config.yaml` are in milliseconds by default; this could be changed in `random.go`, if really necessary
* Before car enters the _stand queue_ for its specific fuel type, it has to go through the _shared queue_
  * Shared queue is meant as a road between e.g. the highway and the main gas station area
  * If shared queue is full, spawning of new cars is paused
* Every stand/register/(shared) queue has its maximum length
  * If queue for the next part of the process is full, car waits and blocks the previous part
  * e.g.: Car _c_ finished refueling at stand _x_ and wants to go to the cash register _y_. _y_'s queue is full though, so _c_ waits until there is room for it. While waiting it still blocks _x_. Car that is behind _c_ in _x_'s queue cannot start refueling until _c_ leaves to the _y_'s queue.
* Queue lengths in `config.yaml` are actually Go channel capacities, meaning [one more car can fit into it](https://stackoverflow.com/a/25539742)
* Car always chooses stand/register with shortest queue (where there are currently least other cars)
* Times in `output.yaml` count time spent in queue, doing some action, but also waiting for queue (and blocking)
  * e.g. `register:handle_time` = time spent after refueling waiting __for__ _y_'s queue + time spent waiting __in__ _y_'s queue + time spent paying

# Přepsané zadání

- Benzínová stanice si udržeje FIFO s auty (config soubor), postupně random přijíždějí
- Pohony aut:
    - benzin
    - diesel
    - LPG
    - elektrika
- Každy typ pohonu má ve stanici *n* stojanů
    - Čas, jak dlouho se u stojanu auto zdrží je definován intervalem v config souboru
- Každý stojan dokáže obsluhovat pouze jedno auto naráz
- Jakmile auto natankuje, jde ke kase
    - Počet kas a čas obsluhy je opět definován intervalem v configu
- Až auto zaplatí, opouští čerpací stanici
- Program shromažďuje statistiky čekací doby pro:
    - každý druh pohonu
    - kasy
    - ...a pak z toho dělá yaml

# Struktura appky

- fronta aut (FIFO)
- obecná classa "obsluhovacího zařízení" 
    - z toho bude dědit classa pro každý typ pohonu (vyplňuje se hodnota intervalu)
    - má metodu "obsloužit auto", která vygeneruje náhodné číslo z intervalu
- 
# Příklad configu

```yaml
cars:
  count: 1000
  arrival_time_min: 1
  arrival_time_max: 2
  shared_queue_length_max: 5
stations:
  gas:
    count: 4
    serve_time_min: 2
    serve_time_max: 5
    queue_length_max: 3
  diesel:
    count: 2
    serve_time_min: 3
    serve_time_max: 6
    queue_length_max: 3
  lpg:
    count: 1
    serve_time_min: 4
    serve_time_max: 7
    queue_length_max: 3
  electric:
    count: 1
    serve_time_min: 5
    serve_time_max: 10
    queue_length_max: 3
registers:
  count: 2
  handle_time_min: 1
  handle_time_max: 3
  queue_length_max: 1

```

# Příklad výstupu

```yaml
stations:
  gas:
    total_cars: 50000
    total_time: 7520s
    avg_queue_time: 14s
    max_queue_time: 5s
  diesel:
    total_cars: 40000
    total_time: 6507s
    avg_queue_time: 15s
    max_queue_time: 5s
  lpg:
    total_cars: 30000
    total_time: 9000s
    avg_queue_time: 5s
    max_queue_time: 5s
  electric:
    total_cars: 30000
    total_time: 25055s
    avg_queue_time: 2s
    max_queue_time: 5s
registers:
  total_cars: 150000
  total_time: 170000s
  avg_queue_time: 2s
  max_queue_time: 5s

```