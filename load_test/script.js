import { check, sleep } from "k6";
import http from "k6/http";
import { randomString, randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
    stages: [
        {
            duration: '10s',
            target: 100
        }
    ],
    thresholds: {
        http_req_failed: ['rate<0.001'], // the error rate must be lower than 0.1%
        http_req_duration: ['p(90)<600'], // 90% of requests must complete below 600ms
        http_req_receiving: ['max<500'], // slowest request below 500ms
        iteration_duration: ['p(95)<20000'], // 95% of requests must complete below 20000ms
    },
};

export default function() {
    const title = randomString(15);
    const msg = randomString(101);
    const user_id = randomIntBetween(1,3);
    let body = {
        "ticket_title": title,
        "ticket_msg": msg,
        "user_id": user_id
    }
    let resPost = http.post("http://localhost:8080/tickets", JSON.stringify(body));
    check(resPost, {
        "is status 201": (r) => r.status === 201
    });
    sleep(1);
};
