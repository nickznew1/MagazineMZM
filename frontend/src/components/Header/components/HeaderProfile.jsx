import {useAuth} from "../../../context/AuthContext";
import {useNavigate} from "react-router";
import {memo} from "react";
import "./styles.css"


const HeaderProfile = () => {
    const {user} = useAuth()
    const navigate = useNavigate()
    return (
        <>
        {user ? (
                <div className ="header__cabinet-button">
                    <button onClick = {()=>navigate(`/profile/`)}>Кабинет</button>
                </div>
            ) : (
                <div className ="header__cabinet-button">
                    <button onClick = {()=>navigate('/auth')}>Кабинет</button>
                </div>
            )}
            </>
    )
}

export default memo(HeaderProfile);