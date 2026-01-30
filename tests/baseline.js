import http from 'k6/http'
import { sleep } from 'k6'

export const options = {
    vus: 10,
    duration: '30s'
}

export default function() {
    const url = 'http://localhost:3001/beerstyles/temperature'

    const payload = JSON.stringify({
        temperature: -3
    })

    const params = {
        headers: {
            'Content-Type': 'application/json'
        }
    }

    http.request('GET', url, payload, params);
    sleep(0.5)
}