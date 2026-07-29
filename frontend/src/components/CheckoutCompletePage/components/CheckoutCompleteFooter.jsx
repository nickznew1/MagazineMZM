import "../components/styles.css"
import {useNavigate} from "react-router";
import {memo} from "react";


const CheckoutCompleteFooter = () =>{
    const navigate = useNavigate();
    return (
        <>
        <div className = "checkout-complete-footer">
            <p id ="long-text">Ваша заявка оформлена. В скором времени наш менеджер свяжется с вами по почте или по телефону для согласования заявки, оплате и доставки. Статус заявки можно отслеживать в личном <a id ="profile-link" href ="/profile/">профиле</a>. Также вы можете связаться с нами по нашим контактам для получения дополнительной информации. </p>
            <h2 id ="thanks">Спасибо, что выбрали нас!</h2>
        </div>
    <button id = "button-main-page" onClick={()=>navigate('/')}>Вернуться на главную страницу</button>
        </>
    )
}

export default memo(CheckoutCompleteFooter);