import http from 'k6/http';

export const options = {
  stages: [
    { duration: '5s', target: 20 },
    { duration: '5s', target: 40 },
    { duration: '10s', target: 80 },
    { duration: '10s', target: 160 },
    { duration: '10s', target: 320 },
    { duration: '10s', target: 640 },
    { duration: '10s', target: 1280 },
  ],
};

export default function () {
  http.get('http://localhost:8080/health');
}
