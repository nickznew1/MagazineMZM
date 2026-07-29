
import {useAuth} from "../../../context/AuthContext";
import {useNavigate} from "react-router";
import {memo} from "react";
import "./styles.css"



const HeaderCart = () => {
    const {user, error, loading} = useAuth()

    const navigate = useNavigate()

    if (loading){
        return <div>loading...</div>
    }

    if (error){
        return <div>{error}</div>;
    }

    return (
        <div className= "header__cart">
            {user ? (
                <button onClick={() =>navigate(`/cart/`)}>
                    Корзина
                </button>
            ) : (
                <button onClick={() =>navigate(`/auth`)}>
                    Корзина
                </button>
            )}
        </div>
    )

}

export default memo(HeaderCart);