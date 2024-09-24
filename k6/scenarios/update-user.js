import http from 'k6/http'
import { randomIntBetween, randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

import { check, sleep } from 'k6'


export default function () {

  const data = {
    "id": randomIntBetween(0, 1000000),
    "name": randomString(),
    "email": randomString(),
  }

  let res = http.post('http://localhost:8080/v1/users', JSON.stringify(data))
  check(res, { 'success create user': (r) => r.status === 201 })

  let getRes = http.put(`http://localhost:8080/v1/users/${data["id"]}`, data)
  check(getRes, { 'success get user': (r) => r.status === 202 })
  sleep(0.1)
}