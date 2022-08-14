import express from 'express'
import cors from 'cors'
import morgan from 'morgan'

const app = express()

app.use(cors())
app.use(morgan('dev'))

app.get('/', (_req, res) => {
  res.send({ message: 'hello world' })
})

app.listen(4000, () => console.log('hello world'))
