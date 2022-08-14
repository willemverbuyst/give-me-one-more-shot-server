import express from 'express'
import cors from 'cors'
import morgan from 'morgan'
import { PORT } from './constants'
import { generateDummyData } from './dummyPatients'

const app = express()

app.use(cors())
app.use(morgan('dev'))

app.get('/', (_req, res) => {
  res.send({ message: 'hello world' })
})

app.get('/patients', (_req, res) => {
  const dummyPatients = Array.from({ length: 10 }).map(() =>
    generateDummyData()
  )
  res.send({ status: 'success', data: dummyPatients })
})

app.listen(PORT, () => console.log(`Server is running on ${PORT}`))
