import {PrivateApi, PublicApi} from "typescript-axios";

const publicApi = new PublicApi(undefined, process.env.REACT_APP_API_URL);
const privateApi = new PrivateApi(undefined, process.env.REACT_APP_API_URL);

export {publicApi, privateApi};
