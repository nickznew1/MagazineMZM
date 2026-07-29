import {useAuth} from "../../../context/AuthContext";
import {useNavigate} from "react-router";
import "./style.css"
import {memo} from "react";


const ProfileLogout = () => {
    const {logout} = useAuth()
    const navigate = useNavigate();

    const handleLogout = async () => {
        await logout();
        navigate('/')
    }

    return (
        <div className ="user-profile-logout-wrapper">
            <div className ="user-profile-logout__btn">
        <button type ="button" onClick={handleLogout}>Выйти из профиля</button>
            </div>
        </div>
    )
}

export default memo(ProfileLogout);