import JSEncrypt from "jsencrypt";
import {requestPack} from "./requests.js";

const rsaEncrypt = async(password)=> {
    let encrypt = new JSEncrypt();
    // 这里的公钥需要替换为你自己的公钥
    let publicKey = await requestPack.get("user/pub_key")
    encrypt.setPublicKey(publicKey);
    return encrypt.encrypt(password);
}

export {rsaEncrypt}