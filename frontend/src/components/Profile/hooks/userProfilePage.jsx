import {useAuth} from "../../../context/AuthContext";
import {useEffect, useState} from "react";
import {config} from "../../../config";

export function UserProfilePage(loginURL) {
    const {token, loading:authLoading, roles, userRole} = useAuth()
    const [profile, setProfile] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    useEffect(() => {
        if (authLoading) {return}
        if (!token) {
            setLoading(false);
            setError(error)
            return
        }
        const fetchProfile = async () => {
            try {
                setLoading(true);
                setError(null);
                const [profileData, itemData] = await Promise.all ([
                    fetch(`${config.apiUrl}/profile/`,
                        {headers: token ? { Authorization: `Bearer ${token}` } : {}}).then(res => res.json()),

                    fetch(`${config.apiUrl}/applications/all`,
                        {headers: token ? { Authorization: `Bearer ${token}` } : {}}).then(res =>res.json())
                ])
                const resJson = {
                    profileData:profileData,
                    itemData:itemData,
                }
                setProfile(resJson)
            } catch (error) {
                setProfile(null);
                setError(error.message)
            } finally {
                setLoading(false);
            }
        }
        fetchProfile();
        }, [authLoading]);
return {user:profile, loading, error, setProfile, roles, userRole}
}