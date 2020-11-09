## Тесты 

### Нагрузочное тестирование
#### О системе
* https://cloud.google.com/solutions/distributed-load-testing-using-kubernetes
* https://locust.io/

#### Быстрый старт
* https://docs.locust.io/en/stable/quickstart.html
* https://docs.locust.io/en/stable/writing-a-locustfile.html#validating-responses

#### Установка
Установка Python
```
sudo apt-get install python3 python3-pip
```
Установка системы Locust
```
sudo pip3 install locust
```

#### Запуск тестов на одной машине
Запуск из каталога
```
cd ./stress-tests
```
1. без веб-интерфейса
```
locust -f ./stress-test.py --host=http://127.0.0.1:4000 --no-web -c 1000 -r 100 -n 10000
```
2. с веб-интерфейсом localhost:8089
```
locust -f ./stress-test.py --host=http://127.0.0.1:4000 -c 1000 -r 100 -n 10000
```

