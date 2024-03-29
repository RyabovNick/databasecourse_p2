# Описание

Игра представляет собой текстовый квест, в котором игроку предстоит управлять существом. Управление осуществляется в консоли при помощи цифр. То есть на каждом шаге игроку выдаются все его параметры и задается вопрос, что он собирается делать с вариантами ответа, пронумерованными цифрами.

Изначально у существа есть четыре основных параметра:

- Длина норы (далее Д) = 10
- Здоровье (далее З) = 100
- Уважение (далее У) = 20
- Вес (далее В) = 30

Цель игры – добиться уважения больше 100 и не дать основным параметрам упасть до нуля.

Игра происходит циклично и имитирует день и ночь. Ночью происходит изменение параметров по следующему принципу:

- Длина норы уменьшается на 2
- Здоровье увеличивается на 20
- Уважение уменьшается на 2
- Вес уменьшается на 5

Днем игрок может выбрать одно из следующих действий, которые отразятся на его показателях (подпункты показывают ветвление выбора):

- Копать нору
  - Интенсивно: Д+5, З-30
  - Лениво: Д+2, З-10
- Поесть травки
  - Жухлой: З+10, В+15
  - Зеленой:
    - Если У<30, то З-30
    - Если У>=30, то З+30, В+30
- Пойти подраться. Победа определяется генерацией случайного числа и должна равняться отношению веса существа к сумме весов. Например, если существо с весом 20 идет драться с существом весом 30, то вероятность победы будет 20/50, если у первого 70, а у второго 20 – 70/90 и т.д. Уважение прибавляется в зависимости от разницы в весе (тут нужно подобрать конкретные цифры): если существо слабее противника, дается больше уважения (например, 40), если они равны – среднее количество (например, 20), если существо ощутимо сильнее – минимум (5 или 10). Здоровье убывает аналогично, в зависимости от разницы в весе, тоже нужно подбирать цифры. Но основная логика в том, что можно пытаться идти со слабеньким существом на гиганта и получить при помощи рандома много уважения, с риском лишиться почти всего здоровья, а можно планомерно бить существ своего уровня с вероятностью успеха около 50% и без большого риска. Драться можно:
  - Со слабым существом (вес 30)
  - Со средним существом (вес 50)
  - С сильным существом (вес 70)
- Поспать весь день – параметры меняются как ночью, т.е. по факту у игрока получается три ночи подряд.

Также не забудьте покрыть ваш код тестами.
