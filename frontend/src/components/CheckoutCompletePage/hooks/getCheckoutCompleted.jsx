import {useAuth} from "../../../context/AuthContext";
import {useEffect, useState} from "react";
import {useParams} from "react-router";
import {config} from "../../../config";


export function GetCheckoutCompleted(id) {
        const {token:token}=useAuth();

        const [checkout, setCheckout] = useState(null)
        const [loading, setLoading] = useState(true)
        const [error, setError] = useState(null)
        console.log(token)
        useEffect(() => {
            if (!token) {
                setLoading(false);
                setError('Not autorized')
                return
            }
            const getCheckoutComplete = async () => {
                try{
                    setLoading(true)
                    const response = await fetch (`${config.apiUrl}/checkout/complete/${id}`, {
                        headers: token ? { Authorization: `Bearer ${token}` } : {}
                    });

                    if (response.ok) {
                        const checkout = await response.json()
                        setCheckout(checkout)
                        setLoading(false)
                    }
                } catch (error) {
                    setCheckout(checkout)
                    setError(error)
                } finally {
                    setLoading(false)
                }
            }
            getCheckoutComplete();
        },[token])

        return {checkout, setCheckout, loading, error, token}
}