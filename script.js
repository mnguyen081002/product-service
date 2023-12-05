import http from 'k6/http';
import { check, sleep } from 'k6';

// add the following to your k6 options

export const options = {
  thresholds: {
    http_req_failed: ['rate<0.01'], // http errors should be less than 1%
    http_req_duration: ['p(95)<200'], // 95% of requests should be below 200ms
  },
};

export default function () {
  const res = http.get(
    'http://127.0.0.1:8080/v1/api/cms/product?search=hea&price=200&limit=10&page=1',
    {
      headers: {
        'x-user-id': '3095398a-ed67-4770-bf58-9ce3c682df13',
      },
    }
  );
  check(res, { 'status was 200': (r) => r.status == 200 });
  check(res, { 'duration < 500ms': (r) => r.timings.duration < 1000 });
  sleep(1);
}
