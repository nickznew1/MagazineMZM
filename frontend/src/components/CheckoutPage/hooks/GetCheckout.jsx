import {useEffect, useState} from "react";
import {useAuth} from "../../../context/AuthContext";
import {config} from "../../../config";


export function GetCheckout() {
    const {token}=useAuth();
    const [inputs, setInputs] = useState(null)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        if (!token) {
            setLoading(false);
            setError('error')
            return
        }
        const getCheckout = async () => {
            try{
                setLoading(true)
                const [profileData, cartData] = await Promise.all ([
                    fetch(`${config.apiUrl}/checkout`, {
                        headers: token ? {Authorization: `Bearer ${token}`} : {}
                    }).then(res => res.json()),
                    fetch(`${config.apiUrl}/cart/`,{
                        headers: token ? {Authorization: `Bearer ${token}`} : {}
                    }).then(res =>res.json())
                ])
                const resJson = {
                    profileData : profileData,
                    cartData: cartData,
                }
                if (resJson) {
                    setInputs(resJson)
                    setLoading(false)
                }

            } catch (error) {
                setInputs(null)
                setError(error)
            } finally {
                setLoading(false)
            }
        }
        getCheckout();
    },[token])

    return {inputs, setInputs, loading, error, token}
}