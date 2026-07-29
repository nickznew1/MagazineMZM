import {useAuth} from "../../../context/AuthContext";
import {useEffect, useState} from "react";
import {config} from "../../../config";

export function UserCart() {
    const {token} = useAuth()
    const [cart, setCart] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {

        if (!token) {
            setLoading(false);
            setError(error)
            return
        }

        if (token) {
            const fetchCart = async () => {
                try {
                    setLoading(true);
                    setError(null);
                    const response = await fetch(`${config.apiUrl}/cart/`, {
                        headers: token ? {Authorization: `Bearer ${token}`} : {}
                    });
                    if (!response.ok) {
                        setError(error)
                    } else {
                        const data = await response.json();
                        setCart(data)
                    }
                } catch (error) {
                    setCart(null)
                    setError(error.message)
                } finally {
                    setLoading(false);
                }
            }
            fetchCart();
        }

    }, []);

    return {user:cart, token, setCart, loading, error}
}