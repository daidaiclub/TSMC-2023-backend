import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 600, //代表模擬用戶數量
  rps: 600, //代表每秒發送請求數量
  duration: '10s', //代表執行時間
};

export default function () {
  let ret = {
    "location": "L1",
    "timestamp": new Date(),
    "data": {
        "a": parseInt((Math.random() * 10)),
        "b": parseInt((Math.random() * 10)),
        "c": parseInt((Math.random() * 10)),
        "d": parseInt((Math.random() * 10)),
    }
  }
  http.post('http://34.81.206.60/api/order',JSON.stringify(ret),{headers: { 'Content-Type': 'application/json' },}); //測試目標網址
  // http.get('http://34.81.206.60/api/report?location=l1&date=2023-01-01')
}