import {useEffect, useState} from "react";
import {config} from "../../../config";


export function GetItemCard(idUrl) {
    const [item, setItem] = useState(null)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        const getItem = async () => {
            try {
                setLoading(true)
                setError(null)
                const [itemData, itemSpec] = await Promise.all ([
                    fetch(`${config.apiUrl}/item/${idUrl}`).then(res => res.json()),
                    fetch(`${config.apiUrl}/item/spec/${idUrl}`).then(res =>res.json())
                ])

                const resJson = {
                    itemData : itemData,
                    specData: itemSpec,
                }

                setItem(resJson)
            } catch (error) {
                setError(error)
                setItem(null)
                }finally {
                setLoading(false)
            }
                }
                getItem();
            }, [idUrl]);


return { item, error, loading, setItem };
}