import http from 'k6/http';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export const options = {
  vus: 1000000,
  duration: '15s',
  summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(95)', 'p(99)', 'p(99.99)', 'count'],
};

export default function () {
  const url = 'http://app:8080/shorten';
  const payload = JSON.stringify({
    url: "https://www.google.com"
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
}

export function handleSummary(data) {
  return {
   "/scripts/shorten_link_report.html": htmlReport(data),
  };
}