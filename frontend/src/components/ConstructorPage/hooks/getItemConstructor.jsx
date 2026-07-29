import {useEffect, useState} from "react";
import {config} from "../../../config";


export function GetItemConstructor(idUrl) {
    const [item, setItem] = useState(null)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        const getItemConstructor = async () => {
            try {
                setLoading(true)
                setError(null)
                const response = await fetch(`${config.apiUrl}/item/${idUrl}`)
                if (response.ok){
                    const item = await response.json()
                    setItem(item)
                    setLoading(false)
                } else {
                    setLoading(true)
                    setItem(null)
                }
            } catch (error) {
                setError(error)
                setItem(null)
            }finally {
                setLoading(false)
            }
        }
        getItemConstructor();
    }, [idUrl]);


    return { item, error, loading };
}