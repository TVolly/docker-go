import http from 'k6/http';
import { sleep } from 'k6';
export let options = {
  vus: 20,
  duration: '30s',
};
export default function() {
  let body = JSON.stringify({
    name: "Community"
  });

  http.post('http://localhost:8080/communities', body);
  sleep(1);
}