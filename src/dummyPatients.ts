import { faker } from '@faker-js/faker'
import { generateBSN } from 'bsn-js'

export enum Gender {
  FEMALE = 'female',
  MALE = 'male',
  OTHER = 'other',
}

export type Patient = {
  birthdate: string
  BSN: string
  email: string
  familyName: string
  gender: Gender
  givenName: string
}

export function randomEnumKey<T>(enumeration: T): T[keyof T] {
  const enumValues = Object.values(enumeration) as unknown as T[keyof T][]
  const randomIndex = Math.floor(Math.random() * enumValues.length)
  return enumValues[randomIndex]
}

export function generateGivenName(gender: Gender): string {
  return gender === Gender.FEMALE
    ? faker.name.firstName('female')
    : gender === Gender.MALE
    ? faker.name.firstName('male')
    : faker.name.firstName()
}

export function formmatDate(date: Date): string {
  const yyyy = date.getFullYear()
  const mm = date.getMonth() + 1
  const dd = date.getDate()

  const formattedDate = `${dd < 10 ? '0' + dd : dd}/${
    mm < 10 ? '0' + mm : mm
  }/${yyyy}`

  return formattedDate
}

export function generateDummyData(): Patient {
  const birthdate = formmatDate(
    faker.date.birthdate({ min: 1920, max: 2020, mode: 'year' })
  )
  const gender = randomEnumKey(Gender)
  const givenName = generateGivenName(gender)
  const familyName = faker.name.lastName()
  const email = faker.internet.email(givenName, familyName)
  const BSN = generateBSN()

  return { birthdate, BSN, email, familyName, gender, givenName }
}
