import axios from "axios"


const prefix = "/backend"

const fetchChat = (roomId) => {
  return axios.post(`${prefix}/api/fetchChat/${roomId}`)
}

export default fetchChat