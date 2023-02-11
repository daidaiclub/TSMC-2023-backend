import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 1, //代表模擬用戶數量
  duration: '1s', //代表執行時間
};

export default function () {
  http.post('http://34.81.206.60/api/order','Content-type: application/json',{
    "location": "l1",
    "timestamp": "2023-01-01T20:18:56.424+08:00",
    "data": {
        "a": 18,
        "b": 2,
        "c": 6,
        "d": 28}}); //測試目標網址
}