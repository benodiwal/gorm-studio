import Axios from "axios";

const AxiosClient = Axios.create({
    baseURL: import.meta.env.VITE_BACKEND_BASE_URL
});

export default AxiosClient;
