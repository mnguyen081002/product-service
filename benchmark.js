import http from 'k6/http';
import { check, sleep } from 'k6';

const TOTAL_RPS = 200;
function rate(percentage) {
  return TOTAL_RPS * percentage;
}

export const options = {
  scenarios: {
    busy: {
      exec: 'get_product_endpoint',
      executor: 'constant-arrival-rate',
      duration: '60s',
      preAllocatedVUs: 80,
      rate: rate(0.9), // 90%
      tags: { test_type: 'get-api' },
    },
    lazy: {
      exec: 'decrease_quantity_endpoint',
      executor: 'constant-arrival-rate',
      duration: '60s',
      preAllocatedVUs: 20,
      tags: { test_type: 'des-api' },
      rate: rate(0.1), // // 10%
    },
  },
  thresholds: {
    http_req_failed: ['rate<0.01'], // http errors should be less than 1%
    http_req_duration: ['p(95)<200'], // 95% of requests should be below 200ms
  },
};

export function get_product_endpoint() {
  const res = http.get(
    'http://127.0.0.1:8081/v1/api/cms/product/?search=hea&price=200&limit=20&page=1',
    {
      headers: {
        'x-user-id': '3095398a-ed67-4770-bf58-9ce3c682df13',
      },
    }
  );
  check(res, { 'get_product_endpoint status was 200': (r) => r.status == 200 });
  check(res, { 'get_product_endpoint duration < 500ms': (r) => r.timings.duration < 500 });
  sleep(1);
}

export function decrease_quantity_endpoint() {
  const payload = JSON.stringify({
    quantity: 1,
  });
  const res = http.patch(
      'http://127.0.0.1:8081/v1/api/cms/product/00001921-8e37-4862-ba6a-2c8a2742cdab',
      payload,
      {
        headers: {
          'x-user-id': '3095398a-ed67-4770-bf58-9ce3c682df13',
          'Content-Type': 'application/json',
        },
      }
  );
  check(res, { 'decrease_quantity_endpoint status was 200': (r) => r.status == 200 });
  check(res, { 'decrease_quantity_endpoint duration < 500ms': (r) => r.timings.duration < 500 });
  sleep(1);
}