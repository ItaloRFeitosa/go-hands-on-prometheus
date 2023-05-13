import http from 'k6/http';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export const options = {
  discardResponseBodies: true,

  scenarios: {
    shorten: {
      executor: 'ramping-arrival-rate',

      // Start at 300 iterations per `timeUnit`
      startRate: 300,

      // Start `startRate` iterations per minute
      timeUnit: '1m',

      // Pre-allocate necessary VUs.
      preAllocatedVUs: 50,


      stages: [
        // Start 300 iterations per `timeUnit` for the first minute.
        { target: 300, duration: '1m' },

        // Linearly ramp-up to starting 600 iterations per `timeUnit` over the following two minutes.
        { target: 3000, duration: '2m' },

        // Cntinue starting 600 iterations per `timeUnit` for the following four minutes.
        { target: 6000, duration: '4m' },

        // Linearly ramp-down to starting 60 iterations per `timeUnit` over the last two minute.
        { target: 600, duration: '2m' },
      ],
    },
  },
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