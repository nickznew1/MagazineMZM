import CheckoutCompletedItems from "../components/CheckoutCompletePage/components/CheckoutCompletedItems";
import {
    GetCheckoutCompleted,
} from "../components/CheckoutCompletePage/hooks/getCheckoutCompleted";
import {useParams} from "react-router";
import CheckoutCompletedInfo from "../components/CheckoutCompletePage/components/CheckoutCompleteInfo";
import CheckoutCompleteFooter from "../components/CheckoutCompletePage/components/CheckoutCompleteFooter";
import NotAuthorized from "../features/NotAuthorized/NotAuthorized";


const CheckoutComplete = () =>{
    const {id} = useParams()
    const checkoutComplete = GetCheckoutCompleted(id)
     if (checkoutComplete.loading){
         return <div>loading...</div>
     }
     if (checkoutComplete.error){
        return <NotAuthorized/>
     }

    return (
        <div className="checkout-complete-wrapper">
            <div className="checkout-complete">
        <CheckoutCompletedItems
            checkoutComplete = {checkoutComplete.checkout}
        />
            <CheckoutCompletedInfo
                checkoutComplete={checkoutComplete.checkout}
                />
                <CheckoutCompleteFooter />
            </div>
        </div>
    )
}

export default CheckoutComplete;