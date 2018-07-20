import axios from 'axios';

export async function getToken(params) {
  const res = await axios.post('http://localhost:3000/api/token', params)
  return res.data;
}