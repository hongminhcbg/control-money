import http from 'k6/http';
import {check, group, sleep} from 'k6';
import { Rate } from 'k6/metrics';

export let errorRate = new Rate('errors');
export let options = {
  stages: [
    { duration: "1m", target: 30 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
    { duration: "1m", target: 20 }, // stay at 100 users for 10 minutes
    { duration: "1m", target: 10 }, // ramp-down to 0 users
  ],
  thresholds: {
    'http_req_duration': ['p(99)<1500'], // 99% of requests must complete below 1.5s
    'logged in successfully': ['p(99)<1500'], // 99% of requests must complete below 1.5s
    errors: ['rate<0.1'], // <10% errors
  }
};

const BASE_URL = 'http://localhost:8080'; 
const USERNAME = 'hongminh6';
const PASSWORD = 'hongminh229297';
const API_KEY = 'hongminh229297'
const apiHeaders = { headers: {
    'api-key': `${API_KEY}`,
    'Content-Type': 'application/json',
  }};

export default () => {
    let body = {
        username: USERNAME,
        password: PASSWORD,
    }
  let loginRes = http.post(`${BASE_URL}/api/v1/account/login`, JSON.stringify(body), apiHeaders);  

  let isSuccess = check(loginRes, { 'logged in successfully': (resp) => resp.status === 200 });

  if (!isSuccess) {
    console.log("send request error")
    errorRate.add(1)
  }

  sleep(1);
}