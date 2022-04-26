import http from 'k6/http';

export const options = {
  stages: [
    { duration: '10s', target: 20 },
    { duration: '10s', target: 40 },
    { duration: '20s', target: 80 },
    { duration: '20s', target: 160 },
    { duration: '20s', target: 320 },
    { duration: '20s', target: 640 },
  ],
};

export default function () {
  http.get('http://localhost:8080/health');
}
