import http from 'k6/http';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export const options = {
  scenarios: {
    shorten: {
      executor: 'ramping-arrival-rate',

      // Start at 300 iterations per `timeUnit`
      startRate: 300,

      // Start `startRate` iterations per minute
      timeUnit: '1s',

      // Pre-allocate necessary VUs.
      preAllocatedVUs: 50,


      stages: [
        { target: 300, duration: '30s' },
        { target: 5000, duration: '10s' },
        { target: 600, duration: '60s' },
        { target: 2000, duration: '60s' },
        { target: 5000, duration: '10s' },
        { target: 1000, duration: '60s' },
        { target: 100, duration: '60s' },
        { target: 5000, duration: '10s' },
        { target: 100, duration: '60s' },
      ],
    },
  },
};

export function setup() {
  const url = 'http://app:8080/shorten';
  const payload = JSON.stringify({
    url: 'http://app:8080/redirect-test'
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);

  return {
    slug: res.json().slug
  }
}

export default function (data) {
  const url = 'http://app:8080/' + data.slug;

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.get(url, params);
}

export function handleSummary(data) {
  return {
    "/scripts/shorten_link_report.html": htmlReport(data),
  };
}