import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 1000, //代表模擬用戶數量
  rps: 1000, //代表每秒發送請求數量
  duration: '1s', //代表執行時間
};

export default function () {
  let ret = {
    "location": "L1",
    "timestamp": new Date(),
    "data": {
        "a": 18,
        "b": 2,
        "c": 6,
        "d": 28
    }
  }
  http.post('http://34.81.206.60/api/order',JSON.stringify(ret),{headers: { 'Content-Type': 'application/json' },}); //測試目標網址
  // http.get('http://34.81.206.60/api/report?location=l1&date=2023-01-01')
}