import {useAuth} from "../../../context/AuthContext";


export function GetAdminPage() {
    const {user, loading, error} = useAuth();

    return {user, loading,error}
}