/**
 * randomFunc can return result or error
 */
async function randomFunc() {
	const rand = Math.random() * 10
	if (rand <= 4) {
		return true
	}

	throw new Error('my error')
}

randomFunc()
	.then((val) => {
		console.log('SUCCESS:', val)
	})
	.catch((err) => {
		console.error('ERROR:', err)
	})

function randomFuncPromise() {
	return new Promise((resolve, reject) => {
		const rand = Math.random() * 10
		if (rand <= 4) {
			resolve(true)
		}

		reject(new Error('my error, promise'))
	})
}

async function randomFuncPromiseExecutor() {
	try {
		const val = await randomFuncPromise()
		return val
	} catch (error) {
		console.log('error: ', error)
		throw new Error('handled error')
	}
}

randomFuncPromiseExecutor()
	.then((val) => {
		console.log('SUCCESS:', val)
	})
	.catch((err) => {
		console.error('ERROR:', err)
	})
