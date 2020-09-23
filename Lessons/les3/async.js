// help https://github.com/frontarm/async-javascript-cheatsheet

/**
 * Задача рассмотреть 3 варианта решения (callback, Promise, async / await) одной задачи:
 * Есть массив со значениями 1 2 3, требуется добавить сразу цифру 1, через
 * 500 мс 2, ещё через 500 мс 3 (для Promise ещё добавить 4 через также 500 мс)
 */
const timeout = 500

let elCb = [1, 2, 3]

function diffTime(start) {
	return new Date() - start
}

/**
 * Callback вызывается в качестве ответа на событие, такой подход позволяет JS быть мультизадачным
 *
 * Callback Hell:
 */
function addValuesCb() {
	// save start date
	const start = new Date()

	elCb.push(1)
	console.log('(cb)(1):', elCb, 'time:', diffTime(start))

	// wait 500 ms
	setTimeout(() => {
		elCb.push(2)
		console.log('(cb)(2):', elCb, 'time:', diffTime(start))

		// wait another 500 ms
		setTimeout(() => {
			elCb.push(3)
			console.log('(cb)(3):', elCb, 'time:', diffTime(start))
		}, timeout)
	}, timeout)
}

addValuesCb()

/**
 * delay wait for specified ms before to go next
 * @param {number} ms - time to wait
 */
function delay(ms) {
	return new Promise((resolve) => {
		setTimeout(resolve, ms)
	})
}

let elP = [1]

// Promise wrong realization
function addValuesPromise() {
	const start = new Date()
	elP.push(1)
	console.log('(p)(1):', elP, 'time:', diffTime(start))

	delay(timeout)
		.then(() => {
			// после того как delay resolve Promise
			// выполнится этот код, который добавит ещё один элемент
			elP.push(2)
			console.log('(p)(2):', elP, 'time:', diffTime(start))
		})
		.then(() => {
			// после предыдущего then выполнится этот
			// на первый взгляд можно подумать, что тут мы не избавляемся от callback hell
			// и если нам требуется зависимость от предыдущего, то это и правда так
			// Это потому, что мы допустили тут ошибку
			//
			// Попробуйте запустить и вы увидите, что 3 и 4 надо одинаковое примерно время
			// чтобы выполнить запрос, но мы хотим, чтобы 4 запускался после того как 3
			// отработал полностью
			//
			// Почему это происходит? Обратите внимания, в этой функции нет ничего особенного
			// т.е. запускается delay, а затем сразу же код идёт дальше и запускает следующий then
			// чтобы исправить эту проблему, эта функция должна возвращать Promise только тогда,
			// когда отработает delay!
			delay(timeout).then(() => {
				elP.push(3)
				console.log('(p)(3):', elP, 'time:', diffTime(start))
			})
		})
		.then(() => {
			delay(timeout).then(() => {
				elP.push(4)
				console.log('(p)(4):', elP, 'time:', diffTime(start))
			})
		})
}

addValuesPromise()

let elPT = [1]

// Promise correct realization
function addValuesPromiseT() {
	const start = new Date()
	elPT.push(1)
	console.log('(pT)(1):', elPT, 'time:', diffTime(start))

	delay(timeout)
		.then(() => {
			elPT.push(2)
			console.log('(pT)(2):', elPT, 'time:', diffTime(start))
		})
		.then(() => {
			// теперь мы добавили Promise, он делает только resolve, т.е. никогда не
			// может вернуть ошибку, функция не будет завершаться сразу же, она
			// завершится только тогда, когда мы вызываем resolve
			// Обратите внимание, что в resolve добавлена строка, это сделано просто для
			// примера, в следующем then мы получим это значение
			return new Promise((resolve) => {
				delay(timeout).then(() => {
					elPT.push(3)
					console.log('(pT)(3):', elPT, 'time:', diffTime(start))
					resolve('Hello!')
				})
			})
		})
		.then((val) => {
			// тут мы получаем то, что вернул предыдущий then, также этот then выполнится
			// только когда сработает resolve предыдущей функции, не раньше!
			console.log('val: ', val)
			delay(timeout).then(() => {
				elPT.push(4)
				console.log('(pT)(4):', elPT, 'time:', diffTime(start))
			})
		})
}

addValuesPromiseT()

let elA = [1, 2, 3]

async function addValuesAsync() {
	// последний вариант и то, что мы будем использовать постоянно - async
	// знать, что такое callback, Promise - необходимо, но сейчас самый удобный вариант
	// и в плане читаемости и в плане реализации - async await.

	// обратите внимание, что await может применяться только в async функциях!
	// Это синтаксический сахар для Promise.
	const start = new Date()
	elA.push(1)
	console.log('(a)(1):', elA, 'time:', diffTime(start))
	await delay(timeout)
	elA.push(2)
	console.log('(a)(2):', elA, 'time:', diffTime(start))
	await delay(timeout)
	console.log('(a)(3):', elA, 'time:', diffTime(start))
}

addValuesAsync()
